package fieldType

type Date struct {
	IsNullable      bool
	IsArray         bool
	IsGoBaseType    bool
	WithoutTimezone bool
}

func (d Date) Name() string {
	return name("Date", d)
}

func (d Date) SqlType() string {
	if d.WithoutTimezone {
		return sqlType("TIMESTAMP", d)
	}
	return sqlType("TIMESTAMPTZ", d)
}

func (d Date) WhereTemplate() string {
	if d.IsArray {
		return DateArrayTemplate
	}
	if d.IsNullable {
		return DateNullableTemplate
	}
	return DateTemplate
}

func (d Date) FmtReplacer() string {
	return "%s"
}

func (d Date) DefaultValue() string {
	return "0"
}

func (d Date) GoBaseType() string {
	if d.IsGoBaseType {
		return goType("time.Time{}", d)
	}
	return d.GoScannerType()
}

func (d Date) GoScannerType() string {
	if d.IsArray {
		if d.IsNullable {
			return "*ggm.DateArray"
		} else {
			return "ggm.DateArray"
		}
	}
	if d.IsNullable {
		return "pq.NullTime"
	}
	return "time.Time"
}

func (d Date) MaxSize() int {
	return 0
}

func (d Date) Nullable() bool {
	return d.IsNullable
}

func (d Date) Array() bool {
	return d.IsArray
}
func (d Date) ConstType() ConstFieldType {
	return DateType
}
func (d Date) ImplementScannerInterface() bool {
	return d.GoBaseType() == d.GoScannerType()
}

const DateTemplate = `
{{if not .}}
type whereFieldDate struct {
	name  string
	where modelWhere
}
func(wfd whereFieldDate) sqlName() string {
	if strings.HasPrefix(wfd.name, "\"") { // already has table name
		return wfd.name
	}
	return "\""+wfd.name+"\""
}
func (wfd whereFieldDate) is(d time.Time) modelWhere {
	wfd.where.andOr()
	wfd.where.addCond("\"" + wfd.name + "\" = '" + d.Format("2006-02-01") + "'")
	return wfd.where.modelWhere()
}
func (wfd whereFieldDate) isNull() modelWhere {
	wfd.where.andOr()
	wfd.where.addCond(wfd.sqlName() + " IS NULL")
	return wfd.where.modelWhere()
}
func (wfd whereFieldDate) isNotNull() modelWhere {
	wfd.where.andOr()
	wfd.where.addCond(wfd.sqlName() + " IS NOT NULL")
	return wfd.where.modelWhere()
}
{{else}}
type whereFieldDate{{.ModelName}} struct {
	whereFieldDate
}
func (wfd whereFieldDate{{.ModelName}}) Is(d time.Time) *{{lower .ModelName}}Where {
	return wfd.is(d).(*{{lower .ModelName}}Where)
}
{{end}}
`

const DateNullableTemplate = `
{{if .}}
type whereFieldDateNullable{{.ModelName}} struct {
	whereFieldDate{{.ModelName}}
}
func (wfd whereFieldDateNullable{{.ModelName}}) IsNull() *{{lower .ModelName}}Where {
	return wfd.isNull().(*{{lower .ModelName}}Where)
}
func (wfd whereFieldDateNullable{{.ModelName}}) IsNotNull() *{{lower .ModelName}}Where {
	return wfd.isNotNull().(*{{lower .ModelName}}Where)
}
{{end}}
`

// TODO: DateArrayTemplate!!
const DateArrayTemplate = `
`
