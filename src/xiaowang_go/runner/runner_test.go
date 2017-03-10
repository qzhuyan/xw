package runner_test

import (
	"encoding/csv"
	"io"
	"os"
	"reflect"
	"testing"
	"xiaowang_go/runner"
)

func TestRun(t *testing.T) {
	from := "../testdata/test.csv"
	to := "../testdata/result.csv"
	runner.Run("../testdata/testcfg.yml", from, to)
	compare_csv_file(from, to, t)
}

func TestRunMeaningful(t *testing.T) {
	from := "../testdata/test2.csv"
	to := "../testdata/result2.csv"
	runner.Run("../testdata/testcfg2.yml", from, to)
	compare_csv_file(from, to, t)
}

func compare_csv_file(file1, file2 string, t *testing.T) {
	//now check file identical
	src, err := os.Open(file1)
	if err != nil {
		t.Error(err)
	}
	dest, err := os.Open(file2)
	if err != nil {
		t.Error(err)
	}

	s := csv.NewReader(src)
	d := csv.NewReader(dest)

	for {
		data1, err1 := s.Read()
		data2, err2 := d.Read()
		if err1 == io.EOF && err2 == io.EOF {
			break
		} else if err1 != nil || err2 != nil {
			t.Error("compare data failed, one of the read failed", err1, err2)
		} else if !reflect.DeepEqual(data1, data2) {
			t.Errorf("line by line comparing failed \n%s\n%s\n", data1, data2)
		}
	}
}
