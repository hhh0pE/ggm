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
type whereFieldFloat{{.ModelName}} struct {
	name  string
	where *{{lower .ModelName}}Where
}

func (wfi whereFieldFloat{{.ModelName}}) Is(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" = '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) Eq(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" = '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) Equal(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" = '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) IsNot(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" <> '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) GreaterThan(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" > '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) GT(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" > '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) GreaterThanOrEqual(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" >= '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) GTE(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" >= '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) LessThan(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" < '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) LT(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" < '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) LessThanOrEqual(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" <= '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) LTE(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" <= '%f'", val))
	return wfi.where
}

func (wfi whereFieldFloat{{.ModelName}}) Between(left, right float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" BETWEEN '%f' AND '%f'", left, right))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) NotBetween(left, right float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" NOT BETWEEN '%f' AND '%f'", left, right))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) In(nums []float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	var cond string
	cond += "\"" + wfi.name + "\" IN ("
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
func (wfi whereFieldFloat{{.ModelName}}) Any(nums ...float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	var cond string
	cond += "\"" + wfi.name + "\" IN ("
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
`

const FloatNullableTemplate = `
type whereFieldFloatNullable{{.ModelName}} struct {
	whereFieldFloat{{.ModelName}}
}

func (wff whereFieldFloat{{.ModelName}}) IsNull() *{{lower .ModelName}}Where {
	wff.where.andOr()
	wff.where.addCond("\"" + wff.name + "\" IS NULL")
	return wff.where
}
func (wff whereFieldFloat{{.ModelName}}) IsNotNull() *{{lower .ModelName}}Where {
	wff.where.andOr()
	wff.where.addCond("\"" + wff.name + "\" IS NOT NULL")
	return wff.where
}
`

const FloatArrayTemplate = `
type whereFieldFloatArray{{.ModelName}} struct {
	name string
	where *{{lower .ModelName}}Where
}

func(wfba whereFieldFloatArray{{.ModelName}}) Is(val []float64) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" = '"+float64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray{{.ModelName}}) IsNot(val []float64) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <> '"+float64ArrayToSqlValue(val)+"'")
return wfba.where
}
func(wfba whereFieldFloatArray{{.ModelName}}) LessThan(val []float64) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" < '"+float64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray{{.ModelName}}) LT(val []float64) *{{lower .ModelName}}Where {
	return wfba.LessThan(val)
}
func(wfba whereFieldFloatArray{{.ModelName}}) LessThanOrEqual(val []float64) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <= '"+float64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray{{.ModelName}}) LTE(val []float64) *{{lower .ModelName}}Where {
	return wfba.LessThanOrEqual(val)
}
func(wfba whereFieldFloatArray{{.ModelName}}) GreaterThan(val []float64) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" > '"+float64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray{{.ModelName}}) GT(val []float64) *{{lower .ModelName}}Where {
	return wfba.GreaterThan(val)
}
func(wfba whereFieldFloatArray{{.ModelName}}) GreaterThanOrEqual(val []float64) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" > '"+float64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray{{.ModelName}}) GTE(val []float64) *{{lower .ModelName}}Where {
	return wfba.GreaterThanOrEqual(val)
}
func(wfba whereFieldFloatArray{{.ModelName}}) Contains(val float64) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" @> '"+float64ArrayToSqlValue([]float64{val})+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray{{.ModelName}}) ContainedBy(val []float64) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <@ '"+float64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray{{.ModelName}}) Overlap(val []float64) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" && '"+float64ArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray{{.ModelName}}) LengthIs(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) = '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray{{.ModelName}}) LengthLessThan(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) < '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray{{.ModelName}}) LengthLT(len int) *{{lower .ModelName}}Where {
	return wfba.LengthLessThan(len)
}
func(wfba whereFieldFloatArray{{.ModelName}}) LengthLessThanOrEqual(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) <= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray{{.ModelName}}) LengthLTE(len int) *{{lower .ModelName}}Where {
	return wfba.LengthLessThanOrEqual(len)
}
func(wfba whereFieldFloatArray{{.ModelName}}) LengthGreaterThan(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) > '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray{{.ModelName}}) LengthGT(len int) *{{lower .ModelName}}Where {
	return wfba.LengthGreaterThan(len)
}
func(wfba whereFieldFloatArray{{.ModelName}}) LengthGreaterThanOrEqual(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) >= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldFloatArray{{.ModelName}}) LengthGTE(len int) *{{lower .ModelName}}Where {
	return wfba.LengthGreaterThanOrEqual(len)
}
`
