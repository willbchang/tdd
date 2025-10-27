package tdd

import (
	"testing"

	. "willbc.com/tdd/stock"
)

func TestMultiplication(t *testing.T) {
	EUR(10).Times(2).Equal(EUR(20), t)
}

func TestDivisionKRW(t *testing.T) {
	KRW(4002).Divides(4).Equal(KRW(1000.5), t)
}

func TestAddUsdAndEur(t *testing.T) {
	USD(5).Adds(EUR(10)).Equal(USD(17), t)
}

func TestAddUsdToKrw(t *testing.T) {
	KRW(1100).Adds(USD(1)).Equal(KRW(2200.11), t)
}

func TestAddEurAndKrw(t *testing.T) {
	EUR(10).Adds(KRW(4002)).Equal(EUR(13.03), t)
}

func TestAddUsdAndEurToKrw(t *testing.T) {
	KRW(1000).Adds(USD(10)).Adds(EUR(10)).Equal(KRW(25202.42), t)
}
