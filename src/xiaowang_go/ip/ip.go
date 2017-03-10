package ipaddr

import (
	//	"fmt"
	"math/rand"
	"strings"
	"xiaowang_go/conf"
	"xiaowang_go/xwrand"
)

type Ipaddr string

func (ip Ipaddr) Transform(r *rand.Rand, conf *conf.F_spec) string {
	segs := strings.Split(string(ip), ".")
	for loc, _ := range segs {
		segs[loc] = string(xwrand.RandInRange(r, int64(0), int64(255)))
	}
	return strings.Join(segs, ".")
}
