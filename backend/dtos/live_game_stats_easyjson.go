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

func easyjsonEa1a6751DecodeGithubComSkeswaEnbiyayBackendDtos(in *jlexer.Lexer, out *LiveGameStats) {
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
			out.Period = int(in.Int())
		case "2":
			out.Channel = string(in.String())
		case "3":
			out.TimeRemaining = string(in.String())
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
func easyjsonEa1a6751EncodeGithubComSkeswaEnbiyayBackendDtos(out *jwriter.Writer, in LiveGameStats) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"1\":")
	out.Int(int(in.Period))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"2\":")
	out.String(string(in.Channel))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"3\":")
	out.String(string(in.TimeRemaining))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LiveGameStats) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEa1a6751EncodeGithubComSkeswaEnbiyayBackendDtos(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LiveGameStats) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEa1a6751EncodeGithubComSkeswaEnbiyayBackendDtos(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LiveGameStats) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEa1a6751DecodeGithubComSkeswaEnbiyayBackendDtos(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LiveGameStats) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEa1a6751DecodeGithubComSkeswaEnbiyayBackendDtos(l, v)
}
