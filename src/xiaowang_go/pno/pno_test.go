package pno_test

import (
	"testing"
	"xiaowang_go/conf"
	"xiaowang_go/pno"
	"xiaowang_go/xwrand"
)

func TestPnoTransform(t *testing.T) {
	res := pno.Transform("SE19890807-3875", xwrand.NewRandSeed(1), new(conf.F_spec))
	if res != "SE6551919-1410" {
		t.Error("transform failed", res)
	}
}
