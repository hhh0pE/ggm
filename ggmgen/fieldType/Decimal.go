package fieldType

type Decimal struct {
	IsNullable   bool
	IsArray      bool
	IsGoBaseType bool
}

func (d Decimal) Name() string {
	return name("Decimal", d)
}

func (d Decimal) SqlType() string {
	return sqlType("NUMERIC", d)
}

func (d Decimal) WhereTemplate() string {
	if d.IsArray {
		return DecimalArrayTemplate
	}
	if d.IsNullable {
		return DecimalNullableTemplate
	}
	return DecimalTemplate
}

func (d Decimal) FmtReplacer() string {
	return "%s"
}

func (d Decimal) DefaultValue() string {
	return `"0"`
}

func (d Decimal) GoBaseType() string {
	if d.IsGoBaseType {
		return goType("Decimal", d)
	}
	return d.GoScannerType()
}

func (d Decimal) GoScannerType() string {
	if d.IsArray {
		if d.IsNullable {
			return "*ggm.DecimalArray"
		} else {
			return "ggm.DecimalArray"
		}
	}
	if d.IsNullable {
		return "ggm.NullDecimal"
	}
	return "Decimal"
}

func (d Decimal) MaxSize() int {
	return 0
}

func (d Decimal) Nullable() bool {
	return d.IsNullable
}

func (d Decimal) Array() bool {
	return d.IsArray
}
func (d Decimal) ConstType() ConstFieldType {
	return DecimalType
}
func (d Decimal) ImplementScannerInterface() bool {
	return d.GoBaseType() == d.GoScannerType()
}

const DecimalTemplate = `
{{if not .}}
type whereFieldDecimal struct {
	name  string
	where modelWhere
}
func(wfi whereFieldDecimal) sqlName() string {
	return wfi.name
}
func (wfi whereFieldDecimal) is(val ggm.Decimal) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(wfi.sqlName()+" = '"+val.String()+"'")
	return wfi.where.modelWhere()
}
func (wfi whereFieldDecimal) eq(val ggm.Decimal) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(wfi.sqlName()+" = '"+val.String()+"'")
	return wfi.where.modelWhere()
}
func (wfi whereFieldDecimal) equal(val ggm.Decimal) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(wfi.sqlName()+" = '"+val.String()+"'")
	return wfi.where.modelWhere()
}
func (wfi whereFieldDecimal) isNot(val ggm.Decimal) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(wfi.sqlName()+" <> '"+val.String()+"'")
	return wfi.where.modelWhere()
}
func (wfi whereFieldDecimal) greaterThan(val ggm.Decimal) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(wfi.sqlName()+" > '"+val.String()+"'")
	return wfi.where.modelWhere()
}
func (wfi whereFieldDecimal) greaterThanOrEqual(val ggm.Decimal) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(wfi.sqlName()+" >= '"+val.String()+"'")
	return wfi.where.modelWhere()
}
func (wfi whereFieldDecimal) lessThan(val ggm.Decimal) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(wfi.sqlName()+" < '"+val.String()+"'")
	return wfi.where.modelWhere()
}
func (wfi whereFieldDecimal) lessThanOrEqual(val ggm.Decimal) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(wfi.sqlName()+" <= '"+val.String()+"'")
	return wfi.where.modelWhere()
}

func (wfi whereFieldDecimal) between(left, right ggm.Decimal) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(wfi.sqlName()+" BETWEEN '"+left.String()+"' AND '"+right.String()+"'")
	return wfi.where.modelWhere()
}
func (wfi whereFieldDecimal) notBetween(left, right ggm.Decimal) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(wfi.sqlName()+" NOT BETWEEN '"+left.String()+"' AND '"+right.String()+"'")
	return wfi.where.modelWhere()
}
func (wfi whereFieldDecimal) in(nums []ggm.Decimal) modelWhere {
	wfi.where.andOr()
	var cond string
	cond += wfi.sqlName() + " IN ("
	for i, num := range nums {
		cond += "'"+num.String()+"'"
		if i != len(nums) {
			cond += ", "
		}
	}
	cond += ")"
	wfi.where.addCond(cond)
	return wfi.where.modelWhere()
}
func (wfi whereFieldDecimal) isNull() modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(wfi.sqlName()+" IS NULL")
	return wfi.where.modelWhere()
}
func (wfi whereFieldDecimal) isNotNull() modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(wfi.sqlName() + " IS NOT NULL")
	return wfi.where.modelWhere()
}
{{else}}
type whereFieldDecimal{{.ModelName}} struct {
	whereFieldDecimal
}
func (wfi whereFieldDecimal{{.ModelName}}) Is(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.is(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldDecimal{{.ModelName}}) Eq(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.eq(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldDecimal{{.ModelName}}) Equal(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.Eq(val)
}
func (wfi whereFieldDecimal{{.ModelName}}) IsNot(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.isNot(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldDecimal{{.ModelName}}) GreaterThan(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.greaterThan(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldDecimal{{.ModelName}}) GT(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.GreaterThan(val)
}
func (wfi whereFieldDecimal{{.ModelName}}) GreaterThanOrEqual(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.greaterThanOrEqual(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldDecimal{{.ModelName}}) GTE(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.GreaterThanOrEqual(val)
}
func (wfi whereFieldDecimal{{.ModelName}}) LessThan(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.lessThan(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldDecimal{{.ModelName}}) LT(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.LessThan(val)
}
func (wfi whereFieldDecimal{{.ModelName}}) LessThanOrEqual(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.lessThanOrEqual(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldDecimal{{.ModelName}}) LTE(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.LessThanOrEqual(val)
}

func (wfi whereFieldDecimal{{.ModelName}}) Between(left, right ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.between(left, right).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldDecimal{{.ModelName}}) NotBetween(left, right ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.notBetween(left, right).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldDecimal{{.ModelName}}) In(nums []ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.in(nums).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldDecimal{{.ModelName}}) Any(nums ...ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.In(nums)
}
{{end}}
`

const DecimalNullableTemplate = `
{{if .}}
type whereFieldDecimalNullable{{.ModelName}} struct {
	whereFieldDecimal{{.ModelName}}
}
func (wff whereFieldDecimal{{.ModelName}}) IsNull() *{{lower .ModelName}}Where {
	return wff.isNull().(*{{lower .ModelName}}Where)
}
func (wff whereFieldDecimal{{.ModelName}}) IsNotNull() *{{lower .ModelName}}Where {
	return wff.isNotNull().(*{{lower .ModelName}}Where)
}
{{end}}
`

const DecimalArrayTemplate = `
{{if not .}}
type whereFieldDecimalArray struct {
	name string
	where modelWhere
}
func(wfba whereFieldDecimalArray) sqlName() string {
	return wfba.name
}
func(wfba whereFieldDecimalArray) is(val []ggm.Decimal) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" = '"+decimalArrayToSqlValue(val)+"'")
	return wfba.where.modelWhere()
}
func(wfba whereFieldDecimalArray) isNot(val []ggm.Decimal) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" <> '"+decimalArrayToSqlValue(val)+"'")
	return wfba.where.modelWhere()
}
func(wfba whereFieldDecimalArray) lessThan(val []ggm.Decimal) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" < '"+decimalArrayToSqlValue(val)+"'")
	return wfba.where.modelWhere()
}
func(wfba whereFieldDecimalArray) lessThanOrEqual(val []ggm.Decimal) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" <= '"+decimalArrayToSqlValue(val)+"'")
	return wfba.where.modelWhere()
}
func(wfba whereFieldDecimalArray) greaterThan(val []ggm.Decimal) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" > '"+decimalArrayToSqlValue(val)+"'")
	return wfba.where.modelWhere()
}
func(wfba whereFieldDecimalArray) greaterThanOrEqual(val []ggm.Decimal) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" > '"+decimalArrayToSqlValue(val)+"'")
	return wfba.where.modelWhere()
}
func(wfba whereFieldDecimalArray) contains(val ggm.Decimal) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" @> '"+decimalArrayToSqlValue([]ggm.Decimal{val})+"'")
	return wfba.where.modelWhere()
}
func(wfba whereFieldDecimalArray) containedBy(val []ggm.Decimal) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" <@ '"+decimalArrayToSqlValue(val)+"'")
	return wfba.where.modelWhere()
}
func(wfba whereFieldDecimalArray) overlap(val []ggm.Decimal) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" && '"+decimalArrayToSqlValue(val)+"'")
	return wfba.where.modelWhere()
}
func(wfba whereFieldDecimalArray) lengthIs(len int) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond("array_length("+wfba.sqlName()+", 1) = '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where.modelWhere()
}
func(wfba whereFieldDecimalArray) lengthLessThan(len int) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond("array_length("+wfba.sqlName()+", 1) < '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where.modelWhere()
}
func(wfba whereFieldDecimalArray) lengthLessThanOrEqual(len int) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond("array_length("+wfba.sqlName()+", 1) <= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where.modelWhere()
}
func(wfba whereFieldDecimalArray) lengthGreaterThan(len int) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond("array_length("+wfba.sqlName()+", 1) > '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where.modelWhere()
}
func(wfba whereFieldDecimalArray) lengthGreaterThanOrEqual(len int) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond("array_length("+wfba.sqlName()+", 1) >= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where.modelWhere()
}
{{else}}
type whereFieldDecimalArray{{.ModelName}} struct {
	whereFieldDecimalArray
}
func(wfba whereFieldDecimalArray{{.ModelName}}) Is(val []ggm.Decimal) *{{lower .ModelName}}Where {
	return wfba.is(val).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) IsNot(val []ggm.Decimal) *{{lower .ModelName}}Where {
	return wfba.isNot(val).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LessThan(val []ggm.Decimal) *{{lower .ModelName}}Where {
	return wfba.lessThan(val).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LT(val []ggm.Decimal) *{{lower .ModelName}}Where {
	return wfba.LessThan(val)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LessThanOrEqual(val []ggm.Decimal) *{{lower .ModelName}}Where {
	return wfba.lessThanOrEqual(val).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LTE(val []ggm.Decimal) *{{lower .ModelName}}Where {
	return wfba.LessThanOrEqual(val)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) GreaterThan(val []ggm.Decimal) *{{lower .ModelName}}Where {
	return wfba.greaterThan(val).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) GT(val []ggm.Decimal) *{{lower .ModelName}}Where {
	return wfba.GreaterThan(val)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) GreaterThanOrEqual(val []ggm.Decimal) *{{lower .ModelName}}Where {
	return wfba.greaterThanOrEqual(val).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) GTE(val []ggm.Decimal) *{{lower .ModelName}}Where {
	return wfba.GreaterThanOrEqual(val)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) Contains(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfba.contains(val).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) ContainedBy(val []ggm.Decimal) *{{lower .ModelName}}Where {
	return wfba.containedBy(val).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) Overlap(val []ggm.Decimal) *{{lower .ModelName}}Where {
	return wfba.overlap(val).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LengthIs(len int) *{{lower .ModelName}}Where {
	return wfba.lengthIs(len).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LengthLessThan(len int) *{{lower .ModelName}}Where {
	return wfba.lengthLessThan(len).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LengthLT(len int) *{{lower .ModelName}}Where {
	return wfba.LengthLessThan(len)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LengthLessThanOrEqual(len int) *{{lower .ModelName}}Where {
	return wfba.lengthLessThanOrEqual(len).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LengthLTE(len int) *{{lower .ModelName}}Where {
	return wfba.LengthLessThanOrEqual(len)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LengthGreaterThan(len int) *{{lower .ModelName}}Where {
	return wfba.lengthGreaterThanOrEqual(len).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LengthGT(len int) *{{lower .ModelName}}Where {
	return wfba.LengthGreaterThan(len)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LengthGreaterThanOrEqual(len int) *{{lower .ModelName}}Where {
	return wfba.lengthGreaterThanOrEqual(len).(*{{lower .ModelName}}Where)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LengthGTE(len int) *{{lower .ModelName}}Where {
	return wfba.LengthGreaterThanOrEqual(len)
}
{{end}}
`
