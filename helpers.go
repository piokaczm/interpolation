package interpolation

import (
	"math"
	"reflect"
	"testing"
)

// reverse slice
func reverse(numbers []float64) []float64 {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

// round floats to ints
func roundToInt(val float64) int {
	var round float64
	_, div := math.Modf(val)
	if div >= .5 {
		round = math.Ceil(val)
	} else {
		round = math.Floor(val)
	}
	return int(round)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func ExpectNotEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
