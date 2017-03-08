package runner_test

import "testing"
import "xiaowang_go/runner"

func TestRun(t *testing.T) {
	runner.Run("../testdata/testcfg.yml", "../testdata/test.csv", "../testdata/result.csv")
	t.Log("todo")
}
