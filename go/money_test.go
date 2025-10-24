package main

import "testing"

func TestMultiplication(t *testing.T) {
	fiver := Dollar{
		amount:   5,
		currency: "USD",
	}
	tenner := fiver.Times(2)
	if tenner.amount != 10 {
		t.Errorf("Expected 10, got [%d]", tenner.amount)
	}
}

type Dollar struct {
	amount   int
	currency string
}

func (d Dollar) Times(n int) Dollar {
	return Dollar{amount: d.amount * n}
}
