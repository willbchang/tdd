package main

import "testing"

func TestMultiplication(t *testing.T) {
	fiver := Money{
		amount:   5,
		currency: "USD",
	}
	tenner := fiver.Times(2)
	if tenner.amount != 10 {
		t.Errorf("Expected 10, got [%d]", tenner.amount)
	}
}

func TestMultiplicationInEuro(t *testing.T) {
	money := Money{
		amount:   10,
		currency: "EUR",
	}
	result := money.Times(2)
	if result.amount != 20 {
		t.Errorf("Expected 20, got[%d]", result.amount)
	}
}

type Money struct {
	amount   int
	currency string
}

func (d Money) Times(n int) Money {
	return Money{amount: d.amount * n, currency: d.currency}
}
