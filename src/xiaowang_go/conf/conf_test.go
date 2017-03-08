package conf_test

import (
	"testing"
	"xiaowang_go/conf"
)

func TestParseCSVConf(t *testing.T) {
	c := conf.ParseConf("../testdata/testcfg.yml")

	// internal conf
	fs := make([](conf.F_spec), 2)

	fs[0] = conf.F_spec{
		Name:      "orderid",
		Class:     "numeric",
		Conf:      "",
		Pos:       1,
		Rangefrom: "123456",
	}
	fs[1] = conf.F_spec{
		Name:  "item",
		Class: "text",
		Conf:  "",
		Pos:   2,
	}

	if c.Version != 1 {
		t.Error("version is not 1, ", c.Version)
	}

	if !c.Hashead {
		t.Error(" hashead is not true, ", c.Hashead)
	}

	if c.Fields[0] != fs[0] {
		t.Error("Field_spec compared failed, ", c.Fields)
	}

	if c.Fields[1] != fs[1] {
		t.Error("Field_spec compared failed, ", c.Fields)
	}
	if c.Fields[0].Rangefrom != "123456" {
		t.Error("Rangefrom not set")
	}
}
