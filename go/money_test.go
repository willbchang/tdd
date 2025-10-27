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
	actualResult := fiver.Times(2)
	expectedResult := Money{
		amount:   10,
		currency: "USD",
	}
	expectedResult.Equal(t, actualResult)
}

func TestMultiplicationInEuro(t *testing.T) {
	money := Money{
		amount:   10,
		currency: "EUR",
	}
	actualResult := money.Times(2)
	expectedResult := Money{
		amount:   20,
		currency: "EUR",
	}
	expectedResult.Equal(t, actualResult)
}

func TestDivisionKRW(t *testing.T) {
	money := Money{
		amount:   4002,
		currency: "KRW",
	}
	actualResult := money.Divides(4)
	expectedResult := Money{
		amount:   1000.5,
		currency: "KRW",
	}
	expectedResult.Equal(t, actualResult)
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

	actualResult := usd.Adds(eur)
	expectedResult := Money{17, "USD", 1}
	expectedResult.Equal(t, actualResult)

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

	actualResult := usd.AddsTo(krw)
	expectedResult := Money{2200, "KRW", 0.000909}
	expectedResult.Equal(t, actualResult)
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

func (m Money) Equal(t *testing.T, n Money) {
	if m != n {
		t.Errorf("Expected [%+v], got [%+v]", m, n)
	}
}
