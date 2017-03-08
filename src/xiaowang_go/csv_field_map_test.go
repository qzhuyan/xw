package xiaowang_go

import (
	"testing"
)

func TestParseCSVConf(t *testing.T) {
	conf := parse_csv_conf("testdata/testcfg.yml")

	// internal conf
	fs := make([]f_spec, 2)

	fs[0] = f_spec{
		Name:  "orderid",
		Class: "numeric",
		Conf:  "",
		Pos:   1,
	}
	fs[1] = f_spec{
		Name:  "item",
		Class: "text",
		Conf:  "",
		Pos:   2,
	}

	if conf.Version != 1 {
		t.Error("version is not 1, ", conf.Version)
	}

	if !conf.Hashead {
		t.Error(" hashead is not true, ", conf.Hashead)
	}

	if conf.Fields[0] != fs[0] {
		t.Error("Field_spec compared failed, ", conf.Fields)
	}

	if conf.Fields[1] != fs[1] {
		t.Error("Field_spec compared failed, ", conf.Fields)
	}
}
