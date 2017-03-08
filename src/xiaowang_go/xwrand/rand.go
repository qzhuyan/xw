package xwrand

import (
	"math/rand"
)

func RandInRange(s *rand.Rand, from, to int64) int64 {
	diff := to - from
	return from + (*s).Int63n(diff)
}

func NewRandSeed(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}
