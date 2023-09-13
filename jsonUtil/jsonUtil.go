package jsonUtil

import (
	"bytes"
	"encoding/json"
)

func UnmarshalUseNumber(data []byte, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewBuffer(data))
	decoder.UseNumber()
	return decoder.Decode(v)
}
