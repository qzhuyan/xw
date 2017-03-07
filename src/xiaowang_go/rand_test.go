package xiaowang_go

import (
	"testing"
)

func TestRandInRange(t *testing.T) {
	gen := NewRandSeed(1)
	randnum := RandInRange(gen, 1, 10)
	if randnum > 10 || randnum < 1 {
		t.Error("rand num out of range: ", randnum)
	}
}
