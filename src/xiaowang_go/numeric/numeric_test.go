package numeric_test

import (
	"testing"
	"xiaowang_go/conf"
	"xiaowang_go/numeric"
	"xiaowang_go/xwrand"
)

var Number numeric.Numeric = "7207297"

var BigNumber numeric.Numeric = "454794729933338495114986149126587207297"
var NumericLen = 6

func TestLen(t *testing.T) {
	if 7 != Number.Len() {
		t.Error("failed")
	}
}

func TestTransform(t *testing.T) {
	Res := Number.Transform(xwrand.NewRandSeed(1), new(conf.F_spec))
	if 7 != len(Res) {
		t.Error("length is wrong, should be the same", Res)
	}
}

func TestTransformBignum(t *testing.T) {
	Res := BigNumber.Transform(xwrand.NewRandSeed(1), new(conf.F_spec))
	if BigNumber.Len() != len(Res) {
		t.Errorf("length is wrong, should be the same\n%s\n%s", Res, BigNumber)
	}
}
