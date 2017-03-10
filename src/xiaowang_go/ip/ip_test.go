package ipaddr_test

import (
	"strconv"
	"strings"
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

	segs := strings.Split(res, ".")
	for _, seg := range segs {
		i, err := strconv.Atoi(seg)
		if err != nil {
			t.Error("cannot convert to int")
		}
		if !(i < 255 && i > 0) {
			t.Error("range error")
		}

	}

}
