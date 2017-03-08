package conf

import (
	"gopkg.in/yaml.v2"
	//	"io"
	"io/ioutil"
)

type F_spec struct {
	Name, Class, Conf, Rangefrom, Rangeto string
	Pos                                   int
}

type CSV_conf struct {
	Version int
	Fields  []F_spec
	Hashead bool
	Seed    int64
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
	return c
}

func gen_csv_conf(c CSV_conf) string {
	return "todo"
}

//todo: check duplicated fields
