package main

import (
	"math"
	"testing"
)

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

func TestAddUsdAndKrwOutputKrw(t *testing.T) {
	usd := Money{
		amount:     1,
		currency:   "USD",
		ratioToUSD: 1,
	}

	krw := Money{
		amount:     1100,
		currency:   "KRW",
		ratioToUSD: 0.000909,
	}

	result := usd.AddsTo(krw)
	if result.amount != 2200 {
		t.Errorf("Expected 2200, got [%f]", result.amount)
	}

	if result.currency != "KRW" {
		t.Errorf("Exptected KRW, got [%s]", result.currency)
	}
}

type Money struct {
	amount     float64
	currency   string
	ratioToUSD float64
}

func (m Money) Times(n float64) Money {
	return Money{amount: m.amount * n, currency: m.currency}
}

func (m Money) Divides(n float64) Money {
	return Money{amount: m.amount / n, currency: m.currency}
}

func (m Money) Adds(o Money) Money {
	return Money{
		amount:     m.amount + o.amount*o.ratioToUSD,
		currency:   m.currency,
		ratioToUSD: m.ratioToUSD,
	}
}

func (m Money) AddsTo(o Money) Money {
	return Money{
		amount:     math.Round(m.amount/o.ratioToUSD) + o.amount,
		currency:   o.currency,
		ratioToUSD: o.ratioToUSD,
	}
}
