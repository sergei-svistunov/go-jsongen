// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package test

import (
	json "encoding/json"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

var _ = json.RawMessage{} // suppress unused package warning

func easyjson_4abafe4c_decode_github_com_sergei_svistunov_go_jsongen_test_RecursiveStruct(in *jlexer.Lexer, out *RecursiveStruct) {
	if in.IsNull() {
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
		case "IntField":
			out.IntField = int(in.Int())
		case "StrField":
			out.StrField = string(in.String())
		case "RecursiveField":
			if in.IsNull() {
				in.Skip()
				out.RecursiveField = nil
			} else {
				out.RecursiveField = new(RecursiveStruct)
				(*out.RecursiveField).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjson_4abafe4c_encode_github_com_sergei_svistunov_go_jsongen_test_RecursiveStruct(out *jwriter.Writer, in RecursiveStruct) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"IntField\":")
	out.Int(int(in.IntField))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"StrField\":")
	out.String(string(in.StrField))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"RecursiveField\":")
	if in.RecursiveField == nil {
		out.RawString("null")
	} else {
		(*in.RecursiveField).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}
func (v RecursiveStruct) EasyMarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson_4abafe4c_encode_github_com_sergei_svistunov_go_jsongen_test_RecursiveStruct(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}
func (v RecursiveStruct) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson_4abafe4c_encode_github_com_sergei_svistunov_go_jsongen_test_RecursiveStruct(w, v)
}
func (v *RecursiveStruct) EasyUnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson_4abafe4c_decode_github_com_sergei_svistunov_go_jsongen_test_RecursiveStruct(&r, v)
	return r.Error()
}
func (v *RecursiveStruct) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson_4abafe4c_decode_github_com_sergei_svistunov_go_jsongen_test_RecursiveStruct(l, v)
}
