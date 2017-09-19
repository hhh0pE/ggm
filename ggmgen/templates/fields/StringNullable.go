package fields

const StringNullableTemplate = `
type whereFieldStringNullable{{.ModelName}} struct {
	whereFieldString{{.ModelName}}
}

func (wfsn *whereFieldStringNullable{{.ModelName}}) IsNull() *{{lower .ModelName}}Where {
	wfsn.where.andOr()
	wfsn.where.addCond("\"" + wfsn.name + "\" IS NULL")
	return wfsn.where
}
func (wfsn *whereFieldStringNullable{{.ModelName}}) IsNotNull() *{{lower .ModelName}}Where {
	wfsn.where.andOr()
	wfsn.where.addCond("\"" + wfsn.name + "\" IS NOT NULL")
	return wfsn.where
}
`
