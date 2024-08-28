package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSubtract(t *testing.T) {
	total := subtract(6, -5)
	if total != 11 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 11)
	}
}

func TestDomathadd(t *testing.T) {
	total := doMath(6, -5, add)
	if total != 1 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 1)
	}
}

func TestDomathsubtract(t *testing.T) {
	total := doMath(8, 4, subtract)
	if total != 4 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 4)
	}
}
