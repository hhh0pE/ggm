package fields

const BooleanArrayTemplate = `
type whereFieldBooleanArray{{.ModelName}} struct {
	name string
	where *{{lower .ModelName}}Where
}

func(wfba whereFieldBooleanArray{{.ModelName}}) Is(val []bool) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" = '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) IsNot(val []bool) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <> '"+boolArrayToSqlValue(val)+"'")
return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LessThan(val []bool) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" < '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LT(val []bool) *{{lower .ModelName}}Where {
	return wfba.LessThan(val)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LessThanOrEqual(val []bool) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <= '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LTE(val []bool) *{{lower .ModelName}}Where {
	return wfba.LessThanOrEqual(val)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) GreaterThan(val []bool) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" > '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) GT(val []bool) *{{lower .ModelName}}Where {
	return wfba.GreaterThan(val)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) GreaterThanOrEqual(val []bool) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" > '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) GTE(val []bool) *{{lower .ModelName}}Where {
	return wfba.GreaterThanOrEqual(val)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) Contains(val bool) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" @> '"+boolArrayToSqlValue([]bool{val})+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) ContainedBy(val []bool) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <@ '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) Overlap(val []bool) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" && '"+boolArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthIs(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) = '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthLessThan(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) < '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthLT(len int) *{{lower .ModelName}}Where {
	return wfba.LengthLessThan(len)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthLessThanOrEqual(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) <= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthLTE(len int) *{{lower .ModelName}}Where {
	return wfba.LengthLessThanOrEqual(len)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthGreaterThan(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) > '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthGT(len int) *{{lower .ModelName}}Where {
	return wfba.LengthGreaterThan(len)
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthGreaterThanOrEqual(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfs.where.addCond("array_length(\"" + wfs.name + "\", 1) >= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldBooleanArray{{.ModelName}}) LengthGTE(len int) *{{lower .ModelName}}Where {
	return wfba.LengthGreaterThanOrEqual(len)
}
`
