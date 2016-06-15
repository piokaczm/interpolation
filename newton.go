package main

import (
	"errors"
	"github.com/gonum/matrix/mat64"
	"math"
)

type Newton struct {
	Values []float64
	Args   []float64
	N      int
}
