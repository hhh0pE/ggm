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
