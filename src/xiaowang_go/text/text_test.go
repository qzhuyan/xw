package text_test

import "testing"
import "strings"
import "xiaowang_go/text"
import "xiaowang_go/xwrand"
import "xiaowang_go/conf"

const test_string text.Text = "abc efg"

func TestTextTransform(t *testing.T) {
	res := test_string.Transform(xwrand.NewRandSeed(1), new(conf.F_spec))
	words := strings.Split(res, " ")
	if len(words[0]) != 3 {
		t.Error("len should not be changed")
	}

	if len(words[1]) != 3 {
		t.Error("len should not be changed")
	}

	res2 := test_string.Transform(xwrand.NewRandSeed(1), new(conf.F_spec))

	if res2 != res {
		t.Error("should have same result with same seed")
	}

}

//Benchmark which is faster?

func BenchmarkTransformWithNewSeed(b *testing.B) {
	for i := 0; i < 10000; i++ {
		test_string.Transform(xwrand.NewRandSeed(1), new(conf.F_spec))
	}
}

func BenchmarkTransformWithRandSrc(b *testing.B) {
	r := xwrand.NewRandSeed(1)
	for i := 0; i < 10000; i++ {
		test_string.Transform(r, new(conf.F_spec))
	}
}
