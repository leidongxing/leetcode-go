package test

import (
	"os"
	"testing"
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 4, 5}
	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	t.Log(a == d)
	t.Log(c)
}

func TestBitClear(t *testing.T) {
	a := 7
	a = a &^ Readable
	t.Log(a&^Readable == Readable, a&^Writable == Writable, a&Executable == Executable)
	os.Exit(0)
}
