# Interpolation

## Golang lib for calculating interpolations

#### **!!! matrix should work just fine, other methods WIP !!!**

### Installation

```
go get github.com/piokaczm/interpolation
```

### Introduction

Since my use required it, there are some constraints:
* it returns only values for x int > 0
* it rounds f(x) and returns them as ints

However it should be relatively easy to add start-end points to struct for iteration over range with negative numbers, or add function for returning value for given argument. It'll be added when I have some spare time, or you can contribute to the repo.

### Usage

Each method uses the same public interface.

At first you need to initialize the struct with values and args (you can skip it, then lengths are not validated, and you have to set N manually as well)

```
// example with matrix method
package main

import "github.com/piokaczm/interpolation"

values := []float64{1, 2, 3}
args := []float64{3, 2, 1}

struct := interpolation.Matrix{}
err := struct.Prepare(values, args)
check(err) // handle errors

// or if you feel reckless

reckless_struct := interpolation.Matrix{
  Values: values,
  Args:   args,
  N       3,
}
```

Then you get two options:

You can get a map[int]int in format xi: f(xi);

```
// where n is limit for range [0:n] for which results will be returned
interMap := struct.InterpolationMap(n)
```

Or an []int of f(x)'s

```
// where n is limit for range [0:n] for which results will be returned
interArray := struct.InterpolationArray(n)
```


If you need these estimations within given range of values, you can then flatten the f(x)'s which are outside this range.

```
// where max is maximum desired value and min minimum
normalizedMap := NormalizeMap(interMap, max, min)

normalizedArray := NormalizeArray(interArray, max, min)
```

##### Feel free to contribute (fork m8)

Thanks to gonum team for providing mat64 lib
