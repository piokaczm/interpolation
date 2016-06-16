package interpolation

import (
	"math"
	"testing"
)

var test_args = []float64{2, 3, 6, 7, 8, 10}
var test_values = []float64{0, 2, 3, 5, 1, 2}

// skip val/args validation for tests
var matrix = Matrix{
	Values: test_values,
	Args:   test_args,
	N:      6,
}

func TestMap(t *testing.T) {
	test_map := matrix.InterpolationMap(4)
	Expect(t, len(test_map), 4)
	Expect(t, test_map[2], 0)
	Expect(t, test_map[3], 2)
}

func TestSolution(t *testing.T) {
	solution := matrix.matrixSolve()
	Expect(t, solution[5], 0.050000000000003916)
	Expect(t, roundToInt(solution[0]), -150)
}

func TestMatrixRowX(t *testing.T) {
	row := matrix.matrixRowX(2.0)
	Expect(t, len(row), matrix.N)
	Expect(t, row[5], 1.0)
	Expect(t, row[4], 2.0)
	Expect(t, row[3], 4.0)
	Expect(t, row[2], 8.0)
	Expect(t, row[1], 16.0)
	Expect(t, row[0], 32.0)
}

func TestGetValuesX(t *testing.T) {
	values := matrix.getValuesX()
	Expect(t, len(values), int(math.Pow(float64(matrix.N), 2)))
	Expect(t, values[5], 1.0)
	Expect(t, values[1], 16.0)
	Expect(t, values[0], 32.0)
	Expect(t, values[30], 100000.0)
	Expect(t, values[35], 1.0)
}
