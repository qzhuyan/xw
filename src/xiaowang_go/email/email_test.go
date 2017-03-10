package email_test

import (
	"strings"
	"testing"
	"xiaowang_go/conf"
	"xiaowang_go/email"
	"xiaowang_go/xwrand"
)

const test_email email.Email = "alice@abc.com"

func TestEmailTransform(t *testing.T) {
	res := test_email.Transform(xwrand.NewRandSeed(1), new(conf.F_spec))
	words := strings.Split(res, "@")
	if words[0] == "123" {
		t.Error("user name is should not be 123")
	}
	if words[1] != "abc.com" {
		t.Error("mail server should be abc.com")
	}

	res2 := test_email.Transform(xwrand.NewRandSeed(1), new(conf.F_spec))
	if res2 != res {
		t.Error("random not stable", res2, res)
	}
}
