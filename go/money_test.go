package main

import "testing"

func TestMultiplication(t *testing.T) {
	fiver := Money{
		amount:   5,
		currency: "USD",
	}
	tenner := fiver.Times(2)
	if tenner.amount != 10 {
		t.Errorf("Expected 10, got [%f]", tenner.amount)
	}
}

func TestMultiplicationInEuro(t *testing.T) {
	money := Money{
		amount:   10,
		currency: "EUR",
	}
	result := money.Times(2)
	if result.amount != 20 {
		t.Errorf("Expected 20, got[%f]", result.amount)
	}
}

func TestDivisionKRW(t *testing.T) {
	money := Money{
		amount:   4002,
		currency: "KRW",
	}
	result := money.Divides(4)
	if result.amount != 1000.5 {
		t.Errorf("Expected 1000.5, got [%f]", result.amount)
	}
}

func TestAddUsdAndEur(t *testing.T) {
	usd := Money{
		amount:     5,
		currency:   "USD",
		ratioToUSD: 1,
	}

	eur := Money{
		amount:     10,
		currency:   "EUR",
		ratioToUSD: 1.2,
	}

	result := usd.Adds(eur)
	if result.amount != 17 {
		t.Errorf("Expected 17, got [%f]", result.amount)
	}

	if result.currency != "USD" {
		t.Errorf("Exptected USD, got [%s]", result.currency)
	}
}

type Money struct {
	amount     float32
	currency   string
	ratioToUSD float32
}

func (m Money) Times(n float32) Money {
	return Money{amount: m.amount * n, currency: m.currency}
}

func (m Money) Divides(n float32) Money {
	return Money{amount: m.amount / n, currency: m.currency}
}

func (m Money) Adds(o Money) Money {
	return Money{
		amount:     m.amount + o.amount*o.ratioToUSD,
		currency:   m.currency,
		ratioToUSD: m.ratioToUSD,
	}
}
