package ggm

import (
	"github.com/shopspring/decimal"
)

type Decimal struct {
	decimal.Decimal
}
func (d *Decimal) Add(x Decimal) Decimal {
	return Decimal{d.Decimal.Add(x.Decimal)}
}
func (d *Decimal) AddFloat(x float64) Decimal {
	*d = Decimal{d.Decimal.Add(decimal.NewFromFloat(x))}
	return *d
}
func (d *Decimal) Multiply(x Decimal) Decimal {
	return Decimal{d.Decimal.Mul(x.Decimal)}
}
func (d *Decimal) MultiplyFloat(x float64) Decimal {
	*d = Decimal{d.Decimal.Mul(decimal.NewFromFloat(x))}
	return *d
}
func (d *Decimal) Divide(x Decimal) Decimal {
	return Decimal{d.Decimal.Div(x.Decimal)}
}
func (d *Decimal) DivideFloat(x float64) Decimal {
	*d = Decimal{d.Decimal.Div(decimal.NewFromFloat(x))}
	return *d
}
func (d *Decimal) Subtract(x Decimal) Decimal {
	return Decimal{d.Decimal.Sub(x.Decimal)}
}
func (d *Decimal) Sub(x Decimal) Decimal {
	return d.Subtract(x)
}
func (d *Decimal) SubtractFloat(x float64) Decimal {
	*d = Decimal{d.Decimal.Sub(decimal.NewFromFloat(x))}
	return *d
}
func (d *Decimal) SubFloat(x float64) Decimal {
	return d.SubtractFloat(x)
}

func (d Decimal) EqualFloat(x float64) bool {
	return d.Equal(NewDecimal(x).Decimal)
}

type Money = Decimal

func NewDecimal(x float64) Decimal {
	return Decimal{decimal.NewFromFloat(x)}
}
func NewMoney(x float64) Money {
	return Money{decimal.NewFromFloat(x)}
}