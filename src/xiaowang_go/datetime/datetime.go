package datetime

import "time"
import "fmt"
import "math/rand"
import "xiaowang_go/xwrand"
import "xiaowang_go/conf"

var Defaultfmt string = "2006-01-02 15:04:05.000"

func Transform(d string, r *rand.Rand, c *conf.F_spec) string {
	from := c.Rangefrom
	to := c.Rangeto

	if c.FormatIn == "" {
		c.FormatIn = Defaultfmt
	}

	if from == "" && to == "" {
		unixfrom := Parsefmt(d, c.FormatIn).Unix()
		newtime := xwrand.RandInRange(r, unixfrom, unixfrom+10000000)
		return time.Unix(newtime, 0).Format(c.FormatIn)

	}

	Rangeok(d, from, to)
	fmt.Println(d, from, to)
	unixfrom := Parsefmt(from, c.FormatIn).Unix()
	unixto := Parsefmt(to, c.FormatIn).Unix()
	newtime := xwrand.RandInRange(r, unixfrom, unixto)
	return time.Unix(newtime+unixfrom, 0).Format("2006-01-02 15:04:05.000")
}

//todo make time format flexable
func Parse(d string) time.Time {
	return Parsefmt(d, "2006-01-02 15:04:05.000")
}

func Parsefmt(d string, format string) time.Time {
	t, error := time.Parse(format, d)
	if error != nil {
		fmt.Println("Parse datetime string failed", d, error)
		panic("cannot parse time " + d + "format" + format)
	}
	return t
}

//panic if rang check failed
func Rangeok(d, from, to string) bool {
	return from < d && d < to
}
