package json

import (
	"encoding/json"
	"io"
)

func Encode(v any) ([]byte, error) {
	return json.Marshal(v)
}

func NewEncoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}
