package main

import "testing"

func Test_Sum(t *testing.T) {
	x := 5
	y := 5

	want := 10
	got := sum(x, y)

	if want != got {
		t.Logf("testing sum: %d, %d", x, y)
		t.Errorf("Sum incorrect %d, %d", want, got)
	}
}

func Test_sum_negative_numbers(t *testing.T) {
	x := -13
	y := -3

	want := -16
	got := sum(x, y)
	t.Logf("Testing: %d, %d", x, y)
	if want != got {
		t.Errorf("Incorrect %d, %d", want, got)
	}
}
