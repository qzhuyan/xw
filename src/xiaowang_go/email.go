package xiaowang_go

import (
	"math/rand"
	"strings"
)

type email string

func (e email) Transform(r *rand.Rand, conf *f_spec) string {
	earray := strings.Split(string(e), "@")
	user := earray[0]
	mailserver := earray[1]
	newuser := text(user).Transform(r, conf)
	return newuser + "@" + mailserver
}
