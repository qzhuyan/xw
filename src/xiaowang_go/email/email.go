package email

import (
	"math/rand"
	"strings"
	"xiaowang_go/conf"
	"xiaowang_go/text"
)

type Email string

func (e Email) Transform(r *rand.Rand, conf *conf.F_spec) string {
	earray := strings.Split(string(e), "@")
	user := earray[0]
	mailserver := earray[1]
	newuser := text.Text(user).Transform(r, conf)
	return newuser + "@" + mailserver
}
