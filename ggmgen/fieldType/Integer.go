package fieldType

type Integer struct {
	IsNullable bool
	IsArray bool
	IsGoBaseType bool
}

func (i Integer) Name() string {
	return name("Integer", i)
}

func (i Integer) SqlType() string {
	return sqlType("INTEGER", i)
}

func (i Integer) WhereTemplate() string {
	if i.IsArray {
		return IntegerArrayTemplate
	}
	if i.IsNullable {
		return IntegerNullableTemplate
	}
	return IntegerTemplate
}

func (i Integer) FmtReplacer() string {
	return "%d"
}

func (i Integer) DefaultValue() string {
	return "0"
}

func (i Integer) GoBaseType() string {
	if i.IsGoBaseType {
		return goType("int64", i)
	}
	return i.GoScannerType()
}

func (i Integer) GoScannerType() string {
	if i.IsArray {
		if i.IsNullable {
			return "*pq.Int64Array"
		} else {
			return "pq.Int64Array"
		}
	}
	if i.IsNullable {
		return "sql.NullInt64"
	}
	return "int64"
}

func (i Integer) Size() int {
	return 0
}

func (i Integer) MaxSize() int {
	return 0
}

func (i Integer) Nullable() bool {
	return i.IsNullable
}

func (i Integer) Array() bool {
	return i.IsArray
}
func (i Integer) ConstType() ConstFieldType {
	return IntType
}
func (i Integer) ImplementScannerInterface() bool {
	return i.GoBaseType() == i.GoScannerType()
}



const IntegerTemplate = `
type whereFieldInteger{{.ModelName}} struct {
	name  string
	where *{{lower .ModelName}}Where
}

func (wfi whereFieldInteger{{.ModelName}}) Is(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" = '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) Eq(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" = '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) Equal(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" = '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) IsNot(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" <> '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) GreaterThan(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" > '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) GT(val int) *{{lower .ModelName}}Where {
	return wfi.GreaterThan(val)
}
func (wfi whereFieldInteger{{.ModelName}}) GreaterThanOrEqual(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" >= '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) GTE(val int) *{{lower .ModelName}}Where {
	return wfi.GreaterThanOrEqual(val)
}
func (wfi whereFieldInteger{{.ModelName}}) LessThan(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" < '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) LT(val int) *{{lower .ModelName}}Where {
	return wfi.LessThan(val)
}
func (wfi whereFieldInteger{{.ModelName}}) LessThanOrEqual(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" <= '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) LTE(val int) *{{lower .ModelName}}Where {
	return wfi.LessThanOrEqual(val)
}

func (wfi whereFieldInteger{{.ModelName}}) Between(left, right int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" BETWEEN '%d' AND '%d'", left, right))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) NotBetween(left, right int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" NOT BETWEEN '%d' AND '%d'", left, right))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) In(nums []int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	var cond string
	cond += "\"" + wfi.name + "\" IN ("
	for i, num := range nums {
		cond += fmt.Sprintf("'%d'", num)
		if i != len(nums) {
			cond += ", "
		}
	}
	cond += ")"
	wfi.where.addCond(cond)
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) Any(nums ...int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	var cond string
	cond += "\"" + wfi.name + "\" IN ("
	for i, num := range nums {
		cond += fmt.Sprintf("'%d'", num)
		if i != len(nums) {
			cond += ", "
		}
	}
	cond += ")"
	wfi.where.addCond(cond)
	return wfi.where
}
`


const IntegerNullableTemplate = `
type whereFieldIntegerNullable{{.ModelName}} struct {
	whereFieldInteger{{.ModelName}}
}

func (wfin whereFieldIntegerNullable{{.ModelName}}) IsNull() *{{lower .ModelName}}Where {
	wfin.where.andOr()
	wfin.where.addCond("\"" + wfin.name + "\" IS NULL")
	return wfin.where
}
func (wfin whereFieldIntegerNullable{{.ModelName}}) IsNotNull() *{{lower .ModelName}}Where {
	wfin.where.andOr()
	wfin.where.addCond("\"" + wfin.name + "\" IS NOT NULL")
	return wfin.where
}
`

const IntegerArrayTemplate = `
type whereFieldIntegerArray{{.ModelName}} struct {
	name string
	where *{{lower .ModelName}}Where
}

func(wfba whereFieldIntegerArray{{.ModelName}}) Is(val []int64) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" = '"+int64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldIntegerArray{{.ModelName}}) IsNot(val []int64) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <> '"+int64ArrayToSqlValue(val)+"'")
return wfba.where
}
func(wfba whereFieldIntegerArray{{.ModelName}}) LessThan(val []int64) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" < '"+int64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldIntegerArray{{.ModelName}}) LT(val []int64) *{{lower .ModelName}}Where {
	return wfba.LessThan(val)
}
func(wfba whereFieldIntegerArray{{.ModelName}}) LessThanOrEqual(val []int64) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <= '"+int64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldIntegerArray{{.ModelName}}) LTE(val []int64) *{{lower .ModelName}}Where {
	return wfba.LessThanOrEqual(val)
}
func(wfba whereFieldIntegerArray{{.ModelName}}) GreaterThan(val []int64) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" > '"+int64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldIntegerArray{{.ModelName}}) GT(val []int64) *{{lower .ModelName}}Where {
	return wfba.GreaterThan(val)
}
func(wfba whereFieldIntegerArray{{.ModelName}}) GreaterThanOrEqual(val []int64) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" > '"+int64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldIntegerArray{{.ModelName}}) GTE(val []int64) *{{lower .ModelName}}Where {
	return wfba.GreaterThanOrEqual(val)
}
func(wfba whereFieldIntegerArray{{.ModelName}}) Contains(val int64) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" @> '"+int64ArrayToSqlValue([]int64{val})+"'")
	return wfba.where
}
func(wfba whereFieldIntegerArray{{.ModelName}}) ContainedBy(val []int64) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <@ '"+int64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldIntegerArray{{.ModelName}}) Overlap(val []int64) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" && '"+int64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldIntegerArray{{.ModelName}}) LengthIs(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) = '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldIntegerArray{{.ModelName}}) LengthLessThan(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) < '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldIntegerArray{{.ModelName}}) LengthLT(len int) *{{lower .ModelName}}Where {
	return wfba.LengthLessThan(len)
}
func(wfba whereFieldIntegerArray{{.ModelName}}) LengthLessThanOrEqual(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) <= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldIntegerArray{{.ModelName}}) LengthLTE(len int) *{{lower .ModelName}}Where {
	return wfba.LengthLessThanOrEqual(len)
}
func(wfba whereFieldIntegerArray{{.ModelName}}) LengthGreaterThan(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) > '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldIntegerArray{{.ModelName}}) LengthGT(len int) *{{lower .ModelName}}Where {
	return wfba.LengthGreaterThan(len)
}
func(wfba whereFieldIntegerArray{{.ModelName}}) LengthGreaterThanOrEqual(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) >= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldIntegerArray{{.ModelName}}) LengthGTE(len int) *{{lower .ModelName}}Where {
	return wfba.LengthGreaterThanOrEqual(len)
}
`