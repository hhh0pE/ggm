package fields

const StringArrayTemplate = `
type whereFieldStringArray{{.ModelName}} struct {
	name string
	where *{{lower .ModelName}}Where
}

func(wfba whereFieldStringArray{{.ModelName}}) Is(val []string) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" = '"+stringArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldStringArray{{.ModelName}}) IsNot(val []string) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <> '"+stringArrayToSqlValue(val)+"'")
return wfba.where
}
func(wfba whereFieldStringArray{{.ModelName}}) LessThan(val []string) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" < '"+stringArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldStringArray{{.ModelName}}) LT(val []string) *{{lower .ModelName}}Where {
	return wfba.LessThan(val)
}
func(wfba whereFieldStringArray{{.ModelName}}) LessThanOrEqual(val []string) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <= '"+stringArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldStringArray{{.ModelName}}) LTE(val []string) *{{lower .ModelName}}Where {
	return wfba.LessThanOrEqual(val)
}
func(wfba whereFieldStringArray{{.ModelName}}) GreaterThan(val []string) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" > '"+stringArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldStringArray{{.ModelName}}) GT(val []string) *{{lower .ModelName}}Where {
	return wfba.GreaterThan(val)
}
func(wfba whereFieldStringArray{{.ModelName}}) GreaterThanOrEqual(val []string) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" > '"+stringArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldStringArray{{.ModelName}}) GTE(val []string) *{{lower .ModelName}}Where {
	return wfba.GreaterThanOrEqual(val)
}
func(wfba whereFieldStringArray{{.ModelName}}) Contains(val string) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" @> '"+stringArrayToSqlValue([]string{val})+"'")
	return wfba.where
}
func(wfba whereFieldStringArray{{.ModelName}}) ContainedBy(val []string) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <@ '"+stringArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldStringArray{{.ModelName}}) Overlap(val []string) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" && '"+stringArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldStringArray{{.ModelName}}) LengthIs(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("array_length(\"" + wfba.name + "\", 1) = '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldStringArray{{.ModelName}}) LengthLessThan(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("array_length(\"" + wfba.name + "\", 1) < '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldStringArray{{.ModelName}}) LengthLT(len int) *{{lower .ModelName}}Where {
	return wfba.LengthLessThan(len)
}
func(wfba whereFieldStringArray{{.ModelName}}) LengthLessThanOrEqual(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("array_length(\"" + wfba.name + "\", 1) <= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldStringArray{{.ModelName}}) LengthLTE(len int) *{{lower .ModelName}}Where {
	return wfba.LengthLessThanOrEqual(len)
}
func(wfba whereFieldStringArray{{.ModelName}}) LengthGreaterThan(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("array_length(\"" + wfba.name + "\", 1) > '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldStringArray{{.ModelName}}) LengthGT(len int) *{{lower .ModelName}}Where {
	return wfba.LengthGreaterThan(len)
}
func(wfba whereFieldStringArray{{.ModelName}}) LengthGreaterThanOrEqual(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("array_length(\"" + wfba.name + "\", 1) >= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldStringArray{{.ModelName}}) LengthGTE(len int) *{{lower .ModelName}}Where {
	return wfba.LengthGreaterThanOrEqual(len)
}
`
