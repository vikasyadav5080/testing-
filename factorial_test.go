package main

import "testing"

/*
input
expected_output == your_output => you pass
*/
func TestFactorial(t *testing.T) {
	for _, c := range testCases {
		if Factorial(c.input) == c.expected {
			t.Logf("Pass : factorial (%d)", c.input)
		} else {
			t.Fatalf("Fail : %s, wanted %d ,got %d", c.description, c.expected, Factorial(c.input))
		}
	}
}

func TestZeroFactorial(t *testing.T) {
	if Factorial(0) == 1 {
		t.Logf("Pass : factorial for zero")
		return
	} else {
		t.Fatalf("Fail : %s, wanted %d ,got %d", "factorial for zero", 1, Factorial(0))
	}
}

func TestOneFactorial(t *testing.T) {
	if Factorial(1) == 1 {
		t.Logf("Pass : factorial for 1")
		return
	} else {
		t.Fatalf("Fail : %s, wanted %d ,got %d", "factorial for one", 1, Factorial(1))
	}
}

func TestNegativeFactorial(t *testing.T) {
	if Factorial(-1) == 0 {
		t.Logf("Pass : factorial for negative")
		return
	} else {
		t.Fatalf("Fail : %s, wanted %d ,got %d", "factorial for negative", 0, Factorial(-1))
	}
}
