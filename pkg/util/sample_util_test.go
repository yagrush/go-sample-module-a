package util

import (
	"testing"
)

func TestCalcAddInt64(t *testing.T) {
	a, b := int64(1), int64(2)
	answer := a + b
	if calcAddInt64(a, b) != answer {
		t.Errorf("%d + %d != %d", a, b, answer)
	}
}
