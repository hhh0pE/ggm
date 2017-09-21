package fields

const FloatNullableTemplate = `
type whereFieldFloatNullable{{.ModelName}} struct {
	name  string
	where modelWhere
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
