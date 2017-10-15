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
{{if not .}}
type whereFieldBoolean struct {
	name string
	where modelWhere
}
func (wfb whereFieldBoolean) sqlName() string {
	return "\""+wfb.name+"\""
}
func(wfb whereFieldBoolean) is(val bool) modelWhere {
	wfb.where.andOr()
	if val {
		wfb.where.addCond(wfb.sqlName()+" = 'TRUE'")
	} else {
		wfb.where.addCond(wfb.sqlName()+" = 'FALSE'")
	}
	return wfb.where
}
func (wfb whereFieldBoolean) isTrue() modelWhere {
	wfb.where.andOr()
	wfb.where.addCond(wfb.sqlName()+" = 'TRUE'")
	return wfb.where
}
func (wfb whereFieldBoolean) isFalse() modelWhere {
	wfb.where.andOr()
	wfb.where.addCond(wfb.sqlName()+" = 'FALSE'")
	return wfb.where
}
func (wfb whereFieldBoolean) fromStr(str string) modelWhere {
	var val bool
	str = strings.TrimSpace(strings.ToLower(str))
	if str == "1" || str == "true" || str == "t" || str == "y" || str == "on" || str == "yes" {
		val = true
	}
	return wfb.is(val)
}
func (wfb whereFieldBoolean) isNull() modelWhere {
	wfb.where.andOr()
	wfb.where.addCond(wfb.sqlName()+" IS NULL")
	return wfb.where
}
func (wfb whereFieldBoolean) isNotNull() modelWhere {
	wfb.where.andOr()
	wfb.where.addCond(wfb.sqlName()+" IS NOT NULL")
	return wfb.where
}
{{else}}
type whereFieldBoolean{{.ModelName}} struct {
    whereFieldBoolean
}
func (wfb whereFieldBoolean{{.ModelName}}) sqlName() string {
	return "\"{{.ModelTableName}}\".\""+wfb.name+"\""
}

func (wfb whereFieldBoolean{{.ModelName}}) Is(val bool) *{{lower .ModelName}}Where {
	return wfb.is(val).(*{{lower .ModelName}}Where)
}
func (wfb whereFieldBoolean{{.ModelName}}) IsTrue() *{{lower .ModelName}}Where {
	return wfb.isTrue().(*{{lower .ModelName}}Where)
}
func (wfb whereFieldBoolean{{.ModelName}}) IsFalse() *{{lower .ModelName}}Where {
	return wfb.isFalse().(*{{lower .ModelName}}Where)
}
func (wfb whereFieldBoolean{{.ModelName}}) FromStr(str string) *{{lower .ModelName}}Where {
	return wfb.fromStr(str).(*{{lower .ModelName}}Where)
}
{{end}}
`

const BooleanNullableTemplate = `
{{if .}}
type whereFieldBooleanNullable{{.ModelName}} struct {
	whereFieldBoolean{{.ModelName}}
}
func(wfbn whereFieldBooleanNullable{{.ModelName}}) sqlName() string {
	return "\"{{.ModelTableName}}\".\""+wfbn.name+"\""
}
func (wfbn whereFieldBooleanNullable{{.ModelName}}) IsNull() *{{lower .ModelName}}Where {
	return wfbn.isNull().(*{{lower .ModelName}}Where)
}
func (wfbn whereFieldBooleanNullable{{.ModelName}}) IsNotNull() *{{lower .ModelName}}Where {
	return wfbn.isNotNull().(*{{lower .ModelName}}Where)
}
{{end}}
`

const BooleanArrayTemplate = `
{{if not .}}
type whereFieldBooleanArray struct {
	name string
	where modelWhere
}
func(wfba whereFieldBooleanArray) sqlName() string {
	return "\""+wfba.name+"\""
}
func(wfba whereFieldBooleanArray) is(val []bool) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" = '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray) isNot(val []bool) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" <> '"+boolArrayToSqlValue(val)+"'")
return wfba.where
}
func(wfba whereFieldBooleanArray) lessThan(val []bool) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" < '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray) lessThanOrEqual(val []bool) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" <= '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray) greaterThan(val []bool) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" > '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray) greaterThanOrEqual(val []bool) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" > '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray) contains(val bool) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" @> '"+boolArrayToSqlValue([]bool{val})+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray) containedBy(val []bool) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" <@ '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray) overlap(val []bool) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" && '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray) lengthIs(len int) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond("array_length("+wfba.sqlName()+", 1) = '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray) lengthLessThan(len int) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond("array_length("+wfba.sqlName()+", 1) < '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray) lengthLessThanOrEqual(len int) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond("array_length("+wfba.sqlName()+", 1) <= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray) lengthGreaterThan(len int) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond("array_length("+wfba.sqlName()+", 1) > '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray) lengthGreaterThanOrEqual(len int) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond("array_length("+wfba.sqlName()+", 1) >= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
{{else}}
type whereFieldBooleanArray{{.ModelName}} struct {
	whereFieldBooleanArray
}
func (wfba whereFieldBooleanArray{{.ModelName}}) sqlName() string {
	return "\"{{.ModelTableName}}\".\""+wfba.name+"\""
}
func(wfba whereFieldBooleanArray{{.ModelName}}) Is(val []bool) *{{lower .ModelName}}Where {
	return wfba.is(val).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) IsNot(val []bool) *{{lower .ModelName}}Where {
	return wfba.isNot(val).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LessThan(val []bool) *{{lower .ModelName}}Where {
	return wfba.lessThan(val).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LT(val []bool) *{{lower .ModelName}}Where {
	return wfba.LessThan(val)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LessThanOrEqual(val []bool) *{{lower .ModelName}}Where {
	return wfba.lessThanOrEqual(val).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LTE(val []bool) *{{lower .ModelName}}Where {
	return wfba.LessThanOrEqual(val)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) GreaterThan(val []bool) *{{lower .ModelName}}Where {
	return wfba.greaterThan(val).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) GT(val []bool) *{{lower .ModelName}}Where {
	return wfba.GreaterThan(val)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) GreaterThanOrEqual(val []bool) *{{lower .ModelName}}Where {
	return wfba.greaterThanOrEqual(val).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) GTE(val []bool) *{{lower .ModelName}}Where {
	return wfba.GreaterThanOrEqual(val)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) Contains(val bool) *{{lower .ModelName}}Where {
	return wfba.contains(val).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) ContainedBy(val []bool) *{{lower .ModelName}}Where {
	return wfba.containedBy(val).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) Overlap(val []bool) *{{lower .ModelName}}Where {
	return wfba.overlap(val).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthIs(len int) *{{lower .ModelName}}Where {
	return wfba.lengthIs(len).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthLessThan(len int) *{{lower .ModelName}}Where {
	return wfba.lengthLessThan(len).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthLT(len int) *{{lower .ModelName}}Where {
	return wfba.LengthLessThan(len)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthLessThanOrEqual(len int) *{{lower .ModelName}}Where {
	return wfba.lengthLessThanOrEqual(len).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthLTE(len int) *{{lower .ModelName}}Where {
	return wfba.LengthLessThanOrEqual(len)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthGreaterThan(len int) *{{lower .ModelName}}Where {
	return wfba.lengthGreaterThan(len).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthGT(len int) *{{lower .ModelName}}Where {
	return wfba.LengthGreaterThan(len)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthGreaterThanOrEqual(len int) *{{lower .ModelName}}Where {
	return wfba.lengthGreaterThanOrEqual(len).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthGTE(len int) *{{lower .ModelName}}Where {
	return wfba.LengthGreaterThanOrEqual(len)
}
{{end}}
`
