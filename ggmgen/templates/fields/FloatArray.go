package fields

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
