package json

import (
	"encoding/json"
	"io"
)

func Decode(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func NewDecoder(r io.Reader) *json.Decoder {
	return json.NewDecoder(r)
}
