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
	return "ggm.NewDecimal(0)"
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
type whereFieldDecimal{{.ModelName}} struct {
	name  string
	where *{{lower .ModelName}}Where
}

func (wfi whereFieldDecimal{{.ModelName}}) Is(val ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" = '"+val.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) Eq(val ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" = '"+val.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) Equal(val ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" = '"+val.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) IsNot(val ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" <> '"+val.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) GreaterThan(val ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" > '"+val.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) GT(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.GreaterThan(val)
}
func (wfi whereFieldDecimal{{.ModelName}}) GreaterThanOrEqual(val ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" >= '"+val.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) GTE(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.GreaterThanOrEqual(val)
}
func (wfi whereFieldDecimal{{.ModelName}}) LessThan(val ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" < '"+val.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) LT(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.LessThan(val)
}
func (wfi whereFieldDecimal{{.ModelName}}) LessThanOrEqual(val ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" <= '"+val.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) LTE(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.LessThanOrEqual(val)
}

func (wfi whereFieldDecimal{{.ModelName}}) Between(left, right ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" BETWEEN '"+left.String()+"' AND '"+right.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) NotBetween(left, right ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" NOT BETWEEN '"+left.String()+"' AND '"+right.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) In(nums []ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	var cond string
	cond += "\"" + wfi.name + "\" IN ("
	for i, num := range nums {
		cond += "'"+num.String()+"'"
		if i != len(nums) {
			cond += ", "
		}
	}
	cond += ")"
	wfi.where.addCond(cond)
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) Any(nums ...ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.In(nums)
}
`


const DecimalArrayTemplate = `
type whereFieldDecimalArray{{.ModelName}} struct {
	name string
	where *{{lower .ModelName}}Where
}

func(wfba whereFieldDecimalArray{{.ModelName}}) Is(val []ggm.Decimal) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" = '"+decimalArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldDecimalArray{{.ModelName}}) IsNot(val []ggm.Decimal) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <> '"+decimalArrayToSqlValue(val)+"'")
return wfba.where
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LessThan(val []ggm.Decimal) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" < '"+decimalArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LT(val []ggm.Decimal) *{{lower .ModelName}}Where {
	return wfba.LessThan(val)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LessThanOrEqual(val []ggm.Decimal) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <= '"+decimalArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LTE(val []ggm.Decimal) *{{lower .ModelName}}Where {
	return wfba.LessThanOrEqual(val)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) GreaterThan(val []ggm.Decimal) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" > '"+decimalArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldDecimalArray{{.ModelName}}) GT(val []ggm.Decimal) *{{lower .ModelName}}Where {
	return wfba.GreaterThan(val)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) GreaterThanOrEqual(val []ggm.Decimal) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" > '"+decimalArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldDecimalArray{{.ModelName}}) GTE(val []ggm.Decimal) *{{lower .ModelName}}Where {
	return wfba.GreaterThanOrEqual(val)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) Contains(val decimal) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" @> '"+decimalArrayToSqlValue([]ggm.Decimal{val})+"'")
	return wfba.where
}
func(wfba whereFieldDecimalArray{{.ModelName}}) ContainedBy(val []ggm.Decimal) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <@ '"+decimalArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldDecimalArray{{.ModelName}}) Overlap(val []ggm.Decimal) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" && '"+decimalArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LengthIs(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) = '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LengthLessThan(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) < '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LengthLT(len int) *{{lower .ModelName}}Where {
	return wfba.LengthLessThan(len)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LengthLessThanOrEqual(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) <= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LengthLTE(len int) *{{lower .ModelName}}Where {
	return wfba.LengthLessThanOrEqual(len)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LengthGreaterThan(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) > '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LengthGT(len int) *{{lower .ModelName}}Where {
	return wfba.LengthGreaterThan(len)
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LengthGreaterThanOrEqual(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) >= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldDecimalArray{{.ModelName}}) LengthGTE(len int) *{{lower .ModelName}}Where {
	return wfba.LengthGreaterThanOrEqual(len)
}
`



const DecimalNullableTemplate = `
type whereFieldDecimalNullable{{.ModelName}} struct {
	whereFieldDecimal{{.ModelName}}
}

func (wff whereFieldDecimal{{.ModelName}}) IsNull() *{{lower .ModelName}}Where {
	wff.where.andOr()
	wff.where.addCond("\"" + wff.name + "\" IS NULL")
	return wff.where
}
func (wff whereFieldDecimal{{.ModelName}}) IsNotNull() *{{lower .ModelName}}Where {
	wff.where.andOr()
	wff.where.addCond("\"" + wff.name + "\" IS NOT NULL")
	return wff.where
}
`