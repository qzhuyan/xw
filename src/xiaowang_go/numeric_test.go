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
	Res := Number.Transform(NewRandSeed(1), new(Transconf))
	if 6 != Res.Len() {
		t.Error("length is wrong, should be the same", Res)
	}

	if Res == Number {
		t.Error("same number, not transformed?")
	}
}
