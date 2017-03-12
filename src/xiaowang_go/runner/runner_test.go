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
	os.Remove(to)
	runner.Run("../testdata/testcfg.yml", from, to, 0)
	csv_file_eq(from, to, false, t)
}

func TestRunMeaningful(t *testing.T) {
	from := "../testdata/test2.csv"
	to := "../testdata/result2.csv"
	os.Remove(to)
	runner.Run("../testdata/testcfg2.yml", from, to, 0)
	csv_file_not_eq(from, to, true, t)
}

func csv_file_not_eq(file1, file2 string, hasheader bool, t *testing.T) {
	compare_csv_file(false, hasheader, file1, file2, t)
}

func csv_file_eq(file1, file2 string, hasheader bool, t *testing.T) {
	compare_csv_file(true, hasheader, file1, file2, t)
}

func compare_csv_file(expectEq bool, hasheader bool, file1, file2 string, t *testing.T) {
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
	linecnt := 0
	for {
		data1, err1 := s.Read()
		data2, err2 := d.Read()
		if err1 == io.EOF && err2 == io.EOF {
			break
		} else if err1 != nil || err2 != nil {
			t.Error("compare data failed, one of the read failed", err1, err2)
		} else if linecnt == 0 && hasheader {
			if !reflect.DeepEqual(data1, data2) {
				t.Errorf("csv has header but not match \n%s\n%s\n", data1, data2)
			}
		} else if reflect.DeepEqual(data1, data2) && !expectEq {
			t.Errorf("line by line comparing, not expected result lineno: %d\n%s\n%s\n", linecnt, data1, data2)
		}
		linecnt++
	}
}
