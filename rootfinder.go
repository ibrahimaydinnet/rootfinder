package rootfinder

import (
	"math"
)

// Function defines a custom type for the function and derivative of it
type Function func(float64) float64

// RootFinder defines a custom object for root finding
type RootFinder struct {
	maxIteration int
	epsilon      float64
	function     Function
	derivative   Function
}

// New returns a RootFinder object
func New(precision, maxIterarion int, function Function, derivative ...Function) *RootFinder {
	if function == nil {
		panic("function is required")
	}
	if precision <= 0 {
		precision = 2
	}
	if maxIterarion <= 0 {
		maxIterarion = 100
	}

	rf := RootFinder{
		epsilon:      math.Pow10(-precision),
		maxIteration: maxIterarion,
		function:     function,
	}
	if len(derivative) > 0 {
		rf.derivative = derivative[0]
	}
	return &rf
}
