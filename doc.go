/*
Package rootfinder contains numerical methods for finding root of a function.

Methods:
	- Bisection
	- NewtonRaphson
	- Secant

Usage:

	rf := rf.New(precision, maxIteration, function, derivative)
	root, iter, err := rf.NewtonRaphson(guess)
*/
package rootfinder
