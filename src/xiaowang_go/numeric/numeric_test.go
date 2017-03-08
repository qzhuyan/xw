package numeric_test

import (
	"testing"
	"xiaowang_go/conf"
	"xiaowang_go/numeric"
	"xiaowang_go/xwrand"
)

var Number numeric.Numeric = 123456
var NumericLen = 6

func TestLen(t *testing.T) {
	if 6 != Number.Len() {
		t.Error("failed")
	}
}

func TestTransform(t *testing.T) {
	Res := Number.Transform(xwrand.NewRandSeed(1), new(conf.F_spec))
	if 6 != len(Res) {
		t.Error("length is wrong, should be the same", Res)
	}
}
