package money

import (
	"database/sql/driver"
	"errors"

	mon "github.com/Rhymond/go-money"
)

type Money struct {
	Amount mon.Money
}

func (a *Money) String() string {
	return a.Amount.Display()
}

func (a *Money) MarshalJSON() ([]byte, error) {
	return mon.MarshalJSON(a.Amount)
}

func (a *Money) Value() (driver.Value, error) {
	return mon.MarshalJSON(a.Amount)
}

func (a *Money) Scan(val any) error {
	moneyByte, ok := val.([]byte)
	if !ok {
		return errors.New("could not scan money | assertion failed")
	}
	return a.Amount.UnmarshalJSON(moneyByte)
}

func (a *Money) UnmarshalJSON(data []byte) error {
	return a.Amount.UnmarshalJSON(data)
}
