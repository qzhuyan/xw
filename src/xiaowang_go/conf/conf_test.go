package conf_test

import (
	"reflect"
	"testing"
	"xiaowang_go/conf"
)

func TestParseCSVConf(t *testing.T) {
	c := conf.ParseConf("../testdata/testcfg.yml")

	// internal conf
	fs := make(map[string]conf.F_spec)

	fs["f1"] = conf.F_spec{
		Class: "numeric",
		Conf:  "",
	}
	fs["f2"] = conf.F_spec{
		Class: "text",
	}

	if c.Version != 1 {
		t.Error("version is not 1, ", c.Version)
	}

	if c.Skipheader {
		t.Error(" usehader is true, ", c.Skipheader)
	}

	if c.Fields["f1"] != fs["f1"] {
		t.Error("Field_spec compared failed, ", c.Fields)
	}

	if !reflect.DeepEqual(c.Header, []string{"f1", "f2", "f3", "f4", "f5"}) {
		t.Error("header not match", c.Header)
	}

	if c.Fields["f2"] != fs["f2"] {
		t.Error("Field_spec compared failed, ", c.Fields)
	}

}
