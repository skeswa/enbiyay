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

func easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos(in *jlexer.Lexer, out *NBAGameBroadcaster) {
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
		case "longName":
			out.LongName = string(in.String())
		case "shortName":
			out.ShortName = string(in.String())
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
func easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos(out *jwriter.Writer, in NBAGameBroadcaster) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"longName\":")
	out.String(string(in.LongName))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"shortName\":")
	out.String(string(in.ShortName))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NBAGameBroadcaster) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NBAGameBroadcaster) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NBAGameBroadcaster) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NBAGameBroadcaster) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos(l, v)
}
func easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos1(in *jlexer.Lexer, out *NBAGameDuration) {
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
		case "hours":
			out.Hours = string(in.String())
		case "minute":
			out.Minute = string(in.String())
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
func easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos1(out *jwriter.Writer, in NBAGameDuration) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"hours\":")
	out.String(string(in.Hours))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"minute\":")
	out.String(string(in.Minute))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NBAGameDuration) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NBAGameDuration) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NBAGameDuration) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NBAGameDuration) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos1(l, v)
}
func easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos2(in *jlexer.Lexer, out *NBAGamePeriod) {
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
		case "type":
			out.Type = int(in.Int())
		case "current":
			out.Current = int(in.Int())
		case "maxRegular":
			out.MaxRegular = int(in.Int())
		case "isHalftime":
			out.IsHalftime = bool(in.Bool())
		case "isEndOfPeriod":
			out.IsEndOfPeriod = bool(in.Bool())
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
func easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos2(out *jwriter.Writer, in NBAGamePeriod) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"type\":")
	out.Int(int(in.Type))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"current\":")
	out.Int(int(in.Current))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"maxRegular\":")
	out.Int(int(in.MaxRegular))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"isHalftime\":")
	out.Bool(bool(in.IsHalftime))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"isEndOfPeriod\":")
	out.Bool(bool(in.IsEndOfPeriod))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NBAGamePeriod) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NBAGamePeriod) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NBAGamePeriod) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NBAGamePeriod) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos2(l, v)
}
func easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos3(in *jlexer.Lexer, out *NBAGameDeepLink) {
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
		case "iosApp":
			out.IOS = string(in.String())
		case "mobileApp":
			out.Mobile = string(in.String())
		case "androidApp":
			out.Android = string(in.String())
		case "desktopWeb":
			out.Desktop = string(in.String())
		case "broadcaster":
			out.Broadcaster = string(in.String())
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
func easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos3(out *jwriter.Writer, in NBAGameDeepLink) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"iosApp\":")
	out.String(string(in.IOS))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mobileApp\":")
	out.String(string(in.Mobile))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"androidApp\":")
	out.String(string(in.Android))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"desktopWeb\":")
	out.String(string(in.Desktop))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"broadcaster\":")
	out.String(string(in.Broadcaster))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NBAGameDeepLink) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NBAGameDeepLink) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NBAGameDeepLink) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NBAGameDeepLink) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos3(l, v)
}
func easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos4(in *jlexer.Lexer, out *NBAGameStream) {
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
		case "streamId":
			out.ID = string(in.String())
		case "streamType":
			out.Type = string(in.String())
		case "isOnAir":
			out.IsOnAir = bool(in.Bool())
		case "duration":
			out.Duration = int(in.Int())
		case "isArchiveAvailToWatch":
			out.ArchiveAvailable = bool(in.Bool())
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
func easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos4(out *jwriter.Writer, in NBAGameStream) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"streamId\":")
	out.String(string(in.ID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"streamType\":")
	out.String(string(in.Type))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"isOnAir\":")
	out.Bool(bool(in.IsOnAir))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"duration\":")
	out.Int(int(in.Duration))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"isArchiveAvailToWatch\":")
	out.Bool(bool(in.ArchiveAvailable))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NBAGameStream) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NBAGameStream) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NBAGameStream) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NBAGameStream) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos4(l, v)
}
func easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos5(in *jlexer.Lexer, out *NBAGameVideoDetails) {
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
		case "canPurchase":
			out.CanPurchase = bool(in.Bool())
		case "isLeaguePass":
			out.IsOnLeaguePass = bool(in.Bool())
		case "isTNTOT":
			out.IsOnTNT = bool(in.Bool())
		case "streams":
			if in.IsNull() {
				in.Skip()
				out.Streams = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.Streams = make([]NBAGameStream, 0, 1)
				} else {
					out.Streams = []NBAGameStream{}
				}
				for !in.IsDelim(']') {
					var v1 NBAGameStream
					(v1).UnmarshalEasyJSON(in)
					out.Streams = append(out.Streams, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "deepLink":
			if in.IsNull() {
				in.Skip()
				out.DeepLinks = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.DeepLinks = make([]NBAGameDeepLink, 0, 1)
				} else {
					out.DeepLinks = []NBAGameDeepLink{}
				}
				for !in.IsDelim(']') {
					var v2 NBAGameDeepLink
					(v2).UnmarshalEasyJSON(in)
					out.DeepLinks = append(out.DeepLinks, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos5(out *jwriter.Writer, in NBAGameVideoDetails) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"canPurchase\":")
	out.Bool(bool(in.CanPurchase))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"isLeaguePass\":")
	out.Bool(bool(in.IsOnLeaguePass))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"isTNTOT\":")
	out.Bool(bool(in.IsOnTNT))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"streams\":")
	if in.Streams == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v3, v4 := range in.Streams {
			if v3 > 0 {
				out.RawByte(',')
			}
			(v4).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"deepLink\":")
	if in.DeepLinks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in.DeepLinks {
			if v5 > 0 {
				out.RawByte(',')
			}
			(v6).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NBAGameVideoDetails) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NBAGameVideoDetails) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NBAGameVideoDetails) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NBAGameVideoDetails) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos5(l, v)
}
func easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos6(in *jlexer.Lexer, out *NBAGameBroadcasters) {
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
		case "national":
			if in.IsNull() {
				in.Skip()
				out.National = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.National = make([]NBAGameBroadcaster, 0, 2)
				} else {
					out.National = []NBAGameBroadcaster{}
				}
				for !in.IsDelim(']') {
					var v7 NBAGameBroadcaster
					(v7).UnmarshalEasyJSON(in)
					out.National = append(out.National, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "canadian":
			if in.IsNull() {
				in.Skip()
				out.Canadian = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.Canadian = make([]NBAGameBroadcaster, 0, 2)
				} else {
					out.Canadian = []NBAGameBroadcaster{}
				}
				for !in.IsDelim(']') {
					var v8 NBAGameBroadcaster
					(v8).UnmarshalEasyJSON(in)
					out.Canadian = append(out.Canadian, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "hTeam":
			if in.IsNull() {
				in.Skip()
				out.LocalHome = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.LocalHome = make([]NBAGameBroadcaster, 0, 2)
				} else {
					out.LocalHome = []NBAGameBroadcaster{}
				}
				for !in.IsDelim(']') {
					var v9 NBAGameBroadcaster
					(v9).UnmarshalEasyJSON(in)
					out.LocalHome = append(out.LocalHome, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "vTeam":
			if in.IsNull() {
				in.Skip()
				out.LocalAway = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.LocalAway = make([]NBAGameBroadcaster, 0, 2)
				} else {
					out.LocalAway = []NBAGameBroadcaster{}
				}
				for !in.IsDelim(']') {
					var v10 NBAGameBroadcaster
					(v10).UnmarshalEasyJSON(in)
					out.LocalAway = append(out.LocalAway, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos6(out *jwriter.Writer, in NBAGameBroadcasters) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"national\":")
	if in.National == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v11, v12 := range in.National {
			if v11 > 0 {
				out.RawByte(',')
			}
			(v12).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"canadian\":")
	if in.Canadian == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v13, v14 := range in.Canadian {
			if v13 > 0 {
				out.RawByte(',')
			}
			(v14).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"hTeam\":")
	if in.LocalHome == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v15, v16 := range in.LocalHome {
			if v15 > 0 {
				out.RawByte(',')
			}
			(v16).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"vTeam\":")
	if in.LocalAway == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v17, v18 := range in.LocalAway {
			if v17 > 0 {
				out.RawByte(',')
			}
			(v18).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NBAGameBroadcasters) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NBAGameBroadcasters) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NBAGameBroadcasters) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NBAGameBroadcasters) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos6(l, v)
}
func easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos7(in *jlexer.Lexer, out *NBAGameBroadcast) {
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
		case "broadcasters":
			(out.Broadcasters).UnmarshalEasyJSON(in)
		case "video":
			(out.Details).UnmarshalEasyJSON(in)
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
func easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos7(out *jwriter.Writer, in NBAGameBroadcast) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"broadcasters\":")
	(in.Broadcasters).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"video\":")
	(in.Details).MarshalEasyJSON(out)
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NBAGameBroadcast) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NBAGameBroadcast) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NBAGameBroadcast) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NBAGameBroadcast) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos7(l, v)
}
func easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos8(in *jlexer.Lexer, out *NBAGameVideoMetadata) {
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
		case "broadcast":
			(out.Broadcast).UnmarshalEasyJSON(in)
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
func easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos8(out *jwriter.Writer, in NBAGameVideoMetadata) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"broadcast\":")
	(in.Broadcast).MarshalEasyJSON(out)
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NBAGameVideoMetadata) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NBAGameVideoMetadata) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NBAGameVideoMetadata) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NBAGameVideoMetadata) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos8(l, v)
}
func easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos9(in *jlexer.Lexer, out *NBAGame) {
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
		case "gameId":
			out.ID = string(in.String())
		case "startTimeUTC":
			out.StartTime = string(in.String())
		case "isStartTimeTBD":
			out.IsStartTimeTBD = bool(in.Bool())
		case "gameDuration":
			(out.Duration).UnmarshalEasyJSON(in)
		case "period":
			(out.Period).UnmarshalEasyJSON(in)
		case "hTeam":
			easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos10(in, &out.HomeTeam)
		case "vTeam":
			easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos10(in, &out.AwayTeam)
		case "watch":
			(out.VideoMetadata).UnmarshalEasyJSON(in)
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
func easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos9(out *jwriter.Writer, in NBAGame) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"gameId\":")
	out.String(string(in.ID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"startTimeUTC\":")
	out.String(string(in.StartTime))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"isStartTimeTBD\":")
	out.Bool(bool(in.IsStartTimeTBD))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"gameDuration\":")
	(in.Duration).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"period\":")
	(in.Period).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"hTeam\":")
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos10(out, in.HomeTeam)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"vTeam\":")
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos10(out, in.AwayTeam)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"watch\":")
	(in.VideoMetadata).MarshalEasyJSON(out)
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NBAGame) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NBAGame) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NBAGame) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NBAGame) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos9(l, v)
}
func easyjson4bb17c5aDecodeGithubComSkeswaEnbiyayBackendNbaDtos10(in *jlexer.Lexer, out *NBATeam) {
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
		case "teamId":
			out.ID = string(in.String())
		case "win":
			out.Win = string(in.String())
		case "loss":
			out.Loss = string(in.String())
		case "score":
			out.Score = string(in.String())
		case "triCode":
			out.Tricode = string(in.String())
		case "seriesWin":
			out.SeriesWin = string(in.String())
		case "seriesLoss":
			out.SeriesLoss = string(in.String())
		case "linescore":
			if in.IsNull() {
				in.Skip()
				out.PeriodScores = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.PeriodScores = make([]struct {
						Score string "json:\"score\""
					}, 0, 4)
				} else {
					out.PeriodScores = []struct {
						Score string "json:\"score\""
					}{}
				}
				for !in.IsDelim(']') {
					var v19 struct {
						Score string "json:\"score\""
					}
					easyjson4bb17c5aDecode(in, &v19)
					out.PeriodScores = append(out.PeriodScores, v19)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson4bb17c5aEncodeGithubComSkeswaEnbiyayBackendNbaDtos10(out *jwriter.Writer, in NBATeam) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"teamId\":")
	out.String(string(in.ID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"win\":")
	out.String(string(in.Win))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"loss\":")
	out.String(string(in.Loss))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"score\":")
	out.String(string(in.Score))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"triCode\":")
	out.String(string(in.Tricode))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"seriesWin\":")
	out.String(string(in.SeriesWin))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"seriesLoss\":")
	out.String(string(in.SeriesLoss))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"linescore\":")
	if in.PeriodScores == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v20, v21 := range in.PeriodScores {
			if v20 > 0 {
				out.RawByte(',')
			}
			easyjson4bb17c5aEncode(out, v21)
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}
func easyjson4bb17c5aDecode(in *jlexer.Lexer, out *struct {
	Score string "json:\"score\""
}) {
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
		case "score":
			out.Score = string(in.String())
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
func easyjson4bb17c5aEncode(out *jwriter.Writer, in struct {
	Score string "json:\"score\""
}) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"score\":")
	out.String(string(in.Score))
	out.RawByte('}')
}