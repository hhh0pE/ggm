package fields

const StringTemplate = `
type whereFieldString{{.ModelName}} struct {
	name  string
	where *{{lower .ModelName}}Where
}

func (wfs *whereFieldString{{.ModelName}}) Is(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" = '" + val + "'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) IsNot(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" <> '" + val + "'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) Eq(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" = '" + val + "'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) Like(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" LIKE '" + val + "'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) NotLike(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" NOT LIKE '" + val + "'")
	return wfs.where
}

func (wfs *whereFieldString{{.ModelName}}) ILike(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" ILIKE '" + val + "'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) NotILike(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" NOT ILIKE '" + val + "'")
	return wfs.where
}

func (wfs *whereFieldString{{.ModelName}}) HasPrefix(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" LIKE '" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) NotHasPrefix(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" NOT LIKE '" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) IHasPrefix(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" ILIKE '" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) NotIHasPrefix(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" NOT ILIKE '" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) HasSuffix(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" LIKE '" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) NotHasSuffix(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" NOT LIKE '" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) IHasSuffix(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" ILIKE '%" + val + "'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) Contains(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" LIKE '%" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) NotContains(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" NOT LIKE '%" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) NotIContains(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" ILIKE '%" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) IContains(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" ILIKE '%" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) Any(val ...string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" IN ('" + strings.Join(val, "', '") + "')'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) NotAny(val ...string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" NOT IN ('" + strings.Join(val, "', '") + "')'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) In(val []string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" IN ('" + strings.Join(val, "', '") + "')'")
	return wfs.where
}
func (wfs *whereFieldString{{.ModelName}}) NotIn(val []string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" NOT IN ('" + strings.Join(val, "', '") + "')'")
	return wfs.where
}
`
