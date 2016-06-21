package interpolation

import (
	"errors"
)

type Newton struct {
	Values []float64
	Args   []float64
	N      int
}

func (n Newton) Prepare(values []float64, arguments []float64) error {
	if len(values) == len(arguments) {
		n.Values = values
		n.Args = arguments
		n.N = len(values)
		return nil
	} else {
		return errors.New("interpolation: Wrong input - arguments length is not equal to values length")
	}
}

// get array of f(x)'s
func (n Newton) InterpolationArray(limit int) []int {
	// calculations
	return n.makeArray(limit)
}

// get map in format x: f(x) for given range
func (n Newton) InterpolationMap(limit int) map[int]int {
	return n.makeMap(limit)
}

func (n Newton) makeMap(limit int) map[int]int {
	results := make(map[int]int)
	for i := 0; i <= limit-1; i++ {
		results[i] = n.calcValue(i)
	}
	return results
}

func (n Newton) makeArray(limit int) []int {
	results := make([]int, limit)
	for i := 0; i <= limit-1; i++ {
		results[i] = n.calcValue(i)
	}
	return results
}

func (newt Newton) calcValue(k int) int {
	// clean it up mate
	x := append([]float64{0}, newt.Args...)
	y := append([]float64{0}, newt.Values...)
	j := 1
	f1 := 1.0
	f2 := 0.0
	f := y[1]
	p := make([]float64, newt.N)
	for n := newt.N; n > 1; n-- {
		for i := 1; i <= n-1; i++ {
			p[i] = ((y[i+1] - y[i]) / (x[i+j] - x[i]))
			y[i] = p[i]
		}
		f1 = 1
		for i := 1; i <= j; i++ {
			f1 *= (float64(k) - x[i])
		}
		f2 += (y[1] * f1)
		j++
	}
	f += f2
	return roundToInt(f)
}
