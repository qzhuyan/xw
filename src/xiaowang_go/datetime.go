package xiaowang_go

import "time"
import "fmt"
import "math/rand"

type datetime string

func (d datetime) Transform(r *rand.Rand, conf *Transconf) string {
	from := datetime(conf.Rangefrom)
	to := datetime(conf.Rangeto)
	d.rangeok(from, to)
	unixfrom := from.Parse().Unix()
	unixto := to.Parse().Unix()
	newtime := RandInRange(r, unixfrom, unixto)
	return time.Unix(newtime+unixfrom, 0).Format("20060102 15:04:05")
}

func (d datetime) Parse() time.Time {
	t, error := time.Parse("20060102 15:04:05", string(d))
	if error != nil {
		fmt.Println("Parse dattime string failed", d, error)
		panic("invalid datetime")
	}
	return t
}

//panic if rang check failed
func (d datetime) rangeok(from, to datetime) bool {
	return from < d && d < to
}
