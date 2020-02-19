package rootfinder

import (
	"math"
	"testing"
)

var testCases = []struct {
	precision    int
	maxIteration int
	root         float64
	function     Function
	derivative   Function
	a            float64
	b            float64
	guess        float64
	x0           float64
	x1           float64
}{
	{
		precision:    6,
		maxIteration: 100,
		root:         0.739085,
		function:     func(x float64) float64 { return x - math.Cos(x) },
		derivative:   func(x float64) float64 { return 1 + math.Sin(x) },
		a:            0.0,
		b:            1.0,
		guess:        0.5,
		x0:           0,
		x1:           1,
	},
	{
		precision:    6,
		maxIteration: 100,
		root:         3.162278,
		function:     func(x float64) float64 { return x*x - 10 },
		derivative:   func(x float64) float64 { return 2 * x },
		a:            3,
		b:            4,
		guess:        3,
		x0:           3,
		x1:           4,
	},
	{
		precision:    5,
		maxIteration: 100,
		root:         2.09455,
		function:     func(x float64) float64 { return x*x*x - 2*x - 5 },
		derivative:   func(x float64) float64 { return 3*x*x - 2 },
		a:            2,
		b:            3,
		guess:        2,
		x0:           2,
		x1:           3,
	},
}

func TestBisection(t *testing.T) {
	for _, tcase := range testCases {
		rf := New(tcase.precision, tcase.maxIteration, tcase.function, tcase.derivative)

		b, i, _ := rf.Bisection(tcase.a, tcase.b)
		if math.Abs(b-tcase.root) > math.Pow10(-tcase.precision) {
			t.Errorf("expected: %v, got: %v , iter: %v", tcase.root, b, i)
		}
	}
}

func TestNewtonRaphson(t *testing.T) {
	for _, tcase := range testCases {
		rf := New(tcase.precision, tcase.maxIteration, tcase.function, tcase.derivative)

		n, i, _ := rf.NewtonRaphson(tcase.guess)
		if math.Abs(n-tcase.root) > math.Pow10(-tcase.precision) {
			t.Errorf("expected: %v, got: %v , iter: %v", tcase.root, n, i)
		}
	}
}

func TestSecant(t *testing.T) {
	for _, tcase := range testCases {
		rf := New(tcase.precision, tcase.maxIteration, tcase.function, tcase.derivative)

		s, i, _ := rf.Secant(tcase.x0, tcase.x1)
		if math.Abs(s-tcase.root) > math.Pow10(-tcase.precision) {
			t.Errorf("expected: %v, got: %v , iter: %v", tcase.root, s, i)
		}
	}
}
