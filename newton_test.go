package main

import (
	"testing"
	// "math"
)

var n_test_args = []float64{1, 3, 4}
var n_test_values = []float64{11, 25, 35}

// skip val/args validation for tests
var newton = Newton{
	Values: n_test_values,
	Args:   n_test_args,
	N:      3,
}

func TestNMap(t *testing.T) {
	n_test_map := newton.InterpolationMap(3)
	Expect(t, n_test_map[1], 11)
	Expect(t, n_test_map[4], 35)
}
