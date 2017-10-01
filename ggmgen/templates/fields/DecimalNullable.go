package fields

const DecimalNullableTemplate = `
type whereFieldDecimalNullable{{.ModelName}} struct {
	whereFieldDecimal{{.ModelName}}
}

func (wff whereFieldDecimal{{.ModelName}}) IsNull() *{{lower .ModelName}}Where {
	wff.where.andOr()
	wff.where.addCond("\"" + wff.name + "\" IS NULL")
	return wff.where
}
func (wff whereFieldDecimal{{.ModelName}}) IsNotNull() *{{lower .ModelName}}Where {
	wff.where.andOr()
	wff.where.addCond("\"" + wff.name + "\" IS NOT NULL")
	return wff.where
}
`
