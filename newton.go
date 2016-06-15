package main

import (
	"errors"
	"fmt"
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
// func (n Newton) ArrayData(limit int) []int {
// 	// calculations
// 	return n.makeArray(solution, limit)
// }

// get map in format x: f(x) for given range
func (n Newton) InterpolationMap(limit int) map[int]int {
	diffs := n.calcDiffs()
	return n.makeMap(diffs, limit)
}

func (n Newton) makeMap(diffs []float64, limit int) map[int]int {
	results := make(map[int]int)
	for i := 0; i <= limit-1; i++ {
		results[i] = n.calcValue(diffs, float64(i))
	}
	return results
}

func (n Newton) calcValue(diffs []float64, x float64) int {
	fmt.Print("IN CALC-VALUE\n")
	sum := n.Values[0]
	for i, diff := range diffs {
		multipliers := make([]float64, i)
		for j := i; j >= 0; j-- {
			multipliers = append(multipliers, x-n.Args[j])
		}
		multi := 1.0
		for _, m := range multipliers {
			multi = multi * m
		}
		sum = sum + diff*multi
	}
	return roundToInt(sum)
}

func (n Newton) calcDiffs() []float64 {
	values := make([]float64, n.N)
	for i := 0; i <= n.N-1; i++ {
		fmt.Printf("IN CALC-DIFF, i: %v\n", i)
		values[i] = n.singleDiff(i)
	}
	return values
}

func (n Newton) singleDiff(j int) float64 {
	fmt.Print("IN SINGLE DIFF\n")
	if j == 1 {
		return n.diff(j)
	} else {
		nom_vals := make([]float64, 0, j)
		for i := j; i >= 0; i-- {
			nom_vals = append(nom_vals, n.diff(i))
		}
		nominator := nom_vals[0]
		for _, val := range nom_vals[1:] {
			nominator = nominator - val
		}
		return nominator / (n.Args[j] - n.Args[0])
	}
}

func (n Newton) diff(j int) float64 {
	fmt.Printf("in diff, j: %v \n", j)
	if j >= 1 {
		return (n.Values[j] - n.Values[j-1]) / (n.Args[j-1] - n.Args[j])
	} else {
		return 0
	}
}
