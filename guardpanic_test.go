package guardpanic

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	worker := func() {
		var s []int
		s[0] = 1 // panic
	}
	printErrCb := func(err error) {
		if err == nil {
			t.Fatal("Unexpected error")
		}

		s := err.Error()
		if s == "" || !strings.Contains(s, "goroutine") {
			t.Fatal("Unexpected error")
		}
	}

	Run(worker, 0, printErrCb)
}