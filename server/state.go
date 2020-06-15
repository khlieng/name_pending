package server

import (
	"log"
	"sync"
	"time"

	"github.com/khlieng/dispatch/pkg/irc"
	"github.com/khlieng/dispatch/pkg/session"
	"github.com/khlieng/dispatch/storage"
)

const (
	// AnonymousUserExpiration is the time to wait before removing an anonymous
	// user that has no irc or websocket connections
	AnonymousUserExpiration = 1 * time.Minute
)

// State is the live state of a single user
type State struct {
	stateData

	networks        map[string]*storage.Network
	pendingDCCSends map[string]*irc.DCCSend

	ws        map[string]*wsConn
	broadcast chan WSResponse

	srv        *Dispatch
	user       *storage.User
	expiration *time.Timer
	reset      chan time.Duration
	lock       sync.Mutex
}

func NewState(user *storage.User, srv *Dispatch) *State {
	return &State{
		stateData:       stateData{m: map[string]interface{}{}},
		networks:        make(map[string]*storage.Network),
		pendingDCCSends: make(map[string]*irc.DCCSend),
		ws:              make(map[string]*wsConn),
		broadcast:       make(chan WSResponse, 32),
		srv:             srv,
		user:            user,
		expiration:      time.NewTimer(AnonymousUserExpiration),
		reset:           make(chan time.Duration, 1),
	}
}

func (s *State) network(host string) (*storage.Network, bool) {
	s.lock.Lock()
	n, ok := s.networks[host]
	s.lock.Unlock()

	return n, ok
}

func (s *State) client(host string) (*irc.Client, bool) {
	if network, ok := s.network(host); ok {
		return network.Client(), true
	}
	return nil, false
}

func (s *State) setNetwork(host string, network *storage.Network) {
	s.lock.Lock()
	s.networks[host] = network
	s.lock.Unlock()

	s.reset <- 0
}

func (s *State) deleteNetwork(host string) {
	s.lock.Lock()
	delete(s.networks, host)
	s.lock.Unlock()

	s.resetExpirationIfEmpty()
}

func (s *State) numIRC() int {
	s.lock.Lock()
	n := len(s.networks)
	s.lock.Unlock()

	return n
}

func (s *State) pendingDCC(filename string) (*irc.DCCSend, bool) {
	s.lock.Lock()
	pack, ok := s.pendingDCCSends[filename]
	s.lock.Unlock()
	return pack, ok
}

func (s *State) setPendingDCC(filename string, pack *irc.DCCSend) {
	s.lock.Lock()
	s.pendingDCCSends[filename] = pack
	s.lock.Unlock()
}

func (s *State) deletePendingDCC(filename string) {
	s.lock.Lock()
	delete(s.pendingDCCSends, filename)
	s.lock.Unlock()
}

func (s *State) setWS(addr string, w *wsConn) {
	s.lock.Lock()
	s.ws[addr] = w
	s.lock.Unlock()

	s.reset <- 0
}

func (s *State) deleteWS(addr string) {
	s.lock.Lock()
	delete(s.ws, addr)
	s.lock.Unlock()

	s.resetExpirationIfEmpty()
}

func (s *State) numWS() int {
	s.lock.Lock()
	n := len(s.ws)
	s.lock.Unlock()

	return n
}

func (s *State) sendJSON(t string, v interface{}) {
	s.broadcast <- WSResponse{t, v}
}

func (s *State) sendLastMessages(network, channel string, count int) {
	messages, hasMore, err := s.user.LastMessages(network, channel, count)
	if err == nil && len(messages) > 0 {
		res := Messages{
			Network:  network,
			To:       channel,
			Messages: messages,
		}

		if hasMore {
			res.Next = messages[0].ID
		}

		s.sendJSON("messages", res)
	}
}

func (s *State) sendMessages(network, channel string, count int, fromID string) {
	messages, hasMore, err := s.user.Messages(network, channel, count, fromID)
	if err == nil && len(messages) > 0 {
		res := Messages{
			Network:  network,
			To:       channel,
			Messages: messages,
			Prepend:  true,
		}

		if hasMore {
			res.Next = messages[0].ID
		}

		s.sendJSON("messages", res)
	}
}

func (s *State) resetExpirationIfEmpty() {
	if s.numIRC() == 0 && s.numWS() == 0 {
		s.reset <- AnonymousUserExpiration
	}
}

func (s *State) kill() {
	s.lock.Lock()
	for _, ws := range s.ws {
		ws.conn.Close()
	}
	for _, network := range s.networks {
		network.Client().Quit()
	}
	s.lock.Unlock()
}

func (s *State) run() {
	for {
		select {
		case res := <-s.broadcast:
			s.lock.Lock()
			for _, ws := range s.ws {
				ws.out <- res
			}
			s.lock.Unlock()

		case <-s.expiration.C:
			s.srv.states.delete(s.user.ID)
			s.user.Remove()
			return

		case duration := <-s.reset:
			if duration == 0 {
				s.expiration.Stop()
			} else {
				s.expiration.Reset(duration)
			}
		}
	}
}

type stateData struct {
	m    map[string]interface{}
	lock sync.Mutex
}

func (s *stateData) Get(key string) interface{} {
	s.lock.Lock()
	v := s.m[key]
	s.lock.Unlock()
	return v
}

func (s *stateData) Set(key string, value interface{}) {
	s.lock.Lock()
	s.m[key] = value
	s.lock.Unlock()
}

func (s *stateData) String(key string) string {
	if v, ok := s.Get(key).(string); ok {
		return v
	}
	return ""
}

func (s *stateData) Int(key string) int {
	if v, ok := s.Get(key).(int); ok {
		return v
	}
	return 0
}

func (s *stateData) Bool(key string) bool {
	if v, ok := s.Get(key).(bool); ok {
		return v
	}
	return false
}

type stateStore struct {
	states       map[uint64]*State
	sessions     map[string]*session.Session
	sessionStore storage.SessionStore
	lock         sync.Mutex
}

func newStateStore(sessionStore storage.SessionStore) *stateStore {
	store := &stateStore{
		states:       make(map[uint64]*State),
		sessions:     make(map[string]*session.Session),
		sessionStore: sessionStore,
	}

	sessions, err := sessionStore.Sessions()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[Init] %d sessions", len(sessions))

	for _, session := range sessions {
		if !session.Expired() {
			store.sessions[session.Key()] = session
		} else {
			go sessionStore.DeleteSession(session.Key())
		}
	}

	return store
}

func (s *stateStore) run() {
	pruneSessions := time.Tick(time.Minute * 5)
	for {
		select {
		case <-pruneSessions:
			s.lock.Lock()
			for key, session := range s.sessions {
				if session.Expired() {
					s.internalDeleteSession(key)
				}
			}
			s.lock.Unlock()
		}
	}
}

func (s *stateStore) get(id uint64) *State {
	s.lock.Lock()
	state := s.states[id]
	s.lock.Unlock()
	return state
}

func (s *stateStore) set(state *State) {
	s.lock.Lock()
	s.states[state.user.ID] = state
	s.lock.Unlock()
}

func (s *stateStore) delete(id uint64) {
	s.lock.Lock()
	delete(s.states, id)
	for key, session := range s.sessions {
		if session.UserID == id {
			delete(s.sessions, key)
			go s.sessionStore.DeleteSession(key)
		}
	}
	s.lock.Unlock()
}

func (s *stateStore) getSession(key string) *session.Session {
	s.lock.Lock()
	session := s.sessions[key]
	s.lock.Unlock()
	return session
}

func (s *stateStore) setSession(session *session.Session) {
	s.lock.Lock()
	s.sessions[session.Key()] = session
	s.lock.Unlock()
	go s.sessionStore.SaveSession(session)
}

func (s *stateStore) deleteSession(key string) {
	s.lock.Lock()
	s.internalDeleteSession(key)
	s.lock.Unlock()
}

func (s *stateStore) internalDeleteSession(key string) {
	id := s.sessions[key].UserID
	delete(s.sessions, key)

	n := 0
	for _, session := range s.sessions {
		if session.UserID == id {
			n++
		}
	}

	state := s.states[id]
	if n == 0 {
		delete(s.states, id)
	}

	go func() {
		if n == 0 {
			// This anonymous user is not reachable anymore since all sessions have
			// expired, so we clean it up
			state.kill()
			state.user.Remove()
		}

		s.sessionStore.DeleteSession(key)
	}()
}
