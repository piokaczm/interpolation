package interpolation

import (
	"errors"
	"github.com/gonum/matrix/mat64"
	"math"
)

type Matrix struct {
	Values []float64
	Args   []float64
	N      int
}

// call it to feed values to struct for further calculations
// they must be in proper order => xi: arguments[i]; f(xi): values[i]
func (m Matrix) Prepare(values []float64, arguments []float64) error {
	if len(values) == len(arguments) {
		m.Values = values
		m.Args = arguments
		m.N = len(values)
		return nil
	} else {
		return errors.New("interpolation: Wrong input - arguments length is not equal to values length")
	}
}

// get array of f(x)'s
func (m Matrix) InterpolationArray(limit int) []int {
	solution := m.matrixSolve()
	return m.makeArray(solution, limit)
}

// get map in format x: f(x) for given range
func (m Matrix) InterpolationMap(limit int) map[int]int {
	solution := m.matrixSolve()
	return m.makeMap(solution, limit)
}

func (m Matrix) makeMap(solution []float64, limit int) map[int]int {
	results := make(map[int]int)
	for i := 0; i <= limit-1; i++ {
		results[i] = m.extrapolate(solution, float64(i))
	}
	return results
}

func (m Matrix) makeArray(solution []float64, limit int) []int {
	results := make([]int, limit)
	for i := 0; i <= limit-1; i++ {
		results[i] = m.extrapolate(solution, float64(i))
	}
	return results
}

func (m Matrix) extrapolate(solution []float64, arg float64) int {
	exp := float64(m.N)
	values := make([]float64, 0, m.N)
	for i := 1; i <= m.N; i++ {
		values = append(values, solution[m.N-i]*math.Pow(arg, exp-float64(i)))
	}

	var sum float64
	for _, val := range values {
		sum = sum + val
	}
	return roundToInt(sum)
}

func (m Matrix) matrixSolve() []float64 {
	var result mat64.Dense
	matrix_a := m.matrixX(m.getValuesX())
	matrix_b := m.matrixFx(m.Values)
	err := result.Solve(matrix_a, matrix_b)
	check(err)
	solution := m.matrixToArray(result)
	return solution
}

func (m Matrix) matrixToArray(matrix mat64.Dense) []float64 {
	array := make([]float64, m.N)
	limit := m.N - 1
	for i := 0; i <= limit; i++ {
		array[i] = matrix.At(limit-i, 0)
	}
	return array
}

// create matrixes

func (m Matrix) matrixX(values []float64) *mat64.Dense {
	return mat64.NewDense(m.N, m.N, values)
}

func (m Matrix) matrixFx(values []float64) *mat64.Dense {
	return mat64.NewDense(m.N, 1, values)
}

func (m Matrix) getValuesX() []float64 {
	result := make([]float64, 0, int(math.Pow(float64(m.N), 2.0)))
	for _, arg := range m.Args {
		result = append(result, m.matrixRowX(arg)...)
	}
	return result
}

func (m Matrix) matrixRowX(arg float64) []float64 {
	row := make([]float64, 0, m.N)
	for i, _ := range m.Args {
		row = append(row, math.Pow(arg, float64(i)))
	}
	return reverse(row)
}
