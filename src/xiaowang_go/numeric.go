package xiaowang_go

import (
	"math"
	"math/rand"
	"strconv"
)

type numeric int64

func (n numeric) Transform(r *rand.Rand, _conf *f_spec) string {
	max := int64(math.Pow(10, float64(n.Len())))
	min := int64(math.Pow(10, float64(n.Len()-1)))
	return strconv.FormatInt(min+RandInRange(r, min, max), 10)
}

func (n numeric) Len() int {
	return len(strconv.FormatInt(int64(n), 10))
}

func NewNumeric(s string) numeric {
	i, error := strconv.ParseInt(s, 10, 0)
	if error != nil {
		panic(error)
	}
	return numeric(i)
}
