package stock

import (
	"math"
	"testing"
)

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
		amount:     math.Round(m.amount*n*100) / 100,
		currency:   m.currency,
		ratioToUSD: m.ratioToUSD,
	}
}

func (m Money) Divides(n float64) Money {
	return Money{
		amount:     math.Round(m.amount/n*100) / 100,
		currency:   m.currency,
		ratioToUSD: m.ratioToUSD,
	}
}

// Adds (ax + by) / a ->
// ax / a + by / a ->
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

func (m Money) Equal(n Money, t *testing.T) {
	if m != n {
		t.Errorf("Expected [%+v], got [%+v]", m, n)
	}
}
