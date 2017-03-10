package datetime_test

import (
	"testing"
	"time"
	"xiaowang_go/conf"
	"xiaowang_go/datetime"
	"xiaowang_go/xwrand"
)

const test_dt datetime.Datetime = "2016-07-10 11:00:20.000"
const start_dt datetime.Datetime = "2016-01-01 00:07:33.000"
const end_dt datetime.Datetime = "2016-01-02 08:07:33.000"

var testConf conf.F_spec = conf.F_spec{Rangefrom: string(start_dt), Rangeto: string(end_dt)}

func TestDatetimeTransform(t *testing.T) {
	t1 := test_dt.Transform(xwrand.NewRandSeed(1), &testConf)
	t2 := test_dt.Transform(xwrand.NewRandSeed(2), &testConf)

	if t1 == t2 {
		t.Error("transform failed", t1)
	}
}

func TestDatetimeParse(t *testing.T) {
	if time.Date(2016, 7, 10, 11, 00, 20, 0, time.UTC) != test_dt.Parse() {
		t.Error(test_dt.Parse())
	}
}

func TestDatetimeRangecheck_early(t *testing.T) {
	const early_dt datetime.Datetime = "2006-01-02 15:04:05"
	if early_dt.Rangeok(start_dt, end_dt) {
		t.Error("range should not be ok")
	}
	return
}

func TestDatetimeRangeOK_late(t *testing.T) {
	const late_dt datetime.Datetime = "2019-01-02 15:04:05"
	if late_dt.Rangeok(start_dt, end_dt) {
		t.Error("range should not be ok")
	}
	return
}

//todo: test invalid datetime
