# RootFinder

Rootfinder package helps you to find root of a function with numerical methods.

Let f(x) be a function.

If a is root of f(x), then
```
f(a) = 0
```
f(x) may be algebraic, trigonometric or transcendental function.

## Methods

- Bisection
- Newton-Raphson
- Secant

## Usage
```go
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

func main() {
	rf := rf.New(precision, maxIter, f, df)

	root, iter, err := rf.NewtonRaphson(guess)
	if err == nil {
		fmt.Printf("found the root: %v after %v iters with Newton-Raphson method.\n", iter, root)
	}
}
```