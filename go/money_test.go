package main

import (
	"math"
	"testing"
)

func TestMultiplication(t *testing.T) {
	fiver := USD(5)
	actualResult := fiver.Times(2)
	expectedResult := USD(10)
	expectedResult.Equal(t, actualResult)
}

func TestMultiplicationInEuro(t *testing.T) {
	money := EUR(10)
	actualResult := money.Times(2)
	expectedResult := EUR(20)
	expectedResult.Equal(t, actualResult)
}

func TestDivisionKRW(t *testing.T) {
	money := KRW(4002)
	actualResult := money.Divides(4)
	expectedResult := KRW(1000.5)
	expectedResult.Equal(t, actualResult)
}

func TestAddUsdAndEur(t *testing.T) {
	usd := USD(5)
	eur := EUR(10)
	actualResult := usd.Adds(eur)
	expectedResult := USD(17)
	expectedResult.Equal(t, actualResult)

}

func TestAddUsdAndKrwOutputKrw(t *testing.T) {
	usd := USD(1)
	krw := KRW(1100)
	actualResult := usd.AddsTo(krw)
	expectedResult := KRW(2200.11)
	expectedResult.Equal(t, actualResult)
}

func TestAddEurAndKrw(t *testing.T) {
	eur := EUR(10)
	krw := KRW(4002)
	actualResult := eur.Adds(krw)
	expectedResult := EUR(13.03)
	expectedResult.Equal(t, actualResult)
}

func TestAddEurToKrw(t *testing.T) {
	eur := EUR(10)
	krw := KRW(4002)
	actualResult := eur.AddsTo(krw)
	expectedResult := KRW(17203.32)
	expectedResult.Equal(t, actualResult)
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

func (m Money) Adds(o Money) Money {
	amount := m.amount + o.amount*o.ratioToUSD/m.ratioToUSD
	amount = math.Round(amount*100) / 100

	return Money{
		amount:     amount,
		currency:   m.currency,
		ratioToUSD: m.ratioToUSD,
	}
}

func (m Money) AddsTo(o Money) Money {
	amount := m.amount*m.ratioToUSD/o.ratioToUSD + o.amount
	amount = math.Round(amount*100) / 100

	return Money{
		amount:     amount,
		currency:   o.currency,
		ratioToUSD: o.ratioToUSD,
	}
}

func (m Money) Equal(t *testing.T, n Money) {
	if m != n {
		t.Errorf("Expected [%+v], got [%+v]", m, n)
	}
}
