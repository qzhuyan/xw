package xiaowang_go

import (
	//"bufio"
	"encoding/csv"
	"io"
	"log"
	"math/rand"
	"os"
)

func run(conffile, srcfile, destfile string) {
	//read config
	conf := parse_csv_conf(conffile)
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

	seed := NewRandSeed(conf.Seed)
	if conf.Hashead {
		r.Read()
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		for _, fspec := range conf.Fields {
			// todo: if Pos not defined use header or ?
			record[fspec.Pos] = anonymize(record[fspec.Pos], seed, &fspec)
			w.Write(record)
		}
	}
}

func anonymize(s string, r *rand.Rand, spec *f_spec) string {
	if s == "" {
		return ""
	}
	switch spec.Class {
	case "numeric":
		return NewNumeric(s).Transform(r, spec)
	case "text":
		return text(s).Transform(r, spec)
	default:
		panic("unknow class")
	}
}
