package wok

import (
	"testing"
)

func TestBool(t *testing.T) {
	boolArray := true
	if !boolArray {
		t.Error("Test failed")
	}
}
