package numeric

import (
	"math"
	"math/rand"
	"strconv"
	"xiaowang_go/conf"
	"xiaowang_go/xwrand"
)

type Numeric int64

func (n Numeric) Transform(r *rand.Rand, _conf *conf.F_spec) string {
	max := int64(math.Pow(10, float64(n.Len())))
	min := int64(math.Pow(10, float64(n.Len()-1)))
	return strconv.FormatInt(min+xwrand.RandInRange(r, min, max), 10)
}

func (n Numeric) Len() int {
	return len(strconv.FormatInt(int64(n), 10))
}

func NewNumeric(s string) Numeric {
	i, error := strconv.ParseInt(s, 10, 0)
	if error != nil {
		panic(error)
	}
	return Numeric(i)
}
