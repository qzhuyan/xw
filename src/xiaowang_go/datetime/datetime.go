package datetime

import "time"
import "fmt"
import "math/rand"
import "xiaowang_go/xwrand"
import "xiaowang_go/conf"

type Datetime string

func (d Datetime) Transform(r *rand.Rand, c *conf.F_spec) string {
	from := Datetime(c.Rangefrom)
	to := Datetime(c.Rangeto)
	d.Rangeok(from, to)
	unixfrom := from.Parse().Unix()
	unixto := to.Parse().Unix()
	newtime := xwrand.RandInRange(r, unixfrom, unixto)
	return time.Unix(newtime+unixfrom, 0).Format("20060102 15:04:05.000")
}

func (d Datetime) Parse() time.Time {
	t, error := time.Parse("20060102 15:04:05.000", string(d))
	if error != nil {
		fmt.Println("Parse dattime string failed", d, error)
		panic("invalid Datetime")
	}
	return t
}

//panic if rang check failed
func (d Datetime) Rangeok(from, to Datetime) bool {
	return from < d && d < to
}
