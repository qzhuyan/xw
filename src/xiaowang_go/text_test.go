package xiaowang_go

import "testing"
import "strings"

const test_string text = "abc efg"

func TestTextTransform(t *testing.T) {
	res := test_string.Transform(1, new(Transconf))
	words := strings.Split(res, " ")
	if len(words[0]) != 3 {
		t.Error("len should not be changed")
	}

	if len(words[1]) != 3 {
		t.Error("len should not be changed")
	}
}
