// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjson11d1a9baDecode20211NoskoolTeamInternalAppProfilesModels(in *jlexer.Lexer, out *RequestLogin) {
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
		case "nickname":
			out.Login = string(in.String())
		case "password":
			out.Password = string(in.String())
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
func easyjson11d1a9baEncode20211NoskoolTeamInternalAppProfilesModels(out *jwriter.Writer, in RequestLogin) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nickname\":"
		out.RawString(prefix[1:])
		out.String(string(in.Login))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestLogin) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncode20211NoskoolTeamInternalAppProfilesModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestLogin) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncode20211NoskoolTeamInternalAppProfilesModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestLogin) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecode20211NoskoolTeamInternalAppProfilesModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestLogin) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecode20211NoskoolTeamInternalAppProfilesModels(l, v)
}
func easyjson11d1a9baDecode20211NoskoolTeamInternalAppProfilesModels1(in *jlexer.Lexer, out *ProfileRequest) {
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
		case "email":
			out.Email = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "nickname":
			out.Nickname = string(in.String())
		case "first_name":
			out.Name = string(in.String())
		case "second_name":
			out.Surname = string(in.String())
		case "favorite_genre":
			if in.IsNull() {
				in.Skip()
				out.FavoriteGenre = nil
			} else {
				in.Delim('[')
				if out.FavoriteGenre == nil {
					if !in.IsDelim(']') {
						out.FavoriteGenre = make([]string, 0, 4)
					} else {
						out.FavoriteGenre = []string{}
					}
				} else {
					out.FavoriteGenre = (out.FavoriteGenre)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.FavoriteGenre = append(out.FavoriteGenre, v1)
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
func easyjson11d1a9baEncode20211NoskoolTeamInternalAppProfilesModels1(out *jwriter.Writer, in ProfileRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix[1:])
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	{
		const prefix string = ",\"nickname\":"
		out.RawString(prefix)
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"first_name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"second_name\":"
		out.RawString(prefix)
		out.String(string(in.Surname))
	}
	{
		const prefix string = ",\"favorite_genre\":"
		out.RawString(prefix)
		if in.FavoriteGenre == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.FavoriteGenre {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ProfileRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson11d1a9baEncode20211NoskoolTeamInternalAppProfilesModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ProfileRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson11d1a9baEncode20211NoskoolTeamInternalAppProfilesModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProfileRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson11d1a9baDecode20211NoskoolTeamInternalAppProfilesModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ProfileRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson11d1a9baDecode20211NoskoolTeamInternalAppProfilesModels1(l, v)
}
