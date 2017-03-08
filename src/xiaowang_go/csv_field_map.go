package xiaowang_go

import (
	"gopkg.in/yaml.v2"
	//	"io"
	"io/ioutil"
)

type f_spec struct {
	Transconf
	Name, Class, Conf string
	Pos               int
}

type csv_conf struct {
	Version int
	Fields  []f_spec
	Hashead bool
}

func parse_csv_conf(filepath string) csv_conf {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	var c csv_conf
	err = yaml.Unmarshal(data, &c)
	if nil != err {
		panic(err)
	}
	return c
}

func gen_csv_conf(c csv_conf) string {
	return "todo"
}

//todo: check duplicated fields
