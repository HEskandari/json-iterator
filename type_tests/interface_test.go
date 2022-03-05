package test

import (
	jsoniter "github.com/heskandari/json-iterator"
	"testing"
)

type StringerType int
type StructWithStringer struct {
	MyVal StringerType
}

func (t StringerType) String() string {
	switch t {
	case 1:
		return "True"
	case 0:
		return "False"
	}
	return ""
}

func TestStringer(t *testing.T) {
	cfg := jsoniter.Config{}.Froze()

	st := StructWithStringer{
		MyVal: StringerType(1),
	}
	b, err := cfg.Marshal(st)
	js := string(b)

	if err != nil {
		t.Fatalf("failed to marshal with jsoniter: %v", err)
	}
	if js != "{\"MyVal\":\"True\"}" {
		t.Fatalf("failed to marshal Stringer with jsoniter: %v", err)
	}
}