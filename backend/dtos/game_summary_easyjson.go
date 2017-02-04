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

func easyjson861ee97bDecodeGithubComSkeswaEnbiyayBackendDtos(in *jlexer.Lexer, out *GameSummary) {
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
		case "1.1":
			out.ID = string(in.String())
		case "1.2":
			(out.LiveGameStats).UnmarshalEasyJSON(in)
		case "1.3":
			out.GameStartTime = int(in.Int())
		case "1.4":
			out.GameStartTimeTBD = bool(in.Bool())
		case "2.1":
			out.HomeTeamWins = int(in.Int())
		case "2.2":
			out.HomeTeamScore = int(in.Int())
		case "2.3":
			out.HomeTeamLosses = int(in.Int())
		case "2.4":
			out.HomeTeamTeamID = string(in.String())
		case "2.5":
			out.HomeTeamTriCode = string(in.String())
		case "3.1":
			out.AwayTeamWins = int(in.Int())
		case "3.2":
			out.AwayTeamScore = int(in.Int())
		case "3.3":
			out.AwayTeamLosses = int(in.Int())
		case "3.4":
			out.AwayTeamTeamID = string(in.String())
		case "3.5":
			out.AwayTeamTriCode = string(in.String())
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
func easyjson861ee97bEncodeGithubComSkeswaEnbiyayBackendDtos(out *jwriter.Writer, in GameSummary) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"1.1\":")
	out.String(string(in.ID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"1.2\":")
	(in.LiveGameStats).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"1.3\":")
	out.Int(int(in.GameStartTime))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"1.4\":")
	out.Bool(bool(in.GameStartTimeTBD))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"2.1\":")
	out.Int(int(in.HomeTeamWins))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"2.2\":")
	out.Int(int(in.HomeTeamScore))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"2.3\":")
	out.Int(int(in.HomeTeamLosses))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"2.4\":")
	out.String(string(in.HomeTeamTeamID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"2.5\":")
	out.String(string(in.HomeTeamTriCode))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"3.1\":")
	out.Int(int(in.AwayTeamWins))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"3.2\":")
	out.Int(int(in.AwayTeamScore))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"3.3\":")
	out.Int(int(in.AwayTeamLosses))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"3.4\":")
	out.String(string(in.AwayTeamTeamID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"3.5\":")
	out.String(string(in.AwayTeamTriCode))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GameSummary) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson861ee97bEncodeGithubComSkeswaEnbiyayBackendDtos(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GameSummary) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson861ee97bEncodeGithubComSkeswaEnbiyayBackendDtos(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GameSummary) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson861ee97bDecodeGithubComSkeswaEnbiyayBackendDtos(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GameSummary) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson861ee97bDecodeGithubComSkeswaEnbiyayBackendDtos(l, v)
}
