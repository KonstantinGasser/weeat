package json

import (
	"encoding/json"
	"io"
)

func FromRequest(body io.ReadCloser, v interface{}) error {
	return json.NewDecoder(body).Decode(v)
}
