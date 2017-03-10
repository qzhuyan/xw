package numeric

import (
	//	"fmt"
	"math"
	"math/rand"
	"strconv"
	"xiaowang_go/conf"
	"xiaowang_go/xwrand"
)

type Numeric string

func (n Numeric) Transform(r *rand.Rand, conf *conf.F_spec) string {
	if len(n) <= 9 {
		max := int64(math.Pow(10, float64(n.Len())))
		min := int64(math.Pow(10, float64(n.Len()-1)))
		res := strconv.FormatInt(xwrand.RandInRange(r, min, max), 10)
		return res
	} else {
		return NewNumeric(string(n)[0:8]).Transform(r, conf) + NewNumeric(string(n)[8:]).Transform(r, conf)
	}
}

func (n Numeric) Len() int {
	return len(n)
}

func NewNumeric(s string) Numeric {
	return Numeric(s)
}
