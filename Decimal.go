package ggm

import (
	"math/big"

	"errors"

	"encoding/json"

	"github.com/shopspring/decimal"
)

type DecimalArray []Decimal

type NullDecimal struct {
	Decimal Decimal
	IsValid bool
}

type Decimal struct {
	decimal.Decimal
}

func (d Decimal) MarshalJSON() ([]byte, error) {
	return d.MarshalJSON()
}

func (d *Decimal) UnmarshalJSON(data []byte) error {
	return d.Decimal.UnmarshalJSON(data)
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

func ParseDecimal(x interface{}) (*Decimal, error) {
	switch val := x.(type) {
	case []byte:
		var newDecimal Decimal
		unmarshalling_err := newDecimal.UnmarshalJSON(val)
		return &newDecimal, unmarshalling_err
	case float64:
		return NewDecimal(val), nil
	case float32:
		return NewDecimal(float64(val)), nil
	case int:
		return NewDecimal(float64(val)), nil
	case int32:
		return NewDecimal(float64(val)), nil
	case int64:
		return NewDecimal(float64(val)), nil
	case json.Number:
		return NewDecimalFromString(string(val))
	case string:
		return NewDecimalFromString(val)
	case Decimal:
		return &val, nil
	case big.Float:
		return NewDecimalFromString(val.String())
	}
	return nil, errors.New("Unsupported type to parse")
}

func NewDecimal(x float64) *Decimal {
	return &Decimal{decimal.NewFromFloat(x)}
}
func NewDecimalFromString(str string) (*Decimal, error) {
	dec, err := decimal.NewFromString(str)
	return &Decimal{dec}, err
}
func NewDecimalFromBigInt(bigInt *big.Int, exp int32) *Decimal {
	return &Decimal{decimal.NewFromBigInt(bigInt, exp)}
}
func NewFromFloatWithExponent(val float64, exp int32) *Decimal {
	return &Decimal{decimal.NewFromFloatWithExponent(val, exp)}
}
