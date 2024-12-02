package parse

import (
	"bytes"
	"testing"
)

func TestParse(t *testing.T) {
	d := []byte(`{"a":"b"}`)
	rdr := bytes.NewReader(d)
	result, err := ParseFromReader(rdr)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", result)
}
