package fields

const DateTemplate = `
type whereFieldDate{{.ModelName}} struct {
	name  string
	where *{{lower .ModelName}}Where
}

func (wfd whereFieldDate{{.ModelName}}) Is(d time.Time) *{{lower .ModelName}}Where {
	wfd.where.andOr()
	wfd.where.addCond("\"" + wfd.name + "\" = '" + d.Format("2006-02-01") + "'")
	return wfd.where
}
`
