package main

import (
	"math"
	"testing"
)

func TestMultiplication(t *testing.T) {
	USD(5).Times(2).Equal(USD(10), t)
}

func TestMultiplicationInEuro(t *testing.T) {
	EUR(10).Times(2).Equal(EUR(20), t)
}

func TestDivisionKRW(t *testing.T) {
	KRW(4002).Divides(4).Equal(KRW(1000.5), t)
}

func TestAddUsdAndEur(t *testing.T) {
	USD(5).Adds(EUR(10)).Equal(USD(17), t)
}

func TestAddUsdToKrw(t *testing.T) {
	USD(1).AddsTo(KRW(1100)).Equal(KRW(2200.11), t)
}

func TestAddEurAndKrw(t *testing.T) {
	EUR(10).Adds(KRW(4002)).Equal(EUR(13.03), t)
}

func TestAddEurToKrw(t *testing.T) {
	EUR(10).AddsTo(KRW(4002)).Equal(KRW(17203.32), t)
}

func TestAddUsdAndEurToKrw(t *testing.T) {
	USD(10).Adds(EUR(10)).AddsTo(KRW(1000)).Equal(KRW(25202.42), t)
}

type Money struct {
	amount     float64
	currency   string
	ratioToUSD float64
}

func USD(amount float64) Money {
	return Money{
		amount:     amount,
		currency:   "USD",
		ratioToUSD: 1,
	}
}

func EUR(amount float64) Money {
	return Money{
		amount:     amount,
		currency:   "EUR",
		ratioToUSD: 1.2,
	}
}

func KRW(amount float64) Money {
	return Money{
		amount:     amount,
		currency:   "KRW",
		ratioToUSD: 0.000909,
	}
}

func (m Money) Times(n float64) Money {
	return Money{
		amount:     m.amount * n,
		currency:   m.currency,
		ratioToUSD: m.ratioToUSD,
	}
}

func (m Money) Divides(n float64) Money {
	return Money{
		amount:     m.amount / n,
		currency:   m.currency,
		ratioToUSD: m.ratioToUSD,
	}
}

// (ax + by) / a
// ax / a + by / a
// x + by / a
func (m Money) Adds(o Money) Money {
	amount := m.amount + o.amount*o.ratioToUSD/m.ratioToUSD
	amount = math.Round(amount*100) / 100

	return Money{
		amount:     amount,
		currency:   m.currency,
		ratioToUSD: m.ratioToUSD,
	}
}

// (ax + by) / b
// ax / b + by / b
// ax / b + y
func (m Money) AddsTo(o Money) Money {
	amount := m.amount*m.ratioToUSD/o.ratioToUSD + o.amount
	amount = math.Round(amount*100) / 100

	return Money{
		amount:     amount,
		currency:   o.currency,
		ratioToUSD: o.ratioToUSD,
	}
}

func (m Money) Equal(n Money, t *testing.T) {
	if m != n {
		t.Errorf("Expected [%+v], got [%+v]", m, n)
	}
}
