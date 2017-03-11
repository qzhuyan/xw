package pno

import (
	"fmt"
	"math/rand"
	"strings"
	"xiaowang_go/conf"
	"xiaowang_go/datetime"
	"xiaowang_go/numeric"
)

func Transform(p string, r *rand.Rand, conf *conf.F_spec) string {
	pstr := string(p)
	if "SE" == pstr[0:2] {
		//sweden
		segs := strings.Split(pstr[2:], "-")
		birth := segs[0]
		last4 := numeric.Numeric(segs[1]).Transform(r, conf)
		conf.FormatIn = "20060102"
		birth = datetime.Transform(birth, r, conf)
		/*
			year := numeric.Numeric(birth[0:4]).Transform(r, conf)
			month := numeric.Numeric(birth[5:7]).Transform(r, conf)
			day := numeric.Numeric(birth[7:]).Transform(r, conf)
		*/

		return fmt.Sprintf("SE%s-%s", birth, last4)
	} else {
		panic("not supported country:" + pstr[0:2])
	}
}
