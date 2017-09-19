package fields

const BooleanNullableTemplate = `
type whereFieldBooleanNullable{{.ModelName}} struct {
	whereFieldBoolean{{.ModelName}}
}

func (wfbn whereFieldBooleanNullable{{.ModelName}}) IsNull() *{{lower .ModelName}}Where {
	wfbn.where.andOr()
	wfbn.where.addCond("\"" + wfbn.name + "\" IS NULL")
	return wfbn.where
}
func (wfbn whereFieldBooleanNullable{{.ModelName}}) IsNotNull() *{{lower .ModelName}}Where {
	wfbn.where.andOr()
	wfbn.where.addCond("\"" + wfbn.name + "\" IS NOT NULL")
	return wfbn.where
}
`
