package main

import (
	"fmt"
	rf "github.com/ibrahimaydinnet/rootfinder"
)

var f = func(x float64) float64 {
	return x*x - 10
}
var df = func(x float64) float64 {
	return 2 * x
}
var precision = 6
var maxIter = 100
var guess = 3.0
var a, b = 3.0, 4.0
var x0, x1 = 0.0, 1.0

func main() {
	rf := rf.New(precision, maxIter, f, df)

	root, iter, err := rf.Bisection(a, b)
	if err == nil {
		fmt.Printf("after %v iters, found the root: %v with Bisection method.\n", iter, root)
	}

	root, iter, err = rf.NewtonRaphson(guess)
	if err == nil {
		fmt.Printf("after %v iters, found the root: %v with Newton-Raphson method.\n", iter, root)
	}

	root, iter, err = rf.Secant(x0, x1)
	if err == nil {
		fmt.Printf("after %v iters, found the root: %v with Secant method.\n", iter, root)
	}
}
