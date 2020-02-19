package rootfinder

import (
	"errors"
	"math"
)

// Bisection finds root of the function by using bisection method
func (rf *RootFinder) Bisection(a, b float64) (root float64, iter int, err error) {
	root = 0
	iter = 0
	err = nil
	next := func(a, b, x0 float64, f Function) (float64, float64, float64) {
		iter++
		if f(a)*f(x0) < 0 {
			return a, x0, (a + x0) / 2
		}
		return x0, b, (x0 + b) / 2
	}
	stop := func(x0, x1 float64, iter int) bool {
		return math.Abs(rf.function(x1)) <= rf.epsilon || math.Abs(x1-x0) <= rf.epsilon
	}

	if math.Abs(rf.function(a)) <= rf.epsilon {
		root = a
		return
	}
	if math.Abs(rf.function(b)) <= rf.epsilon {
		root = b
		return
	}
	if rf.function(a)*rf.function(b) > 0 {
		err = errors.New("[a,b] is invalid interval")
		return
	}

	x0 := (a + b) / 2
	a, b, x1 := next(a, b, x0, rf.function)

	for iter <= rf.maxIteration && !stop(x0, x1, iter) {
		x0 = x1
		a, b, x1 = next(a, b, x0, rf.function)
	}

	if !stop(x0, x1, iter) {
		err = errors.New("root not found")
		return
	}

	rounder := 1 / rf.epsilon
	root = math.Round(x1*rounder) / rounder
	return
}
