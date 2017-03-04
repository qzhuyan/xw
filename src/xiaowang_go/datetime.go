package xiaowang_go

import "time"
import "fmt"

type datetime string

func (d datetime) Transform(seed uint64) time.Time {
	epoch, error := time.Parse("20060102 15:04:05", string(d))
	if error != nil {
		fmt.Println("same datetime with diffent seeds", d, error)
	}

	return epoch
}
