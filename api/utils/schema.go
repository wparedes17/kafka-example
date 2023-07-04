package utils

import "github.com/gorilla/schema"

var decoder *schema.Decoder

func init() {
	decoder = schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
}

func DecodeSchema(dst interface{}, src map[string][]string) error {
	return decoder.Decode(dst, src)
}
