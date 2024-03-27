package types

import "encoding/json"

func FromJson[From *string | *[]byte, To any](data From, to To) (e error) {
	if v, ok := any(data).(*[]byte); ok {
		e = json.Unmarshal(*v, to)
	} else if v, ok := any(data).(*string); ok {
		e = json.Unmarshal([]byte(*v), to)
	}
	return
}

func Json[From any, To *string | *[]byte](from From, to To) (e error) {
	var b []byte
	if v, ok := any(to).(*[]byte); ok {
		*v, e = json.Marshal(from)
	} else if v, ok := any(to).(*string); ok {
		b, e = json.Marshal(from)
		*v = string(b)
	}
	return
}
