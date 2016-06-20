package interpolation

import (
	"testing"
	// "math"
)

var n_test_args = []float64{1, 3, 4, 6}
var n_test_values = []float64{11, 25, 35, 61}

// skip val/args validation for tests
var newton = Newton{
	Values: n_test_values,
	Args:   n_test_args,
	N:      4,
}

func TestNMap(t *testing.T) {
	n_test_map := newton.InterpolationMap(7)
	Expect(t, n_test_map[1], 11)
	Expect(t, n_test_map[3], 25)
	Expect(t, n_test_map[4], 35)
	Expect(t, n_test_map[6], 61)
}

func TestNArray(t *testing.T) {
	n_test_array := newton.InterpolationMap(7)
	Expect(t, n_test_array[1], 11)
	Expect(t, n_test_array[3], 25)
	Expect(t, n_test_array[4], 35)
	Expect(t, n_test_array[6], 61)
}

func TestCalcValue(t *testing.T) {
	val_6 := newton.calcValue(6)
	Expect(t, val_6, 61)
	val_4 := newton.calcValue(4)
	Expect(t, val_4, 35)
	val_3 := newton.calcValue(3)
	Expect(t, val_3, 25)
	val_1 := newton.calcValue(1)
	Expect(t, val_1, 11)
}
