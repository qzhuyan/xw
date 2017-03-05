package xiaowang_go

import (
	"math"
	"math/rand"
	"strconv"
)

type numeric int64

func (n numeric) Transform(seed int64) numeric {
	gen := rand.New(rand.NewSource(seed))
	max := int64(math.Pow(10, float64(n.Len())))
	min := int64(math.Pow(10, float64(n.Len()-1)))
	return numeric(min + RandInRange(gen, min, max))
}

func (n numeric) Len() int {
	return len(strconv.FormatInt(int64(n), 10))
}
