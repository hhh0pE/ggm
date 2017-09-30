package fields

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
