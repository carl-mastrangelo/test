package test

import (
	"testing"
)

func Test(t *testing.T) {
	if doThing() != 1 {
		t.Error("bad")
	}
}
