package cell

import "testing"

func testDefaultEmpty(t *testing.T) {
	var c Cell
	if c != EMPTY {
		t.Errorf("Default cell value is not Empty: %v", c)
	}
}
