package fieldType

type Integer struct {
	IsNullable   bool
	IsArray      bool
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
{{if not .}}
type whereFieldInteger struct {
	name  string
	where modelWhere
}
func(wfi whereFieldInteger) sqlName() string {
	return wfi.name
}
func (wfi whereFieldInteger) is(val int64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" = '%d'", val))
	return wfi.where.modelWhere()
}
func (wfi whereFieldInteger) eq(val int64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" = '%d'", val))
	return wfi.where.modelWhere()
}
func (wfi whereFieldInteger) equal(val int64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" = '%d'", val))
	return wfi.where.modelWhere()
}
func (wfi whereFieldInteger) isNot(val int64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" <> '%d'", val))
	return wfi.where.modelWhere()
}
func (wfi whereFieldInteger) greaterThan(val int64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" > '%d'", val))
	return wfi.where.modelWhere()
}
func (wfi whereFieldInteger) greaterThanOrEqual(val int64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" >= '%d'", val))
	return wfi.where.modelWhere()
}
func (wfi whereFieldInteger) lessThan(val int64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" < '%d'", val))
	return wfi.where.modelWhere()
}
func (wfi whereFieldInteger) lessThanOrEqual(val int64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" <= '%d'", val))
	return wfi.where.modelWhere()
}

func (wfi whereFieldInteger) between(left, right int64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" BETWEEN '%d' AND '%d'", left, right))
	return wfi.where.modelWhere()
}
func (wfi whereFieldInteger) notBetween(left, right int64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" NOT BETWEEN '%d' AND '%d'", left, right))
	return wfi.where.modelWhere()
}
func (wfi whereFieldInteger) in(nums []int64) modelWhere {
	wfi.where.andOr()
	var cond string
	cond += wfi.sqlName() + " IN ("
	for i, num := range nums {
		cond += fmt.Sprintf("'%d'", num)
		if i != len(nums) {
			cond += ", "
		}
	}
	cond += ")"
	wfi.where.addCond(cond)
	return wfi.where.modelWhere()
}
func (wfi whereFieldInteger) any(nums ...int64) modelWhere {
	wfi.where.andOr()
	var cond string
	cond += wfi.sqlName() + " IN ("
	for i, num := range nums {
		cond += fmt.Sprintf("'%d'", num)
		if i != len(nums) {
			cond += ", "
		}
	}
	cond += ")"
	wfi.where.addCond(cond)
	return wfi.where.modelWhere()
}
func (wfi whereFieldInteger) isNull() modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(wfi.sqlName() + " IS NULL")
	return wfi.where.modelWhere()
}
func (wfi whereFieldInteger) isNotNull() modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(wfi.sqlName() + " IS NOT NULL")
	return wfi.where.modelWhere()
}
{{else}}
type whereFieldInteger{{.ModelName}} struct {
	whereFieldInteger
}
func (wfi whereFieldInteger{{.ModelName}}) Is(val int64) *{{lower .ModelName}}Where {
	return wfi.is(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldInteger{{.ModelName}}) Eq(val int64) *{{lower .ModelName}}Where {
	return wfi.eq(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldInteger{{.ModelName}}) Equal(val int64) *{{lower .ModelName}}Where {
	return wfi.equal(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldInteger{{.ModelName}}) IsNot(val int64) *{{lower .ModelName}}Where {
	return wfi.isNot(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldInteger{{.ModelName}}) GreaterThan(val int64) *{{lower .ModelName}}Where {
	return wfi.greaterThan(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldInteger{{.ModelName}}) GT(val int64) *{{lower .ModelName}}Where {
	return wfi.GreaterThan(val)
}
func (wfi whereFieldInteger{{.ModelName}}) GreaterThanOrEqual(val int64) *{{lower .ModelName}}Where {
	return wfi.greaterThanOrEqual(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldInteger{{.ModelName}}) GTE(val int64) *{{lower .ModelName}}Where {
	return wfi.GreaterThanOrEqual(val)
}
func (wfi whereFieldInteger{{.ModelName}}) LessThan(val int64) *{{lower .ModelName}}Where {
	return wfi.lessThan(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldInteger{{.ModelName}}) LT(val int64) *{{lower .ModelName}}Where {
	return wfi.LessThan(val)
}
func (wfi whereFieldInteger{{.ModelName}}) LessThanOrEqual(val int64) *{{lower .ModelName}}Where {
	return wfi.lessThanOrEqual(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldInteger{{.ModelName}}) LTE(val int64) *{{lower .ModelName}}Where {
	return wfi.LessThanOrEqual(val)
}

func (wfi whereFieldInteger{{.ModelName}}) Between(left, right int64) *{{lower .ModelName}}Where {
	return wfi.between(left, right).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldInteger{{.ModelName}}) NotBetween(left, right int64) *{{lower .ModelName}}Where {
	return wfi.notBetween(left, right).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldInteger{{.ModelName}}) In(nums []int64) *{{lower .ModelName}}Where {
	return wfi.in(nums).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldInteger{{.ModelName}}) Any(nums ...int64) *{{lower .ModelName}}Where {
	return wfi.any(nums...).(*{{lower .ModelName}}Where)
}
{{end}}
`

const IntegerNullableTemplate = `
{{if .}}
type whereFieldIntegerNullable{{.ModelName}} struct {
	whereFieldInteger{{.ModelName}}
}

func (wfin whereFieldIntegerNullable{{.ModelName}}) IsNull() *{{lower .ModelName}}Where {
	return wfin.isNull().(*{{lower .ModelName}}Where)
}
func (wfin whereFieldIntegerNullable{{.ModelName}}) IsNotNull() *{{lower .ModelName}}Where {
	return wfin.isNotNull().(*{{lower .ModelName}}Where)
}
{{end}}
`

const IntegerArrayTemplate = `
{{if not .}}
type whereFieldIntegerArray struct {
	name string
	where modelWhere
}
func(wfia whereFieldIntegerArray) sqlName() string {
	return wfia.name
}
func(wfia whereFieldIntegerArray) is(val []int64) modelWhere {
	wfia.where.andOr()
	wfia.where.addCond(wfia.sqlName() + " = '"+int64ArrayToSqlValue(val)+"'")
	return wfia.where.modelWhere()
}
func(wfia whereFieldIntegerArray) isNot(val []int64) modelWhere {
	wfia.where.andOr()
	wfia.where.addCond(wfia.sqlName() + " <> '"+int64ArrayToSqlValue(val)+"'")
return wfia.where.modelWhere()
}
func(wfia whereFieldIntegerArray) lessThan(val []int64) modelWhere {
	wfia.where.andOr()
	wfia.where.addCond(wfia.sqlName() + " < '"+int64ArrayToSqlValue(val)+"'")
	return wfia.where.modelWhere()
}
func(wfia whereFieldIntegerArray) lessThanOrEqual(val []int64) modelWhere {
	wfia.where.andOr()
	wfia.where.addCond(wfia.sqlName() + " <= '"+int64ArrayToSqlValue(val)+"'")
	return wfia.where.modelWhere()
}
func(wfia whereFieldIntegerArray) greaterThan(val []int64) modelWhere {
	wfia.where.andOr()
	wfia.where.addCond(wfia.sqlName() + " > '"+int64ArrayToSqlValue(val)+"'")
	return wfia.where.modelWhere()
}
func(wfia whereFieldIntegerArray) greaterThanOrEqual(val []int64) modelWhere {
	wfia.where.andOr()
	wfia.where.addCond(wfia.sqlName() + " > '"+int64ArrayToSqlValue(val)+"'")
	return wfia.where.modelWhere()
}
func(wfia whereFieldIntegerArray) contains(val int64) modelWhere {
	wfia.where.andOr()
	wfia.where.addCond(wfia.sqlName() + " @> '"+int64ArrayToSqlValue([]int64{val})+"'")
	return wfia.where.modelWhere()
}
func(wfia whereFieldIntegerArray) containedBy(val []int64) modelWhere {
	wfia.where.andOr()
	wfia.where.addCond(wfia.sqlName() + " <@ '"+int64ArrayToSqlValue(val)+"'")
	return wfia.where.modelWhere()
}
func(wfia whereFieldIntegerArray) overlap(val []int64) modelWhere {
	wfia.where.andOr()
	wfia.where.addCond(wfia.sqlName() + " && '"+int64ArrayToSqlValue(val)+"'")
	return wfia.where.modelWhere()
}
func(wfia whereFieldIntegerArray) lengthIs(len int64) modelWhere {
	wfia.where.andOr()
	wfia.where.addCond("array_length("+wfia.sqlName() + ", 1) = '"+fmt.Sprintf("%d", len)+"'")
	return wfia.where.modelWhere()
}
func(wfia whereFieldIntegerArray) lengthLessThan(len int64) modelWhere {
	wfia.where.andOr()
	wfia.where.addCond("array_length("+wfia.sqlName() + ", 1) < '"+fmt.Sprintf("%d", len)+"'")
	return wfia.where.modelWhere()
}
func(wfia whereFieldIntegerArray) lengthLessThanOrEqual(len int64) modelWhere {
	wfia.where.andOr()
	wfia.where.addCond("array_length("+wfia.sqlName() + ", 1) <= '"+fmt.Sprintf("%d", len)+"'")
	return wfia.where.modelWhere()
}
func(wfia whereFieldIntegerArray) lengthGreaterThan(len int64) modelWhere {
	wfia.where.andOr()
	wfia.where.addCond("array_length("+wfia.sqlName() + ", 1) > '"+fmt.Sprintf("%d", len)+"'")
	return wfia.where.modelWhere()
}
func(wfia whereFieldIntegerArray) lengthGreaterThanOrEqual(len int64) modelWhere {
	wfia.where.andOr()
	wfia.where.addCond("array_length("+wfia.sqlName() + ", 1) >= '"+fmt.Sprintf("%d", len)+"'")
	return wfia.where.modelWhere()
}
{{else}}
type whereFieldIntegerArray{{.ModelName}} struct {
	whereFieldIntegerArray
}
func(wfia whereFieldIntegerArray{{.ModelName}}) Is(val []int64) *{{lower .ModelName}}Where {
	return wfia.is(val).(*{{lower .ModelName}}Where)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) IsNot(val []int64) *{{lower .ModelName}}Where {
	return wfia.isNot(val).(*{{lower .ModelName}}Where)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) LessThan(val []int64) *{{lower .ModelName}}Where {
	return wfia.lessThan(val).(*{{lower .ModelName}}Where)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) LT(val []int64) *{{lower .ModelName}}Where {
	return wfia.LessThan(val)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) LessThanOrEqual(val []int64) *{{lower .ModelName}}Where {
	return wfia.lessThanOrEqual(val).(*{{lower .ModelName}}Where)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) LTE(val []int64) *{{lower .ModelName}}Where {
	return wfia.LessThanOrEqual(val)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) GreaterThan(val []int64) *{{lower .ModelName}}Where {
	return wfia.greaterThan(val).(*{{lower .ModelName}}Where)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) GT(val []int64) *{{lower .ModelName}}Where {
	return wfia.GreaterThan(val)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) GreaterThanOrEqual(val []int64) *{{lower .ModelName}}Where {
	return wfia.greaterThanOrEqual(val).(*{{lower .ModelName}}Where)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) GTE(val []int64) *{{lower .ModelName}}Where {
	return wfia.GreaterThanOrEqual(val)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) Contains(val int64) *{{lower .ModelName}}Where {
	return wfia.contains(val).(*{{lower .ModelName}}Where)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) ContainedBy(val []int64) *{{lower .ModelName}}Where {
	return wfia.containedBy(val).(*{{lower .ModelName}}Where)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) Overlap(val []int64) *{{lower .ModelName}}Where {
	return wfia.overlap(val).(*{{lower .ModelName}}Where)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) LengthIs(len int64) *{{lower .ModelName}}Where {
	return wfia.lengthIs(len).(*{{lower .ModelName}}Where)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) LengthLessThan(len int64) *{{lower .ModelName}}Where {
	return wfia.lengthLessThan(len).(*{{lower .ModelName}}Where)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) LengthLT(len int64) *{{lower .ModelName}}Where {
	return wfia.LengthLessThan(len)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) LengthLessThanOrEqual(len int64) *{{lower .ModelName}}Where {
	return wfia.lengthLessThanOrEqual(len).(*{{lower .ModelName}}Where)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) LengthLTE(len int64) *{{lower .ModelName}}Where {
	return wfia.LengthLessThanOrEqual(len)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) LengthGreaterThan(len int64) *{{lower .ModelName}}Where {
	return wfia.lengthGreaterThan(len).(*{{lower .ModelName}}Where)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) LengthGT(len int64) *{{lower .ModelName}}Where {
	return wfia.LengthGreaterThan(len)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) LengthGreaterThanOrEqual(len int64) *{{lower .ModelName}}Where {
	return wfia.lengthGreaterThanOrEqual(len).(*{{lower .ModelName}}Where)
}
func(wfia whereFieldIntegerArray{{.ModelName}}) LengthGTE(len int64) *{{lower .ModelName}}Where {
	return wfia.LengthGreaterThanOrEqual(len)
}
{{end}}
`
