package scriptor_test

import (
	"testing"

	"github.com/death12358/digitalopen/games/scriptor"

	"github.com/shopspring/decimal"
)

// JackpotPush
func JackpotPush(currnecy, name string, point decimal.Decimal) (*scriptor.JackpotPool, error) {
	return testScrtiptor.JackpotPush("test", "test", currnecy, name, point)
}

// JackpotSweep
func JackpotSweep(currnecy, name string, default_point decimal.Decimal) (*scriptor.JackpotPool, error) {
	return testScrtiptor.JackpotSweep("test", "test", currnecy, name, default_point)
}

// JackpotPeek
func JackpotPeek(currnecy, name string) (*scriptor.JackpotPool, error) {
	return testScrtiptor.JackpotPeek("test", "test", currnecy, name)
}

// JackpotPeeks
func JackpotPeeks(currnecy string) (*scriptor.JackpotPools, error) {
	return testScrtiptor.JackpotPeeks("test", "test", currnecy)
}

// JackpotCurrencyPeeks
func JackpotCurrencyPeeks() (*scriptor.JackpotCurrencyPools, error) {
	return testScrtiptor.JackpotCurrencyPeeks("test", "test")
}

// TestJackpotPush
func TestJackpotPush(t *testing.T) {
	res, err := JackpotPush("test", "test", decimal.NewFromFloat(1))
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

// TestJackpotSweep
func TestJackpotSweep(t *testing.T) {
	res, err := JackpotPush("test", "test", decimal.NewFromFloat(1))
	if err != nil {
		t.Error(err)
	}
	t.Log(res)

	res, err = JackpotSweep("test", "test", decimal.Zero)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

// TestJackpotCurrencyPeeks
func TestJackpotCurrencyPeeks(t *testing.T) {
	JackpotPush("AAA", "test", decimal.NewFromFloat(0.000001))
	JackpotPush("AAB", "test", decimal.NewFromFloat(20.01))
	JackpotPush("ABC", "test", decimal.NewFromFloat(65535))
	JackpotPush("ADE", "test", decimal.NewFromFloat(3.1415926))

	res, err := JackpotCurrencyPeeks()
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
