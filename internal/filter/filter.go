package filter

import (
	"strings"

	"github.com/sondelll/rodeo/internal/rodeo"
)

func FilterByKeys(d rodeo.UnstructuredObject, k []string) rodeo.UnstructuredObject {
	out := rodeo.UnstructuredObject{}

	for dataKey, val := range d {
		for _, filterKey := range k {
			if strings.EqualFold(dataKey, filterKey) {
				out[dataKey] = val
			}
		}
	}
	return out
}

func ParseFilterKeys(s string) []string {
	s = strings.ReplaceAll(s, "\"", "")
	return strings.Split(s, ",")
}
