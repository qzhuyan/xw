package datetime

import "time"
import "fmt"
import "math/rand"
import "xiaowang_go/xwrand"
import "xiaowang_go/conf"

func Transform(d string, r *rand.Rand, c *conf.F_spec) string {
	from := c.Rangefrom
	to := c.Rangeto
	if from != "" && to != "" {
		Rangeok(d, from, to)
		unixfrom := Parse(from).Unix()
		unixto := Parse(to).Unix()
		newtime := xwrand.RandInRange(r, unixfrom, unixto)
		return time.Unix(newtime+unixfrom, 0).Format("2006-01-02 15:04:05.000")
	} else {
		unixfrom := Parse(d).Unix()
		unixto := Parse(d).AddDate(0, 0, 30).Unix()
		newtime := xwrand.RandInRange(r, unixfrom, unixto)
		return time.Unix(newtime+unixfrom, 0).Format("2006-01-02 15:04:05.000")
	}
}

//todo make time format flexable
func Parse(d string) time.Time {
	t, error := time.Parse("2006-01-02 15:04:05.000", d)
	if error != nil {
		fmt.Println("Parse datetime string failed", d, error)
		panic("invalid Datetime")
	}
	return t
}

//panic if rang check failed
func Rangeok(d, from, to string) bool {
	return from < d && d < to
}
