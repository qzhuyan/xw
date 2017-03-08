package xiaowang_go

import (
	"testing"
)

var Number numeric = 123456
var NumericLen = 6

func TestLen(t *testing.T) {
	if 6 != Number.Len() {
		t.Error("failed")
	}
}

func TestTransform(t *testing.T) {
	Res := Number.Transform(NewRandSeed(1), new(f_spec))
	if 6 != len(Res) {
		t.Error("length is wrong, should be the same", Res)
	}
}
