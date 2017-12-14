package lpsimplex

import (
	"fmt"
	"math"
	"testing"
)

func TestLinprog(t *testing.T) {
	A, b, c := GetModelSmall_1()
	tests := []struct {
		// min cx s.t. (a:ae)x <= b:be
		a      [][]float64
		b      []float64
		c      []float64
		ae     [][]float64
		be     []float64
		bounds []Bound
		x      []float64
		opt    float64
		intr   int
		errstr string
	}{
		// Basic feasible LP
		// Case 0
		{[][]float64{{-1, 2, 1, 0}, {3, 1, 0, 1}},
			[]float64{4, 9},
			[]float64{-1, -2, 0, 0},
			nil,
			nil,
			[]Bound{},
			[]float64{2, 2, 0, 0},
			-8,
			2,
			"",
		},
		// Case 1
		{[][]float64{{-3, 1}, {1, 2}},
			[]float64{6, 4},
			[]float64{-1, 4},
			nil,
			nil,
			[]Bound{},
			[]float64{0, 0},
			-4,
			1,
			"",
		},
		// Case 2
		{[][]float64{{-3, 1}, {1, 2}},
			[]float64{6, 4},
			[]float64{-1, 4},
			nil,
			nil,
			[]Bound{{math.Inf(-1), math.Inf(1)}, {-3, math.Inf(1)}},
			[]float64{0, 0},
			-22,
			1,
			"",
		},
		// Case 3
		// Taken from an Introduction to Linear Programming and Game Theory, Thie and Keough
		{[][]float64{{4, -1, 0, 1}, {7, -8, -1, 0}}, //Ch. 3, page 59
			[]float64{6, -7},
			[]float64{-3, 2, 1, -1},
			[][]float64{{1, 1, 0, 4}},
			[]float64{12},
			[]Bound{
				{0, math.Inf(1)},
				{0, math.Inf(1)},
				{0, math.Inf(1)},
				{math.Inf(-1), math.Inf(1)},
			},
			[]float64{0, 0, 0, 0},
			-2.235294117647059,
			3,
			"",
		},
		// Case 4
		{nil, // hand converted case 3
			nil,
			[]float64{-3, 2, 1, -1, 1, 0, 0},
			[][]float64{{4, -1, 0, 1, -1, 1, 0}, {-7, 8, 1, 0, 0, 0, -1}, {1, 1, 0, 4, -4, 0, 0}}, //Ch. 3, page 59
			[]float64{6, 7, 12},
			[]Bound{},
			[]float64{0, 0, 0, 0},
			-2.235294117647059,
			3,
			"",
		},
		// Case 5
		{A,
			b,
			c,
			nil,
			nil,
			[]Bound{},
			[]float64{2, 2, 0, 0},
			-490152.5485805898,
			77,
			"",
		},
	}

	tol := 1.0E-12
	//tol := float64(0)
	bland := false
	maxiter := 4000
	//bounds := []Bound(nil)
	//callback := LPSimplexVerboseCallback
	//callback := LPSimplexTerseCallback
	callback := Callbackfunc(nil)
	disp := false //true

	for i, elem := range tests {
		res := LPSimplex(elem.c, elem.a, elem.b, elem.ae, elem.be, elem.bounds, callback, disp, maxiter, tol, bland)
		//fmt.Printf("Res: %+v\n", res)
		//fmt.Printf("Case %d returned with success value of %v and objective value %v\n", i, res.Success, res.Fun)
		if !res.Success {
			t.Errorf("TestLinprog Case %d: failed with message %s\n", i, res.Message)
		}
		if elem.opt != res.Fun {
			t.Errorf("TestLinprog Case %d: Fun: %f but expected %f\n", i, res.Fun, elem.opt)
		}
		if elem.intr != res.Nitr {
			t.Errorf("TestLinprog Case %d: Nitr: %d but expected %d\n", i, res.Nitr, elem.intr)
		}
		/*
			fmt.Printf("linprob_simplex says: %s\n", message)
			fmt.Printf("\ninterations: %d\n", nit)
			if len(x) != len(elem.c) {
				if successful == true {
					fmt.Printf("len(x:%d) != len(c:%d) but should be\n", len(x), len(elem.c))
				} else {
					if x[0] != math.NaN() { // This does not work correctly TODO
						fmt.Printf("Simplex failed: BUT x[0] != NaN it is: %v\n", x[0])
					}
				}
			}
			fmt.Printf("Status: %d\n", status)
			if successful && len(slack) != len(elem.a) {
				fmt.Printf("Slack is not the currect lenth. Should be %d but is %d\n", len(elem.a), len(slack))
			}
			fmt.Printf("Case %d returned with success value of %v and objective value %f\n", i, successful, fun)
		*/
		//fmt.Printf("Case %d returned with success value of %v and objective value %f\n", i, res.Success, res.Fun)
	}
}

func TestPrintT(t *testing.T) {
	tests := []struct {
		a [][]float64
	}{
		// Case 0
		{[][]float64{{-1, 2, 1, 0}, {3, 1, 0, 1}}},
		// Case 1
		{[][]float64{{-1, 2, 1, 0, 8, 8, 8, 8}, {3, 1, 0, 1, 8, 8, 8, 8}}},
		// Case 2
		{[][]float64{{-1, 2, 3, 4, 5, 6, 7, 8, 9}, {-2, 2, 3, 4, 5, 6, 7, 8, 9},
			{-3, 2, 3, 4, 5, 6, 7, 8, 9}, {-4, 2, 3, 4, 5, 6, 7, 8, 9},
			{-5, 2, 3, 4, 5, 6, 7, 8, 9}, {-6, 2, 3, 4, 5, 6, 7, 8, 9},
			{-7, 2, 3, 4, 5, 6, 7, 8, 9}, {-8, 2, 3, 4, 5, 6, 7, 8, 9},
			{-9, 2, 3, 4, 5, 6, 7, 8, 9}, {-10, 2, 3, 4, 5, 6, 7, 8, 9},
			{-11, 2, 3, 4, 5, 6, 7, 8, 9}, {-12, 2, 3, 4, 5, 6, 7, 8, 9}}},
	}

	for i, elem := range tests {
		fmt.Printf("Case %d: \n", i)
		err := TersPrintMatrix(elem.a)
		if err != nil {
			fmt.Printf("Error - %s\n", err)
		}
		// TODO: redirect IO and compare with expected result
	}
}

func TestCountNegEntries(t *testing.T) {
	tests := []struct {
		a      []float64
		expect int
	}{
		{[]float64{1, 2, 3, 4, 5, 6}, 0},
		{[]float64{1, 2, 3, -4, 5, 6}, 1},
		{[]float64{}, 0},
		{nil, 0},
		{[]float64{-1, -2, -3, -4, -5, -6}, 6},
	}
	for i, elem := range tests {
		r := countNegEntries(elem.a)
		if r != elem.expect {
			t.Errorf("Case %d failed: countNegEntries() returned \"%d\" but expected \"%d\"\n", i, r, elem.expect)
		}
	}
}

func TestCheckRectangle(t *testing.T) {
	tests := []struct {
		a          [][]float64
		rows, cols int
		errstr     string
	}{
		{[][]float64{{1, 2}, {3, 4}, {5, 6}}, 3, 2, ""},
		{[][]float64{{1}, {2}, {3}, {4}}, 0, 0, "Invalid input, must be two-dimensional"},
		{[][]float64{}, 0, 0, "Invalid input, must be two-dimensional"},
		{[][]float64{{1, 2}, {3, 4, 10}, {5, 6}}, 0, 0, "Invalid input, all rows must have the same length"},
		{nil, 0, 0, ""},
	}
	for i, elem := range tests {
		rows, cols, err := checkRectangle(elem.a)
		if err != nil {
			if err.Error() != elem.errstr {
				t.Errorf("Case %d failed: checkRectangle() returned \"%v\" but expected \"%s\"\n", i, err, elem.errstr)
			}
		}
		if err == nil && elem.errstr != "" {
			t.Errorf("Case %d failed: checkRectangle() returned nil error but expected error \"%s\"\n", i, elem.errstr)
		}
		if elem.rows != rows {
			t.Errorf("Case %d failed: checkRectangle() returned rows: %d but expected %d\n", i, rows, elem.rows)
		}
		if elem.cols != cols {
			t.Errorf("Case %d failed: checkRectangle() returned cols: %d but expected %d\n", i, cols, elem.cols)
		}
	}
}

/*
func TestGetQuantity(t *testing.T) {
	testCases := []struct {
		line  string
		quant int
	}{
		{"- 1 other stuff", 1},
		{" some stuff 5 other stuff", 5},
		{"- 10 other stuff", 10},
		{"- 162 other stuff", 162},
		{"-1 other stuff", 1},
	}
	for _, testCase := range testCases {
		quantity := getQuantity(testCase.line)
		if quantity != testCase.quant {
			t.Errorf("getQuantity(%s) returned %d but expected %d\n", testCase.line, quantity, testCase.quant)
		}
	}
}

func TestGetTitle(t *testing.T) {
	tstCases := []struct {
		line     string
		expected string
	}{
		{"- 56 A title 76 -- not this $ 6.43", "A title 76"},
		{"- A title -- not this $ 6.43", "A title"},
		{"- A title not this $ 6.43", "A title not this"},
		{"- A title ", "A title"},
		{"A title ", "A title"},
	}
	for _, tstCase := range tstCases {
		title := getTitle(tstCase.line)
		if title != tstCase.expected {
			t.Errorf("getTitle(%s) returned [%s] but [%s] was expected\n", tstCase.line, title, tstCase.expected)
		}
	}
}

func TestGetPrice(t *testing.T) {
	tstCases := []struct {
		line     string
		expected float32
	}{
		{"$6.50", 6.50},
		{" - 23 This is the title -- other stuff $ 6.50  ", 6.50},
	}
	for _, tstCase := range tstCases {
		price := getPrice(tstCase.line)
		if price != tstCase.expected {
			t.Errorf("getPrice(%s) returned [%6.2f] expected [%6.2f]\n", tstCase.line, price, tstCase.expected)
		}
	}
}
*/

func meLPSimplexVerboseCallback(xk []float64, tableau [][]float64, nit, pivrow, pivcol, phase int, basis []int, complete bool) {
	if complete {
		fmt.Printf("--------- Iteration Complete - Phase %d -------\n", phase)
		fmt.Printf("Tableau:\n")
	} else if nit == 0 {
		fmt.Printf("--------- Initial Tableau - Phase %d ----------\n", phase)
	} else {
		fmt.Printf("--------- Iteration %d  - Phase %d --------\n", nit, phase)
		fmt.Printf("Tableau:\n")
	}

	if nit >= 0 {
		err := TersPrintMatrix(tableau)
		if err != nil {
			fmt.Printf("Error - %s\n", err)
		}
		if !complete {
			fmt.Printf("Pivot Element: T[%d, %d]\n", pivrow, pivcol)
		}
		//fmt.Printf("Basic Variables: %v\n", basis)
		fmt.Printf("Basic Variables:\n")
		TersPrintIntArray(basis)
		fmt.Printf("\n")
		fmt.Printf("Current Solution:\n")
		fmt.Printf("x = \n")
		TersPrintArray(xk)
		fmt.Printf("\n")
		fmt.Printf("Current Objective Value:\n")
		fmt.Printf("f = %f\n", -tableau[len(tableau)-1][len(tableau[0])-1])
		fmt.Printf("\n")
	}
}
