package fieldType

type Date struct {
	IsNullable   bool
	IsArray      bool
	IsGoBaseType bool
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


// TODO: DateArrayTemplate!!
const DateArrayTemplate = `
`