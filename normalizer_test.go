package interpolation

import "testing"

var real_args = []float64{0, 180, 720, 960, 1380}
var real_vals = []float64{4000.0, 400.0, 600.0, 800.0, 3600.0}

// skip val/args validation for tests
var m = Matrix{
	Values: real_vals,
	Args:   real_args,
	N:      5,
}

func TestNormalizeMap(t *testing.T) {
	norm_map := m.InterpolationMap(1440)
	val1 := norm_map[0]
	val2 := norm_map[400]
	NormalizeMap(norm_map, 4000, 0)
	Expect(t, val1, norm_map[0])
	ExpectNotEqual(t, val2, norm_map[400])
}

func TestNormalizeArray(t *testing.T) {
	norm_arr := m.InterpolationArray(1440)
	val1 := norm_arr[0]
	val2 := norm_arr[400]
	NormalizeArray(norm_arr, 4000, 0)
	Expect(t, val1, norm_arr[0])
	ExpectNotEqual(t, val2, norm_arr[400])
}

func TestNormalizeWithEqualMinMax(t *testing.T) {
	norm_arr := m.InterpolationArray(1440)
	NormalizeArray(norm_arr, 15, 15)
	Expect(t, 15, norm_arr[0])
	Expect(t, 15, norm_arr[100])
	Expect(t, 15, norm_arr[1000])
}
