package ipaddr_test

import (
	"testing"
	"xiaowang_go/conf"
	"xiaowang_go/ip"
	"xiaowang_go/xwrand"
)

func TestIPTransform(t *testing.T) {
	var testip ipaddr.Ipaddr = "1.1.1.1"
	res := testip.Transform(xwrand.NewRandSeed(1), new(conf.F_spec))
	if res == string(testip) {
		t.Error("transform failed!")
	}
	//todo: checkmore about range

}
