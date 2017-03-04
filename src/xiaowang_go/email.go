package xiaowang_go

import (
	"strings"
)

type email string

func (e email) Transform(seed int64) string {
	earray := strings.Split(string(e), "@")
	user := earray[0]
	mailserver := earray[1]
	newuser := text(user).Transform(seed)
	return newuser + "@" + mailserver
}
