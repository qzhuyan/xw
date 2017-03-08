package xiaowang_go

import "testing"

func TestRun(t *testing.T) {
	run("testdata/testcfg.yml", "testdata/test.csv", "testdata/result.csv")
	t.Log("todo")
}
