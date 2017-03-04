package xiaowang_go

import "testing"

const test_dt datetime = "20160101 08:07:33"

func TestDatetimeTransform(t *testing.T) {
	t1 := test_dt.Transform(1)
	t2 := test_dt.Transform(2)
	if t1 == t2 {
		t.Error("transform failed", t1)
	}
}
