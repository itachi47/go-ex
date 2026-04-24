package main

import "testing"

func TestDivide(t *testing.T) {
	q, r, err := Divide(10, 3)

	if err != nil {
		t.Fatalf("unexpected error")
	}

	if q != 3 || r != 1 {
		t.Fatalf("got %d %d", q, r)
	}
}

func TestDivideByZero(t *testing.T) {
	_, _, err := Divide(10, 0)

	if err == nil {
		t.Fatalf("expected error")
	}
}
