package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/sondelll/rodeo/internal/filter"
	"github.com/sondelll/rodeo/internal/parse"
)

func main() {
	var kFil string = ""
	var usePretty bool = false
	flag.BoolVar(&usePretty, "p", false, "pretty print mode")
	flag.StringVar(&kFil, "f", "all", "key filter")
	flag.Parse()

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	b := sc.Bytes()

	data, err := parse.ParseFromReader(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	if kFil == "all" {
		fmt.Printf("%s\n", asJson(data, usePretty))
	} else {
		kf := filter.ParseFilterKeys(kFil)
		flt := filter.FilterByKeys(data, kf)

		fmt.Printf("%s\n", asJson(flt, usePretty))
	}
}

func asJson(d any, pretty bool) []byte {
	var b []byte
	var err error
	if pretty {
		b, err = json.MarshalIndent(d, "", "  ")
	} else {
		b, err = json.Marshal(d)
	}
	if err != nil {
		return []byte("{\"error\":\"serializeErr\"}")
	}
	return b
}
