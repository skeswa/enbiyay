// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package dtos

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

func easyjsonDabaa7ccDecodeGithubComSkeswaEnbiyayBackendDtos(in *jlexer.Lexer, out *GameLeader) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "1":
			out.ID = string(in.String())
		case "2":
			out.Name = string(in.String())
		case "3":
			out.Minutes = string(in.String())
		case "4":
			out.StatValue = string(in.String())
		case "5":
			out.JerseyNumber = string(in.String())
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
func easyjsonDabaa7ccEncodeGithubComSkeswaEnbiyayBackendDtos(out *jwriter.Writer, in GameLeader) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"1\":")
	out.String(string(in.ID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"2\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"3\":")
	out.String(string(in.Minutes))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"4\":")
	out.String(string(in.StatValue))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"5\":")
	out.String(string(in.JerseyNumber))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GameLeader) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDabaa7ccEncodeGithubComSkeswaEnbiyayBackendDtos(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GameLeader) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDabaa7ccEncodeGithubComSkeswaEnbiyayBackendDtos(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GameLeader) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDabaa7ccDecodeGithubComSkeswaEnbiyayBackendDtos(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GameLeader) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDabaa7ccDecodeGithubComSkeswaEnbiyayBackendDtos(l, v)
}