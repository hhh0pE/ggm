package fieldType

import "fmt"

type Text struct {
	IsNullable   bool
	IsArray      bool
	IsGoBaseType bool
	MaxLength    int
}

func (t Text) Name() string {
	return name("Text", t)
}

func (t Text) SqlType() string {
	if t.MaxLength == 0 {
		return sqlType("TEXT", t)
	}
	return sqlType(fmt.Sprintf("VARCHAR(%d)", t.MaxLength), t)
}

func (t Text) WhereTemplate() string {
	if t.IsArray {
		return TextArrayTemplate
	}
	if t.IsNullable {
		return TextNullableTemplate
	}
	return TextTemplate
}

func (t Text) FmtReplacer() string {
	return "%s"
}

func (t Text) DefaultValue() string {
	if t.IsArray {
		return "nil"
	}
	return `""`
}

func (t Text) GoBaseType() string {
	if t.IsGoBaseType {
		return goType("string", t)
	}
	return t.GoScannerType()
}

func (t Text) GoScannerType() string {
	if t.IsArray {
		if t.IsNullable {
			return "*pq.StringArray"
		} else {
			return "pq.StringArray"
		}
	}
	if t.IsNullable {
		return "sql.NullString"
	}
	return "string"
}

func (t Text) MaxSize() int {
	return t.MaxLength
}

func (t Text) Nullable() bool {
	return t.IsNullable
}

func (t Text) Array() bool {
	return t.IsArray
}

func (t Text) ConstType() ConstFieldType {
	return TextType
}
func (t Text) ImplementScannerInterface() bool {
	return t.GoBaseType() == t.GoScannerType()
}

const TextTemplate = `
type whereFieldText{{.ModelName}} struct {
	name  string
	where *{{lower .ModelName}}Where
}

func (wfs *whereFieldText{{.ModelName}}) Is(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" = '" + val + "'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) IsNot(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" <> '" + val + "'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) Eq(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" = '" + val + "'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) Like(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" LIKE '" + val + "'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) NotLike(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" NOT LIKE '" + val + "'")
	return wfs.where
}

func (wfs *whereFieldText{{.ModelName}}) ILike(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" ILIKE '" + val + "'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) NotILike(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" NOT ILIKE '" + val + "'")
	return wfs.where
}

func (wfs *whereFieldText{{.ModelName}}) HasPrefix(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" LIKE '" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) NotHasPrefix(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" NOT LIKE '" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) IHasPrefix(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" ILIKE '" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) NotIHasPrefix(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" NOT ILIKE '" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) HasSuffix(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" LIKE '" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) NotHasSuffix(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" NOT LIKE '" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) IHasSuffix(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" ILIKE '%" + val + "'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) Contains(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" LIKE '%" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) NotContains(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" NOT LIKE '%" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) NotIContains(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" ILIKE '%" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) IContains(val string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" ILIKE '%" + val + "%'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) Any(val ...string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" IN ('" + strings.Join(val, "', '") + "')'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) NotAny(val ...string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" NOT IN ('" + strings.Join(val, "', '") + "')'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) In(val []string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" IN ('" + strings.Join(val, "', '") + "')'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) NotIn(val []string) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("\"" + wfs.name + "\" NOT IN ('" + strings.Join(val, "', '") + "')'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) LengthIs(len int) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("length(\"" + wfs.name + "\") = '"+fmt.Sprintf("%d", len)+"'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) LengthLessThan(len int) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("length(\"" + wfs.name + "\") < '"+fmt.Sprintf("%d", len)+"'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) LengthLT(len int) *{{lower .ModelName}}Where {
	return wfs.LengthLessThan(len)
}
func (wfs *whereFieldText{{.ModelName}}) LengthGreaterThan(len int) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("length(\"" + wfs.name + "\") > '"+fmt.Sprintf("%d", len)+"'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) LengthGT(len int) *{{lower .ModelName}}Where {
	return wfs.LengthGreaterThan(len)
}
func (wfs *whereFieldText{{.ModelName}}) LengthGreaterThanOrEqual(len int) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("length(\"" + wfs.name + "\") >= '"+fmt.Sprintf("%d", len)+"'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) LengthGTE(len int) *{{lower .ModelName}}Where {
	return wfs.LengthGreaterThanOrEqual(len)
}
func (wfs *whereFieldText{{.ModelName}}) LengthLessThanOrEqual(len int) *{{lower .ModelName}}Where {
	wfs.where.andOr()
	wfs.where.addCond("length(\"" + wfs.name + "\") <= '"+fmt.Sprintf("%d", len)+"'")
	return wfs.where
}
func (wfs *whereFieldText{{.ModelName}}) LengthLTE(len int) *{{lower .ModelName}}Where {
	return wfs.LengthLessThanOrEqual(len)
}
`

const TextNullableTemplate = `
type whereFieldTextNullable{{.ModelName}} struct {
	whereFieldText{{.ModelName}}
}

func (wfsn *whereFieldTextNullable{{.ModelName}}) IsNull() *{{lower .ModelName}}Where {
	wfsn.where.andOr()
	wfsn.where.addCond("\"" + wfsn.name + "\" IS NULL")
	return wfsn.where
}
func (wfsn *whereFieldTextNullable{{.ModelName}}) IsNotNull() *{{lower .ModelName}}Where {
	wfsn.where.andOr()
	wfsn.where.addCond("\"" + wfsn.name + "\" IS NOT NULL")
	return wfsn.where
}
`

const TextArrayTemplate = `
type whereFieldTextArray{{.ModelName}} struct {
	name string
	where *{{lower .ModelName}}Where
}

func(wfba whereFieldTextArray{{.ModelName}}) Is(val []string) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" = '"+stringArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldTextArray{{.ModelName}}) IsNot(val []string) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <> '"+stringArrayToSqlValue(val)+"'")
return wfba.where
}
func(wfba whereFieldTextArray{{.ModelName}}) LessThan(val []string) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" < '"+stringArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldTextArray{{.ModelName}}) LT(val []string) *{{lower .ModelName}}Where {
	return wfba.LessThan(val)
}
func(wfba whereFieldTextArray{{.ModelName}}) LessThanOrEqual(val []string) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <= '"+stringArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldTextArray{{.ModelName}}) LTE(val []string) *{{lower .ModelName}}Where {
	return wfba.LessThanOrEqual(val)
}
func(wfba whereFieldTextArray{{.ModelName}}) GreaterThan(val []string) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" > '"+stringArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldTextArray{{.ModelName}}) GT(val []string) *{{lower .ModelName}}Where {
	return wfba.GreaterThan(val)
}
func(wfba whereFieldTextArray{{.ModelName}}) GreaterThanOrEqual(val []string) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" > '"+stringArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldTextArray{{.ModelName}}) GTE(val []string) *{{lower .ModelName}}Where {
	return wfba.GreaterThanOrEqual(val)
}
func(wfba whereFieldTextArray{{.ModelName}}) Contains(val string) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" @> '"+stringArrayToSqlValue([]string{val})+"'")
	return wfba.where
}
func(wfba whereFieldTextArray{{.ModelName}}) ContainedBy(val []string) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" <@ '"+stringArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldTextArray{{.ModelName}}) Overlap(val []string) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("\"" + wfba.name + "\" && '"+stringArrayToSqlValue(val)+"'")
	return wfba.where
}
func(wfba whereFieldTextArray{{.ModelName}}) LengthIs(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("array_length(\"" + wfba.name + "\", 1) = '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldTextArray{{.ModelName}}) LengthLessThan(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("array_length(\"" + wfba.name + "\", 1) < '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldTextArray{{.ModelName}}) LengthLT(len int) *{{lower .ModelName}}Where {
	return wfba.LengthLessThan(len)
}
func(wfba whereFieldTextArray{{.ModelName}}) LengthLessThanOrEqual(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("array_length(\"" + wfba.name + "\", 1) <= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldTextArray{{.ModelName}}) LengthLTE(len int) *{{lower .ModelName}}Where {
	return wfba.LengthLessThanOrEqual(len)
}
func(wfba whereFieldTextArray{{.ModelName}}) LengthGreaterThan(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("array_length(\"" + wfba.name + "\", 1) > '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldTextArray{{.ModelName}}) LengthGT(len int) *{{lower .ModelName}}Where {
	return wfba.LengthGreaterThan(len)
}
func(wfba whereFieldTextArray{{.ModelName}}) LengthGreaterThanOrEqual(len int) *{{lower .ModelName}}Where {
	wfba.where.andOr()
	wfba.where.addCond("array_length(\"" + wfba.name + "\", 1) >= '"+fmt.Sprintf("%d", len)+"'")
	return wfba.where
}
func(wfba whereFieldTextArray{{.ModelName}}) LengthGTE(len int) *{{lower .ModelName}}Where {
	return wfba.LengthGreaterThanOrEqual(len)
}
`
