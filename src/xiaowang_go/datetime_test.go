package xiaowang_go

import (
	"testing"
	"time"
)

const test_dt datetime = "20160101 08:07:33"
const start_dt datetime = "20160101 00:07:33"
const end_dt datetime = "20160102 08:07:33"

var testConf Transconf = Transconf{rangeFrom: string(start_dt), rangeTo: string(end_dt)}

func TestDatetimeTransform(t *testing.T) {
	t1 := test_dt.Transform(NewRandSeed(1), &testConf)
	t2 := test_dt.Transform(NewRandSeed(2), &testConf)

	if t1 == t2 {
		t.Error("transform failed", t1)
	}

}

func TestDatetimeParse(t *testing.T) {
	if time.Date(2016, 1, 1, 8, 7, 33, 0, time.UTC) != test_dt.Parse() {
		t.Error(test_dt.Parse())
	}
}

func TestDatetimeRangecheck_early(t *testing.T) {
	const early_dt datetime = "20060102 15:04:05"
	if early_dt.rangeok(start_dt, end_dt) {
		t.Error("range should not be ok")
	}
	return
}

func TestDatetimeRangeOK_late(t *testing.T) {
	const late_dt datetime = "20190102 15:04:05"
	if late_dt.rangeok(start_dt, end_dt) {
		t.Error("range should not be ok")
	}
	return
}

//todo: test invalid datetime
