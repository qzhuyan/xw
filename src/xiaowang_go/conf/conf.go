package conf

import (
	"gopkg.in/yaml.v2"
	//	"io"
	"io/ioutil"
	//"strings"
)

type F_spec struct {
	Class                          string `yaml:"transform"`
	Name, Conf, Rangefrom, Rangeto string
}

type CSV_conf struct {
	Version    int
	Fields     map[string]F_spec
	Skipheader bool
	Seed       int64
	Header     []string
}

func ParseConf(filepath string) CSV_conf {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	var c CSV_conf
	err = yaml.Unmarshal(data, &c)
	if nil != err {
		panic(err)
	}
	if c.Skipheader && len(c.Header) > 0 {
		panic("in config file both hashead and filenames are defined, which one to use?")
	}

	return c
}

func gen_csv_conf(c CSV_conf) string {
	return "todo"
}

//todo: check duplicated fields
