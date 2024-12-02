package parse

import (
	"encoding/json"
	"io"

	"github.com/sondelll/rodeo/internal/rodeo"
)

func ParseFromReader(r io.Reader) (rodeo.UnstructuredObject, error) {
	data := rodeo.UnstructuredObject{}
	dec := json.NewDecoder(r)
	err := dec.Decode(&data)
	return data, err
}
