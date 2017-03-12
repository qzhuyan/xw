package runner

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"xiaowang_go/conf"
	"xiaowang_go/datetime"
	"xiaowang_go/email"
	"xiaowang_go/ip"
	"xiaowang_go/numeric"
	"xiaowang_go/pno"
	"xiaowang_go/text"
	"xiaowang_go/xwrand"
)

func Run(conffile, srcfile, destfile string, seed int64) {
	//read config
	c := conf.ParseConf(conffile)
	//iter input file
	src, err := os.Open(srcfile)

	if nil != err {
		panic(err)
	}

	dest, err := os.Create(destfile)
	if nil != err {
		panic(err)
	}

	r := csv.NewReader(src)
	w := csv.NewWriter(dest)

	if seed != 0 {
		c.Seed = seed
	}

	rsrc := xwrand.NewRandSeed(c.Seed)

	if !c.Skipheader {
		data, err := r.Read()
		if err != nil {
			panic("not able to read header")
		}
		c.Header = data
		w.Write(data)
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// decide csv file head

		for pos, name := range c.Header {
			spec := c.Fields[name]
			record[pos] = anonymize(record[pos], rsrc, &spec)
		}
		w.Write(record)
	}

	//flush it
	w.Flush()
}

func anonymize(s string, r *rand.Rand, spec *conf.F_spec) string {
	defer func() string {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in anonymize class: %s str: %s\n", spec.Class, r)
		}
		return ""
	}()

	if s == "" {
		return ""
	}
	switch spec.Class {
	case "":
		return s
	case "numeric":
		return numeric.NewNumeric(s).Transform(r, spec)
	case "text":
		return text.Text(s).Transform(r, spec)
	case "datetime":
		return datetime.Transform(s, r, spec)
	case "email":
		return email.Email(s).Transform(r, spec)
	case "ip":
		return ipaddr.Ipaddr(s).Transform(r, spec)
	case "pno4":
		return pno.Transform(s, r, spec)
	default:
		panic("unknow class")
	}
}
