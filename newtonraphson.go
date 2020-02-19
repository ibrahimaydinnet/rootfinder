package rootfinder

import (
	"errors"
	"math"
)

func derivative(e float64, f Function) Function {
	return func(x float64) float64 {
		return (f(x+e) - f(x)) / e
	}
}

// NewtonRaphson finds root of the function by using newton-raphson method
func (rf *RootFinder) NewtonRaphson(initialGuess float64) (root float64, iter int, err error) {
	root = 0
	iter = 0
	err = nil
	next := func(x0 float64, f, df Function) float64 {
		iter++
		return x0 - f(x0)/df(x0)
	}
	stop := func(x0, x1 float64) bool {
		return math.Abs(rf.function(x1)) <= rf.epsilon || math.Abs(x1-x0) <= rf.epsilon
	}

	if rf.derivative == nil {
		rf.derivative = derivative(rf.epsilon, rf.function)
	}

	x0 := initialGuess
	x1 := next(x0, rf.function, rf.derivative)

	for iter <= rf.maxIteration && !stop(x0, x1) {
		x0 = x1
		x1 = next(x0, rf.function, rf.derivative)
	}

	if !stop(x0, x1) {
		err = errors.New("root not found")
		return
	}

	rounder := 1 / rf.epsilon
	root = math.Round(x1*rounder) / rounder
	return
}
