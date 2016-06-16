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

func TestDiff(t *testing.T) {
	n_test_diff_2 := newton.diff(2)
	Expect(t, n_test_diff_2, 10.0)
	n_test_diff_3 := newton.diff(3)
	Expect(t, n_test_diff_3, 13.0)
}

func TestSingleDiff(t *testing.T) {
	n_test_single_diff := newton.singleDiff(2)
	Expect(t, n_test_single_diff, 1.0)
	n_test_single_diff_2 := newton.singleDiff(3)
	Expect(t, n_test_single_diff_2, -0.8)
}

func TestNMap(t *testing.T) {
	n_test_map := newton.InterpolationMap(7)
	Expect(t, n_test_map[1], 11)
	Expect(t, n_test_map[3], 25)
	Expect(t, n_test_map[4], 35)
	Expect(t, n_test_map[6], 61)
}
