package datastructures

import (
	"bytes"
	"strings"
	"testing"
)

func TestFindHourglass(t *testing.T) {
	b := bytes.NewBufferString(strings.TrimSpace(`
1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 9 2 -4 -4 0
0 0 0 -2 0 0
0 0 -1 -2 -4 0
`))
	//result :=
	if got, want := FindHourglass(b), 13; got != want {
		t.Errorf("wrong answer: got=%v", got)
	}
}
