package email

import (
	"strings"
	"testing"
	"xiaowang_go/conf"
	"xiaowang_go/xwrand"
)

const test_email email = "123@abc.com"

func TestEmailTransform(t *testing.T) {
	res := test_email.Transform(xwrand.NewRandSeed(1), new(conf.F_spec))
	words := strings.Split(res, "@")
	if words[0] == "123" {
		t.Error("user name is should not be  123")
	}
	if words[1] != "abc.com" {
		t.Error("mail server should be abc.com")
	}
}
