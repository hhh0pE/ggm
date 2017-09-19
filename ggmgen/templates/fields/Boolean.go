package fields

const BooleanTemplate = `
type whereFieldBoolean{{.ModelName}} struct {
    name string
	where *{{lower .ModelName}}Where
}

func (wfb whereFieldBoolean{{.ModelName}}) Is(val bool) *{{lower .ModelName}}Where {
	wfb.where.andOr()
	if val {
		wfb.where.addCond("\"" + wfb.name + "\" = 'TRUE'")
	} else {
		wfb.where.addCond("\"" + wfb.name + "\" = 'FALSE'")
	}
	return wfb.where
}
func (wfb whereFieldBoolean{{.ModelName}}) IsTrue() *{{lower .ModelName}}Where {
	wfb.where.andOr()
	wfb.where.addCond("\"" + wfb.name + "\" = 'TRUE'")
	return wfb.where
}
func (wfb whereFieldBoolean{{.ModelName}}) IsFalse() *{{lower .ModelName}}Where {
	wfb.where.andOr()
	wfb.where.addCond("\"" + wfb.name + "\" = 'FALSE'")
	return wfb.where
}
func (wfb whereFieldBoolean{{.ModelName}}) FromStr(str string) *{{lower .ModelName}}Where {
	var val bool
	str = strings.TrimSpace(strings.ToLower(str))
	if str == "1" || str == "true" || str == "t" || str == "y" || str == "on" || str == "yes" {
		val = true
	}
	return wfb.Is(val)
}
`
