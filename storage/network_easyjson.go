//v1: false// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package storage

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC5839400DecodeGithubComKhliengDispatchStorage(in *jlexer.Lexer, out *Network) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "host":
			out.Host = string(in.String())
		case "port":
			out.Port = string(in.String())
		case "tls":
			out.TLS = bool(in.Bool())
		case "serverPassword":
			out.ServerPassword = string(in.String())
		case "nick":
			out.Nick = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "realname":
			out.Realname = string(in.String())
		case "account":
			out.Account = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "features":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Features = make(map[string]interface{})
				} else {
					out.Features = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 interface{}
					if m, ok := v1.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v1.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v1 = in.Interface()
					}
					(out.Features)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "connected":
			out.Connected = bool(in.Bool())
		case "error":
			out.Error = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5839400EncodeGithubComKhliengDispatchStorage(out *jwriter.Writer, in Network) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if in.Host != "" {
		const prefix string = ",\"host\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Host))
	}
	if in.Port != "" {
		const prefix string = ",\"port\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Port))
	}
	if in.TLS {
		const prefix string = ",\"tls\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.TLS))
	}
	if in.ServerPassword != "" {
		const prefix string = ",\"serverPassword\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ServerPassword))
	}
	if in.Nick != "" {
		const prefix string = ",\"nick\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Nick))
	}
	if in.Username != "" {
		const prefix string = ",\"username\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Username))
	}
	if in.Realname != "" {
		const prefix string = ",\"realname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Realname))
	}
	if in.Account != "" {
		const prefix string = ",\"account\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Account))
	}
	if in.Password != "" {
		const prefix string = ",\"password\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Password))
	}
	if len(in.Features) != 0 {
		const prefix string = ",\"features\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.Features {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v2Name))
				out.RawByte(':')
				if m, ok := v2Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v2Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v2Value))
				}
			}
			out.RawByte('}')
		}
	}
	if in.Connected {
		const prefix string = ",\"connected\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Connected))
	}
	if in.Error != "" {
		const prefix string = ",\"error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Error))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Network) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5839400EncodeGithubComKhliengDispatchStorage(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Network) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5839400EncodeGithubComKhliengDispatchStorage(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Network) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5839400DecodeGithubComKhliengDispatchStorage(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Network) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5839400DecodeGithubComKhliengDispatchStorage(l, v)
}
func easyjsonC5839400DecodeGithubComKhliengDispatchStorage1(in *jlexer.Lexer, out *Channel) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "network":
			out.Network = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "topic":
			out.Topic = string(in.String())
		case "joined":
			out.Joined = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5839400EncodeGithubComKhliengDispatchStorage1(out *jwriter.Writer, in Channel) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Network != "" {
		const prefix string = ",\"network\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Network))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.Topic != "" {
		const prefix string = ",\"topic\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Topic))
	}
	if in.Joined {
		const prefix string = ",\"joined\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Joined))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Channel) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5839400EncodeGithubComKhliengDispatchStorage1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Channel) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5839400EncodeGithubComKhliengDispatchStorage1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Channel) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5839400DecodeGithubComKhliengDispatchStorage1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Channel) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5839400DecodeGithubComKhliengDispatchStorage1(l, v)
}
