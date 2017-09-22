package fieldType

import (
	"fmt"

	"github.com/hhh0pE/ggm/ggmgen/templates/fields"
)

type ConstFieldType uint

type FieldType struct {
	ConstType  ConstFieldType
	Size       uint
	MaxSize    uint
	IsNullable bool
}

func (ft FieldType) Template() string {
	switch {
	case ft.ConstType == IntType && !ft.IsNullable:
		return fields.IntegerTemplate
	case ft.ConstType == IntType && ft.IsNullable:
		return fields.IntegerNullableTemplate

	case ft.ConstType == BoolType && !ft.IsNullable:
		return fields.BooleanTemplate
	case ft.ConstType == BoolType && ft.IsNullable:
		return fields.BooleanNullableTemplate

	case ft.ConstType == TextType && !ft.IsNullable:
		return fields.StringTemplate
	case ft.ConstType == TextType && ft.IsNullable:
		return fields.StringNullableTemplate

	case ft.ConstType == FloatType && !ft.IsNullable:
		return fields.FloatTemplate
	case ft.ConstType == FloatType && ft.IsNullable:
		return fields.FloatNullableTemplate

	case ft.ConstType == DateType && !ft.IsNullable:
		return fields.DateTemplate
	case ft.ConstType == DateType && ft.IsNullable:
		return fields.DateNullableTemplate
	}
	return ""
}

const (
	BoolType ConstFieldType = iota + 1
	IntType
	FloatType
	TextType
	DateType
)

func (ft FieldType) SqlType() string {
	var sqlType string
	switch ft.ConstType {
	case BoolType:
		sqlType = "BOOLEAN"
	case IntType:
		sqlType = "INTEGER"
	case FloatType:
		sqlType = "REAL"
	case TextType:
		if ft.MaxSize > 0 && ft.MaxSize < 255 {
			sqlType = fmt.Sprintf("VARCHAR(%d)", ft.MaxSize)
		} else {
			sqlType = "TEXT"
		}
	case DateType:
		sqlType = "TIMESTAMP"
	default:
		return ""
	}

	if ft.IsNullable {
		sqlType += " NULL"
	} else {
		sqlType += " NOT NULL"
	}
	return sqlType
}

func (ft FieldType) Name() string {
	var name string
	switch ft.ConstType {
	case IntType:
		name = "Integer"
	case FloatType:
		name = "Float"
	case TextType:
		name = "String"
	case BoolType:
		name = "Boolean"
	case DateType:
		name = "Date"
	}

	if ft.IsNullable {
		name += "Nullable"
	}
	return name
}

func (ft FieldType) FmtReplacer() string {
	switch ft.ConstType {
	case IntType:
		return "%d"
	case FloatType:
		return "%f"
	case TextType:
		return "%s"
	case BoolType:
		return "%t"
	case DateType:
		return "%s"
	}
	return ""
}

func (ft FieldType) DefaultValue() string {
	switch ft.ConstType {
	case IntType:
		return "0"
	case FloatType:
		return "0.0"
	case TextType:
		return "\"\""
	case BoolType:
		return "false"
	case DateType:
		return "0"
	}
	return ""
}

func GetAllFieldTypes() []FieldType {
	return []FieldType{
		Boolean,
		BooleanNullable,
		Date,
		DateNullable,
		Integer,
		IntegerNullable,
		Numeric,
		NumericNullable,
		Text,
		TextNullable,
		VarChar,
		VarCharNullable,
	}
}
