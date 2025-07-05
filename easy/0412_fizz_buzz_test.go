package easy

import "testing"

func TestFizzBuzz(t *testing.T) {
	exp := []string{
		"1",
		"2",
		"Fizz",
		"4",
		"Buzz",
		"Fizz",
		"7",
		"8",
		"Fizz",
		"Buzz",
		"11",
		"Fizz",
		"13",
		"14",
		"FizzBuzz",
	}
	res := fizzBuzz(15)
	if len(res) != len(exp) {
		t.Fatalf("Expected len %d, got %d: %#v", len(exp), len(res), res)
	}
	for i, v := range exp {
		if res[i] != v {
			t.Errorf("Mismatch at index %d: want %s, got %v", i, v, res[i])
		}
	}
}
