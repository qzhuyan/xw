package xiaowang_go

import (
	"math"
	"math/rand"
	"strconv"
)

type numeric uint64

func (n numeric) Transform(seed int64) numeric {
	gen := rand.New(rand.NewSource(seed))
	max := math.Pow(10, float64(n.Len()))
	min := math.Pow(10, float64(n.Len()-1))
	return numeric(min) + numeric(gen.Int63n(int64(max-min)))
}

func (n numeric) Len() int {
	return len(strconv.FormatUint(uint64(n), 10))
}
