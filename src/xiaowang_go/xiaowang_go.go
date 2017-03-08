package xiaowang_go

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	//	"math/rand"
	"os"
)

func run(conffile, srcfile string) {
	//read config
	conf := parse_csv_conf(conffile)
	//iter input file
	src, err := os.Open(srcfile)
	if nil != err {
		panic(err)
	}

	r := csv.NewReader(src)

	if conf.Hashead {
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("handling", record)
		for fspec := range conf.Fields {
			// todo: if Pos not defined use header or ?
			//record[fspec.Pos] = anonymize(record[fspec.Pos], fspec)
			log.Println(fspec)
		}

	}
}

/*
func anonymize(s string, r *rand.Rand, spec *f_spec) {
	switch f_spec.Class {
	case "numeric":
		return numeric(s).Transform(r, spec)
	default:
		panic("unknow class")
	}
}
*/
