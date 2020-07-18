package Records

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	var (
		str = "2+3*9-9+4/2"
		expected = 22
	)
	actual := Calculator(str)
	if actual != expected{
		t.Errorf("Actual : %d; expected %d", actual , expected)
	}
}

