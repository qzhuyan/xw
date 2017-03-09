package runner

import (
	//"bufio"
	"encoding/csv"
	"io"
	"log"
	"math/rand"
	"os"
	"xiaowang_go/conf"
	"xiaowang_go/numeric"
	"xiaowang_go/text"
	"xiaowang_go/xwrand"
)

func Run(conffile, srcfile, destfile string) {
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

	seed := xwrand.NewRandSeed(c.Seed)

	if c.Hashead {
		Fields = r.Read()
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

		for _, fspec := range c.Fields {
			// todo: if Pos not defined use header or ?
			record[fspec.Pos] = anonymize(record[fspec.Pos], seed, &fspec)
		}
		w.Write(record)
	}
}

func anonymize(s string, r *rand.Rand, spec *conf.F_spec) string {
	if s == "" {
		return ""
	}
	switch spec.Class {
	case "numeric":
		return numeric.NewNumeric(s).Transform(r, spec)
	case "text":
		return text.Text(s).Transform(r, spec)
	default:
		panic("unknow class")
	}
}
