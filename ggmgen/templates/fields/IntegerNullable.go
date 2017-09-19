package fields

const IntegerNullableTemplate = `
type whereFieldIntegerNullable{{.ModelName}} struct {
	whereFieldInteger{{.ModelName}}
}

func (wfin whereFieldIntegerNullable{{.ModelName}}) IsNull() *{{lower .ModelName}}Where {
	wfin.where.andOr()
	wfin.where.addCond("\"" + wfin.name + "\" IS NULL")
	return wfin.where
}
func (wfin whereFieldIntegerNullable{{.ModelName}}) IsNotNull() *{{lower .ModelName}}Where {
	wfin.where.andOr()
	wfin.where.addCond("\"" + wfin.name + "\" IS NOT NULL")
	return wfin.where
}
`
