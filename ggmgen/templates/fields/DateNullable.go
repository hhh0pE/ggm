package fields

const DateNullableTemplate = `
type whereFieldDateNullable{{.ModelName}} struct {
	whereFieldDate{{.ModelName}}
}

func (wfd whereFieldDateNullable{{.ModelName}}) IsNull() *{{lower .ModelName}}Where {
	wfd.where.andOr()
	wfd.where.addCond("\"" + wfd.name + "\" IS NULL")
	return wfd.where
}
func (wfd whereFieldDateNullable{{.ModelName}}) IsNotNull() *{{lower .ModelName}}Where {
	wfd.where.andOr()
	wfd.where.addCond("\"" + wfd.name + "\" IS NOT NULL")
	return wfd.where
}
`
