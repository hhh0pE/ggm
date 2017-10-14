package fieldType

type ConstFieldType uint

const (
	BoolType ConstFieldType = iota + 1
	IntType
	FloatType
	DecimalType
	TextType
	DateType
)

func (cft ConstFieldType) BaseType(isNullable, isArray, isGoBaseType bool) FieldType {
	switch cft {
	case BoolType:
		var ft Boolean
		ft.IsNullable = isNullable
		ft.IsArray = isArray
		ft.IsGoBaseType = isGoBaseType
		return &ft
	case IntType:
		var ft Integer
		ft.IsNullable = isNullable
		ft.IsArray = isArray
		ft.IsGoBaseType = isGoBaseType
		return &ft
	case FloatType:
		var ft Float
		ft.IsNullable = isNullable
		ft.IsArray = isArray
		ft.IsGoBaseType = isGoBaseType
		return &ft
	case TextType:
		var ft Text
		ft.IsNullable = isNullable
		ft.IsArray = isArray
		ft.IsGoBaseType = isGoBaseType
		return &ft
	case DateType:
		var ft Date
		ft.IsNullable = isNullable
		ft.IsArray = isArray
		ft.IsGoBaseType = isGoBaseType
		return &ft
	case DecimalType:
		var ft Decimal
		ft.IsNullable = isNullable
		ft.IsArray = isArray
		ft.IsGoBaseType = isGoBaseType
		return &ft
	}
	return nil
}

type FieldType interface {
	Name() string
	SqlType() string
	WhereTemplate() string
	FmtReplacer() string
	DefaultValue() string
	GoBaseType() string
	GoScannerType() string

	MaxSize() int
	Nullable() bool
	Array() bool

	ConstType() ConstFieldType
	ImplementScannerInterface() bool
}

func sqlType(sqlTypeName string, ft FieldType) string {
	if ft.Array() {
		return sqlTypeName + "[]"
	}
	if ft.Nullable() {
		return sqlTypeName + " NOT NULL"
	}
	return sqlTypeName
}

func name(name string, ft FieldType) string {
	if ft.Array() {
		return name + "Array"
	}
	if ft.Nullable() {
		return name + "Nullable"
	}
	return name
}

func goType(typeName string, ft FieldType) string {
	if ft.Array() {
		if ft.Nullable() {
			return "*[]" + typeName
		} else {
			return "[]" + typeName
		}
	}
	if ft.Nullable() {
		return "*" + typeName
	}
	return typeName
}

//type FieldType struct {
//	ConstType  ConstFieldType
//	Size       uint
//	MaxSize    uint
//	IsNullable bool
//	IsArray    bool
//}

//func (ft FieldType) Template() string {
//	switch {
//	case ft.ConstType == IntType:
//		if ft.IsArray {
//			return fields.IntegerArrayTemplate
//		}
//		if ft.IsNullable {
//			return fields.IntegerNullableTemplate
//		}
//		return fields.IntegerTemplate
//	case ft.ConstType == FloatType:
//		if ft.IsArray {
//			return fields.FloatArrayTemplate
//		}
//		if ft.IsNullable {
//			return fields.FloatNullableTemplate
//		}
//		return fields.FloatTemplate
//	case ft.ConstType == DecimalType:
//		if ft.IsArray {
//			return fields.DecimalArrayTemplate
//		}
//		if ft.IsNullable {
//			return fields.DecimalNullableTemplate
//		}
//		return fields.DecimalTemplate
//	case ft.ConstType == BoolType:
//		if ft.IsArray {
//			return fields.BooleanArrayTemplate
//		}
//		if ft.IsNullable {
//			return fields.BooleanNullableTemplate
//		}
//		return fields.BooleanTemplate
//	case ft.ConstType == TextType:
//		if ft.IsArray {
//			return fields.StringArrayTemplate
//		}
//		if ft.IsNullable {
//			return fields.StringNullableTemplate
//		}
//		return fields.StringTemplate
//	case ft.ConstType == DateType:
//		if ft.IsArray {
//
//		}
//		if ft.IsNullable {
//			return fields.DateNullableTemplate
//		}
//		return fields.DateTemplate
//	}
//	return ""
//}

//func (ft FieldType) SqlType() string {
//	var sqlType string
//	switch ft.ConstType {
//	case BoolType:
//		sqlType = "BOOLEAN"
//	case IntType:
//		sqlType = "BIGINT"
//	case FloatType:
//		sqlType = "REAL"
//	case DecimalType:
//		sqlType = "NUMERIC"
//	case TextType:
//		if ft.MaxSize > 0 && ft.MaxSize <= 255 {
//			sqlType = fmt.Sprintf("VARCHAR(%d)", ft.MaxSize)
//		} else {
//			sqlType = "TEXT"
//		}
//	case DateType:
//		sqlType = "TIMESTAMP WITH TIME ZONE"
//	default:
//		return ""
//	}
//
//	if ft.IsArray {
//		sqlType += "[]"
//	}
//
//	if ft.IsNullable {
//		sqlType += " NULL"
//	} else {
//		sqlType += " NOT NULL"
//	}
//
//	return sqlType
//}
//
//func (ft FieldType) Name() string {
//	var name string
//	switch ft.ConstType {
//	case IntType:
//		name = "Integer"
//	case FloatType:
//		name = "Float"
//	case DecimalType:
//		name = "Decimal"
//	case TextType:
//		name = "String"
//	case BoolType:
//		name = "Boolean"
//	case DateType:
//		name = "Date"
//	}
//
//	if ft.IsArray {
//		name += "Array"
//	} else if ft.IsNullable {
//		name += "Nullable"
//	}
//	return name
//}
//
//func (ft FieldType) FmtReplacer() string {
//	switch ft.ConstType {
//	case IntType:
//		return "%d"
//	case FloatType:
//		return "%f"
//	case TextType:
//		return "%s"
//	case BoolType:
//		return "%t"
//	case DateType:
//		return "%s"
//	case DecimalType:
//		return "%s"
//	}
//	return ""
//}
//
//func (ft FieldType) DefaultValue() string {
//	if ft.IsArray {
//		return "nil"
//	}
//	switch ft.ConstType {
//	case IntType:
//		return "0"
//	case FloatType:
//		return "0.0"
//	case DecimalType:
//		return "ggm.NewDecimal(0)"
//	case TextType:
//		return "\"\""
//	case BoolType:
//		return "false"
//	case DateType:
//		return "0"
//	}
//	return ""
//}
//
func GetAllFieldTypes() []FieldType {
	return []FieldType{
		BoolType.BaseType(false, false, false),
		BoolType.BaseType(true, false, false),
		BoolType.BaseType(false, true, false),
		IntType.BaseType(false, false, false),
		IntType.BaseType(true, false, false),
		IntType.BaseType(false, true, false),
		FloatType.BaseType(false, false, false),
		FloatType.BaseType(true, false, false),
		FloatType.BaseType(false, true, false),
		DecimalType.BaseType(false, false, false),
		DecimalType.BaseType(true, false, false),
		DecimalType.BaseType(false, true, false),
		TextType.BaseType(false, false, false),
		TextType.BaseType(true, false, false),
		TextType.BaseType(false, true, false),
		DateType.BaseType(false, false, false),
		DateType.BaseType(true, false, false),
		DateType.BaseType(false, true, false),
	}
}
