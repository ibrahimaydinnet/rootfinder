package rootfinder

import (
	"errors"
	"math"
)

// Secant finds root of the function by using secant method
func (rf *RootFinder) Secant(initialGuesses ...float64) (root float64, iter int, err error) {
	root = 0
	iter = 0
	err = nil
	next := func(x0, x1 float64, f Function) float64 {
		iter++
		return (x0*f(x1) - x1*f(x0)) / (f(x1) - f(x0))
	}
	stop := func(x0, x1 float64) bool {
		return math.Abs(x1-x0) <= rf.epsilon || math.Abs(x1-x0) <= rf.epsilon
	}

	x0 := 0.0
	x1 := 1.0
	if len(initialGuesses) > 2 {
		x1 = initialGuesses[1]
	}
	if len(initialGuesses) > 1 {
		x0 = initialGuesses[0]
	}

	x2 := next(x0, x1, rf.function)
	x0 = x1
	x1 = x2

	for iter <= rf.maxIteration && !stop(x0, x1) {
		x2 = next(x0, x1, rf.function)
		x0 = x1
		x1 = x2
	}

	if !stop(x0, x1) {
		err = errors.New("root not found")
		return
	}

	rounder := 1 / rf.epsilon
	root = math.Round(x1*rounder) / rounder
	return
}
