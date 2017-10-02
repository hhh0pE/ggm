package fieldType

type Boolean struct {
	IsNullable   bool
	IsArray      bool
	IsGoBaseType bool
}

func (bt Boolean) SqlType() string {
	return sqlType("BOOLEAN", bt)
}

func (bt Boolean) WhereTemplate() string {
	if bt.IsArray {
		return BooleanArrayTemplate
	}
	if bt.IsNullable {
		return BooleanNullableTemplate
	}
	return BooleanTemplate
}

func (bt Boolean) FmtReplacer() string {
	return "%t"
}

func (bt Boolean) DefaultValue() string {
	//if bt.IsArray || bt.IsNullable {
	//	return "nil"
	//}
	return "false"
}

func (bt Boolean) GoBaseType() string {
	if bt.IsGoBaseType {
		return goType("bool", bt)
	}
	return bt.GoScannerType()
	//return goType("bool", bt)
}

func (bt Boolean) GoScannerType() string {
	//if bt.IsGoBaseType {
	//	return bt.GoBaseType()
	//}
	if bt.IsArray {
		if bt.IsNullable {
			return "*pq.BoolArray"
		}
		return "pq.BoolArray"
	}
	if bt.IsNullable {
		return "sql.NullBool"
	}

	return "bool"
}
func (bt Boolean) ImplementScannerInterface() bool {
	return bt.GoBaseType() == bt.GoScannerType()
}

//func (bt Boolean) IsNilSuffix() string {
//	if bt.IsArray {
//		return " == nil"
//	}
//	if bt.IsNullable {
//		return ".Valid"
//	}
//	return ""
//}

//func (bt Boolean) Initializer() string {
//	var b sql.NullBool
//	if b.Valid {
//
//	}
//}

func (bt Boolean) Size() int {
	return 0
}

func (bt Boolean) MaxSize() int {
	return 0
}

func (bt Boolean) Nullable() bool {
	return bt.IsNullable
}

func (bt Boolean) Array() bool {
	return bt.IsArray
}

func (bt Boolean) Name() string {
	return name("Boolean", bt)
}
func (bt Boolean) ConstType() ConstFieldType {
	return BoolType
}

const BooleanTemplate = `
type whereFieldBoolean{{.ModelName}} struct {
    name string
	where *{{lower .ModelName}}Where
}

func (wfb whereFieldBoolean{{.ModelName}}) Is(val bool) *{{lower .ModelName}}Where {
	wfb.where.andOr()
	if val {
		wfb.where.addCond("\"" + wfb.name + "\" = 'TRUE'")
	} else {
		wfb.where.addCond("\"" + wfb.name + "\" = 'FALSE'")
	}
	return wfb.where
}
func (wfb whereFieldBoolean{{.ModelName}}) IsTrue() *{{lower .ModelName}}Where {
	wfb.where.andOr()
	wfb.where.addCond("\"" + wfb.name + "\" = 'TRUE'")
	return wfb.where
}
func (wfb whereFieldBoolean{{.ModelName}}) IsFalse() *{{lower .ModelName}}Where {
	wfb.where.andOr()
	wfb.where.addCond("\"" + wfb.name + "\" = 'FALSE'")
	return wfb.where
}
func (wfb whereFieldBoolean{{.ModelName}}) FromStr(str string) *{{lower .ModelName}}Where {
	var val bool
	str = strings.TrimSpace(strings.ToLower(str))
	if str == "1" || str == "true" || str == "t" || str == "y" || str == "on" || str == "yes" {
		val = true
	}
	return wfb.Is(val)
}
`

const BooleanNullableTemplate = `
type whereFieldBooleanNullable{{.ModelName}} struct {
	whereFieldBoolean{{.ModelName}}
}

func (wfbn whereFieldBooleanNullable{{.ModelName}}) IsNull() *{{lower .ModelName}}Where {
	wfbn.where.andOr()
	wfbn.where.addCond("\"" + wfbn.name + "\" IS NULL")
	return wfbn.where
}
func (wfbn whereFieldBooleanNullable{{.ModelName}}) IsNotNull() *{{lower .ModelName}}Where {
	wfbn.where.andOr()
	wfbn.where.addCond("\"" + wfbn.name + "\" IS NOT NULL")
	return wfbn.where
}
`

const BooleanArrayTemplate = `
type whereFieldBooleanArray{{.ModelName}} struct {
	name string
	where *{{lower .ModelName}}Where
}

func(wfba whereFieldBooleanArray{{.ModelName}}) Is(val []bool) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" = '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) IsNot(val []bool) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <> '"+boolArrayToSqlValue(val)+"'")
return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LessThan(val []bool) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" < '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LT(val []bool) *{{lower .ModelName}}Where {
	return wfba.LessThan(val)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LessThanOrEqual(val []bool) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <= '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LTE(val []bool) *{{lower .ModelName}}Where {
	return wfba.LessThanOrEqual(val)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) GreaterThan(val []bool) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" > '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) GT(val []bool) *{{lower .ModelName}}Where {
	return wfba.GreaterThan(val)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) GreaterThanOrEqual(val []bool) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" > '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) GTE(val []bool) *{{lower .ModelName}}Where {
	return wfba.GreaterThanOrEqual(val)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) Contains(val bool) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" @> '"+boolArrayToSqlValue([]bool{val})+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) ContainedBy(val []bool) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <@ '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) Overlap(val []bool) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" && '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthIs(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) = '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthLessThan(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) < '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthLT(len int) *{{lower .ModelName}}Where {
	return wfba.LengthLessThan(len)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthLessThanOrEqual(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) <= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthLTE(len int) *{{lower .ModelName}}Where {
	return wfba.LengthLessThanOrEqual(len)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthGreaterThan(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) > '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthGT(len int) *{{lower .ModelName}}Where {
	return wfba.LengthGreaterThan(len)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthGreaterThanOrEqual(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) >= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthGTE(len int) *{{lower .ModelName}}Where {
	return wfba.LengthGreaterThanOrEqual(len)
}
`
