package fieldType

type Float struct {
	IsNullable   bool
	IsArray      bool
	IsGoBaseType bool
}

func (f Float) Name() string {
	return name("Float", f)
}

func (f Float) SqlType() string {
	return sqlType("REAL", f)
}

func (f Float) WhereTemplate() string {
	if f.IsNullable {
		return FloatNullableTemplate
	}
	if f.IsArray {
		return FloatArrayTemplate
	}
	return FloatTemplate
}

func (f Float) FmtReplacer() string {
	return "%f"
}

func (f Float) DefaultValue() string {
	return "0.0"
}

func (f Float) GoBaseType() string {
	if f.IsGoBaseType {
		return goType("float64", f)
	}
	return f.GoScannerType()
}

func (f Float) GoScannerType() string {
	if f.IsArray {
		if f.IsNullable {
			return "*pq.NullArray"
		} else {
			return "pq.NullArray"
		}
	}
	if f.IsNullable {
		return "sql.NullFloat64"
	}

	return "float64"
}

func (Float) Size() int {
	return 0
}

func (Float) MaxSize() int {
	return 0
}

func (f Float) Nullable() bool {
	return f.IsNullable
}

func (f Float) Array() bool {
	return f.IsArray
}
func (f Float) ConstType() ConstFieldType {
	return FloatType
}
func (f Float) ImplementScannerInterface() bool {
	return f.GoBaseType() == f.GoScannerType()
}



const FloatTemplate = `
{{if not .}}
type whereFieldFloat struct {
	name  string
	where modelWhere
}
func (wff whereFieldFloat) sqlName() string {
	return "\""+wff.name+"\""
}

func (wfi whereFieldFloat) is(val float64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" = '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat) eq(val float64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" = '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat) equal(val float64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" = '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat) isNot(val float64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" <> '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat) greaterThan(val float64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" > '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat) greaterThanOrEqual(val float64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" >= '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat) lessThan(val float64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" < '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat) lessThanOrEqual(val float64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" <= '%f'", val))
	return wfi.where
}

func (wfi whereFieldFloat) between(left, right float64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" BETWEEN '%f' AND '%f'", left, right))
	return wfi.where
}
func (wfi whereFieldFloat) notBetween(left, right float64) modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf(wfi.sqlName()+" NOT BETWEEN '%f' AND '%f'", left, right))
	return wfi.where
}
func (wfi whereFieldFloat) in(nums []float64) modelWhere {
	wfi.where.andOr()
	var cond string
	cond += wfi.sqlName()+" IN ("
	for i, num := range nums {
		cond += fmt.Sprintf("'%f'", num)
		if i != len(nums) {
			cond += ", "
		}
	}
	cond += ")"
	wfi.where.addCond(cond)
	return wfi.where
}
func (wfi whereFieldFloat) any(nums ...float64) modelWhere {
	wfi.where.andOr()
	var cond string
	cond += wfi.sqlName()+" IN ("
	for i, num := range nums {
		cond += fmt.Sprintf("'%f'", num)
		if i != len(nums) {
			cond += ", "
		}
	}
	cond += ")"
	wfi.where.addCond(cond)
	return wfi.where
}
func (wfi whereFieldFloat) isNull() modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(wfi.sqlName()+" IS NULL")
	return wfi.where
}
func (wfi whereFieldFloat) isNotNull() modelWhere {
	wfi.where.andOr()
	wfi.where.addCond(wfi.sqlName()+" IS NOT NULL")
	return wfi.where
}
{{else}}
type whereFieldFloat{{.ModelName}} struct {
	name  string
	where *{{lower .ModelName}}Where
}

func (wff whereFieldFloat{{.ModelName}}) sqlName() string {
	return "\"{{.ModelTableName}}\".\""+wff.name+"\""
}
func (wfi whereFieldFloat{{.ModelName}}) Is(val float64) *{{lower .ModelName}}Where {
	return wfi.is(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldFloat{{.ModelName}}) Eq(val float64) *{{lower .ModelName}}Where {
	return wfi.eq(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldFloat{{.ModelName}}) Equal(val float64) *{{lower .ModelName}}Where {
	return wfi.equal(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldFloat{{.ModelName}}) IsNot(val float64) *{{lower .ModelName}}Where {
	return wfi.isNot(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldFloat{{.ModelName}}) GreaterThan(val float64) *{{lower .ModelName}}Where {
	return wfi.greaterThan(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldFloat{{.ModelName}}) GT(val float64) *{{lower .ModelName}}Where {
	return wfi.GreaterThan(val)
}
func (wfi whereFieldFloat{{.ModelName}}) GreaterThanOrEqual(val float64) *{{lower .ModelName}}Where {
	return wfi.greaterThanOrEqual(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldFloat{{.ModelName}}) GTE(val float64) *{{lower .ModelName}}Where {
	return wfi.GreaterThanOrEqual(val)
}
func (wfi whereFieldFloat{{.ModelName}}) LessThan(val float64) *{{lower .ModelName}}Where {
	return wfi.lessThan(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldFloat{{.ModelName}}) LT(val float64) *{{lower .ModelName}}Where {
	return wfi.LessThan(val)
}
func (wfi whereFieldFloat{{.ModelName}}) LessThanOrEqual(val float64) *{{lower .ModelName}}Where {
	return wfi.lessThanOrEqual(val).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldFloat{{.ModelName}}) LTE(val float64) *{{lower .ModelName}}Where {
	return wfi.LessThanOrEqual(val)
}

func (wfi whereFieldFloat{{.ModelName}}) Between(left, right float64) *{{lower .ModelName}}Where {
	return wfi.between(left, right).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldFloat{{.ModelName}}) NotBetween(left, right float64) *{{lower .ModelName}}Where {
	return wfi.notBetween(left, right).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldFloat{{.ModelName}}) In(nums []float64) *{{lower .ModelName}}Where {
	return wfi.in(nums).(*{{lower .ModelName}}Where)
}
func (wfi whereFieldFloat{{.ModelName}}) Any(nums ...float64) *{{lower .ModelName}}Where {
	return wfi.any(val).(*{{lower .ModelName}}Where)
}
{{end}}
`

const FloatNullableTemplate = `
{{if .}}
type whereFieldFloatNullable{{.ModelName}} struct {
	whereFieldFloat{{.ModelName}}
}

func (wff whereFieldFloatNullable{{.ModelName}}) IsNull() *{{lower .ModelName}}Where {
	return wff.isNull().(*{{lower .ModelName}}Where)
}
func (wff whereFieldFloatNullable{{.ModelName}}) IsNotNull() *{{lower .ModelName}}Where {
	return wfi.isNotNull().(*{{lower .ModelName}}Where)
}
{{end}}
`

const FloatArrayTemplate = `
{{if not .}}
type whereFieldFloatArray struct {
	name string
	where modelWhere
}

func(wfba whereFieldFloatArray) sqlName() string {
	return "\""+wfba.name+"\""
}

func(wfba whereFieldFloatArray) is(val []float64) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" = '"+float64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray) isNot(val []float64) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" <> '"+float64ArrayToSqlValue(val)+"'")
return wfba.where
}
func(wfba whereFieldFloatArray) lessThan(val []float64) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" < '"+float64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray) lessThanOrEqual(val []float64) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" <= '"+float64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray) greaterThan(val []float64) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" > '"+float64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray) greaterThanOrEqual(val []float64) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" > '"+float64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray) contains(val float64) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" @> '"+float64ArrayToSqlValue([]float64{val})+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray) containedBy(val []float64) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" <@ '"+float64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray) overlap(val []float64) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond(wfba.sqlName()+" && '"+float64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray) lengthIs(len int) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond("array_length("+wfba.sqlName()+", 1) = '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray) lengthLessThan(len int) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond("array_length("+wfba.sqlName()+", 1) < '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray) lengthLessThanOrEqual(len int) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond("array_length("+wfba.sqlName()+", 1) <= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray) lengthGreaterThan(len int) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond("array_length("+wfba.sqlName()+", 1) > '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray) lengthGreaterThanOrEqual(len int) modelWhere {
	wfba.where.andOr()
	wfba.where.addCond("array_length("+wfba.sqlName()+", 1) >= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
{{else}}
type whereFieldFloatArray{{.ModelName}} struct {
	whereFieldFloatArray
}

func(wffa whereFieldFloatArray{{.ModelName}}) sqlName() string {
	return "\"{{.ModelTableName}}\".\""+wffa.name+"\""
}

func(wffa whereFieldFloatArray{{.ModelName}}) Is(val []float64) *{{lower .ModelName}}Where {
	return wffa.is(val).(*{{lower .ModelName}})
}
func(wffa whereFieldFloatArray{{.ModelName}}) IsNot(val []float64) *{{lower .ModelName}}Where {
	return wffa.isNot(val).(*{{lower .ModelName}})
}
func(wffa whereFieldFloatArray{{.ModelName}}) LessThan(val []float64) *{{lower .ModelName}}Where {
	return wffa.lessThan(val).(*{{lower .ModelName}})
}
func(wffa whereFieldFloatArray{{.ModelName}}) LT(val []float64) *{{lower .ModelName}}Where {
	return wffa.LessThan(val)
}
func(wffa whereFieldFloatArray{{.ModelName}}) LessThanOrEqual(val []float64) *{{lower .ModelName}}Where {
	return wffa.lessThanOrEqual(val).(*{{lower .ModelName}})
}
func(wffa whereFieldFloatArray{{.ModelName}}) LTE(val []float64) *{{lower .ModelName}}Where {
	return wffa.LessThanOrEqual(val)
}
func(wffa whereFieldFloatArray{{.ModelName}}) GreaterThan(val []float64) *{{lower .ModelName}}Where {
	return wffa.greaterThan(val).(*{{lower .ModelName}})
}
func(wffa whereFieldFloatArray{{.ModelName}}) GT(val []float64) *{{lower .ModelName}}Where {
	return wffa.GreaterThan(val)
}
func(wffa whereFieldFloatArray{{.ModelName}}) GreaterThanOrEqual(val []float64) *{{lower .ModelName}}Where {
	return wffa.greaterThanOrEqual(val).(*{{lower .ModelName}})
}
func(wffa whereFieldFloatArray{{.ModelName}}) GTE(val []float64) *{{lower .ModelName}}Where {
	return wffa.GreaterThanOrEqual(val)
}
func(wffa whereFieldFloatArray{{.ModelName}}) Contains(val float64) *{{lower .ModelName}}Where {
	return wffa.contains(val).(*{{lower .ModelName}})
}
func(wffa whereFieldFloatArray{{.ModelName}}) ContainedBy(val []float64) *{{lower .ModelName}}Where {
	return wffa.containedBy(val).(*{{lower .ModelName}})
}
func(wffa whereFieldFloatArray{{.ModelName}}) Overlap(val []float64) *{{lower .ModelName}}Where {
	return wffa.overlap(val).(*{{lower .ModelName}})
}
func(wffa whereFieldFloatArray{{.ModelName}}) LengthIs(len int) *{{lower .ModelName}}Where {
	return wffa.lengthIs(len).(*{{lower .ModelName}})
}
func(wffa whereFieldFloatArray{{.ModelName}}) LengthLessThan(len int) *{{lower .ModelName}}Where {
	return wffa.lengthLessThan(len).(*{{lower .ModelName}})
}
func(wffa whereFieldFloatArray{{.ModelName}}) LengthLT(len int) *{{lower .ModelName}}Where {
	return wffa.LengthLessThan(len)
}
func(wffa whereFieldFloatArray{{.ModelName}}) LengthLessThanOrEqual(len int) *{{lower .ModelName}}Where {
	return wffa.lengthLessThanOrEqual(val).(*{{lower .ModelName}})
}
func(wffa whereFieldFloatArray{{.ModelName}}) LengthLTE(len int) *{{lower .ModelName}}Where {
	return wffa.LengthLessThanOrEqual(len)
}
func(wffa whereFieldFloatArray{{.ModelName}}) LengthGreaterThan(len int) *{{lower .ModelName}}Where {
	return wffa.lengthGreaterThan(val).(*{{lower .ModelName}})
}
func(wffa whereFieldFloatArray{{.ModelName}}) LengthGT(len int) *{{lower .ModelName}}Where {
	return wffa.LengthGreaterThan(len)
}
func(wffa whereFieldFloatArray{{.ModelName}}) LengthGreaterThanOrEqual(len int) *{{lower .ModelName}}Where {
	return wffa.lengthGreaterThanOrEqual(val).(*{{lower .ModelName}})
}
func(wffa whereFieldFloatArray{{.ModelName}}) LengthGTE(len int) *{{lower .ModelName}}Where {
	return wffa.LengthGreaterThanOrEqual(len)
}
{{end}}
`
