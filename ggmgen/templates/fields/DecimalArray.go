package fields

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
