package rootfinder

import (
	"errors"
	"math"
)

// Function defines a custom type for the function and derivative of it
type Function func(float64) float64

// RootFinder defines a custom object for root finding
type RootFinder interface {
	Bisection(a, b float64) (root float64, iter int, err error)
	NewtonRaphson(initialGuess float64) (root float64, iter int, err error)
	Secant(initialGuesses ...float64) (root float64, iter int, err error)
}

type rootFinder struct {
	maxIteration int
	epsilon      float64
	function     Function
	derivative   Function
}

// New returns a RootFinder object
func New(precision, maxIterarion int, function Function, derivative ...Function) RootFinder {
	if function == nil {
		panic("function is required")
	}
	if precision <= 0 {
		precision = 2
	}
	if maxIterarion <= 0 {
		maxIterarion = 100
	}

	rf := rootFinder{
		epsilon:      math.Pow10(-precision),
		maxIteration: maxIterarion,
		function:     function,
	}
	if len(derivative) > 0 {
		rf.derivative = derivative[0]
	}
	return rf
}

// Bisection finds root of the function by using bisection method
func (rf rootFinder) Bisection(a, b float64) (root float64, iter int, err error) {
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

func derivative(e float64, f Function) Function {
	return func(x float64) float64 {
		return (f(x+e) - f(x)) / e
	}
}

// NewtonRaphson finds root of the function by using newton-raphson method
func (rf rootFinder) NewtonRaphson(initialGuess float64) (root float64, iter int, err error) {
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

// Secant finds root of the function by using secant method
func (rf rootFinder) Secant(initialGuesses ...float64) (root float64, iter int, err error) {
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
