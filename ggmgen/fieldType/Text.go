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
{{if not .}}
type whereFieldText struct {
	name  string
	where modelWhere
}
func(wft whereFieldText) sqlName() string {
	return wft.name
}
func (wft whereFieldText) is(val string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " = '" + val + "'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) isNot(val string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " <> '" + val + "'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) eq(val string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " = '" + val + "'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) like(val string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " LIKE '" + val + "'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) notLike(val string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " NOT LIKE '" + val + "'")
	return wft.where.modelWhere()
}

func (wft whereFieldText) ilike(val string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " ILIKE '" + val + "'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) notILike(val string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " NOT ILIKE '" + val + "'")
	return wft.where.modelWhere()
}

func (wft whereFieldText) hasPrefix(val string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " LIKE '" + val + "%'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) notHasPrefix(val string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " NOT LIKE '" + val + "%'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) ihasPrefix(val string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " ILIKE '" + val + "%'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) notIHasPrefix(val string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " NOT ILIKE '" + val + "%'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) hasSuffix(val string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " LIKE '" + val + "%'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) notHasSuffix(val string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " NOT LIKE '" + val + "%'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) ihasSuffix(val string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " ILIKE '%" + val + "'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) contains(val string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " LIKE '%" + val + "%'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) notContains(val string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " NOT LIKE '%" + val + "%'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) notIContains(val string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " ILIKE '%" + val + "%'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) icontains(val string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " ILIKE '%" + val + "%'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) any(val ...string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " IN ('" + strings.Join(val, "', '") + "')'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) notAny(val ...string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " NOT IN ('" + strings.Join(val, "', '") + "')'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) in(val []string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " IN ('" + strings.Join(val, "', '") + "')'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) notIn(val []string) modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " NOT IN ('" + strings.Join(val, "', '") + "')'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) lengthIs(len int) modelWhere {
	wft.where.andOr()
	wft.where.addCond("length(" + wft.sqlName() + ") = '"+fmt.Sprintf("%d", len)+"'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) lengthLessThan(len int) modelWhere {
	wft.where.andOr()
	wft.where.addCond("length(" + wft.sqlName() + ") < '"+fmt.Sprintf("%d", len)+"'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) lengthGreaterThan(len int) modelWhere {
	wft.where.andOr()
	wft.where.addCond("length(" + wft.sqlName() + ") > '"+fmt.Sprintf("%d", len)+"'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) lengthGreaterThanOrEqual(len int) modelWhere {
	wft.where.andOr()
	wft.where.addCond("length(" + wft.sqlName() + ") >= '"+fmt.Sprintf("%d", len)+"'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) lengthLessThanOrEqual(len int) modelWhere {
	wft.where.andOr()
	wft.where.addCond("length(" + wft.sqlName() + ") <= '"+fmt.Sprintf("%d", len)+"'")
	return wft.where.modelWhere()
}
func (wft whereFieldText) isNull() modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " IS NULL")
	return wft.where.modelWhere()
}
func (wft whereFieldText) isNotNull() modelWhere {
	wft.where.andOr()
	wft.where.addCond(wft.sqlName() + " IS NOT NULL")
	return wft.where.modelWhere()
}
{{else}}
type whereFieldText{{.ModelName}} struct {
	whereFieldText
}
func (wft whereFieldText{{.ModelName}}) Is(val string) *{{lower .ModelName}}Where {
	return wft.is(val).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) IsNot(val string) *{{lower .ModelName}}Where {
	return wft.isNot(val).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) Eq(val string) *{{lower .ModelName}}Where {
	return wft.eq(val).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) Like(val string) *{{lower .ModelName}}Where {
	return wft.like(val).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) NotLike(val string) *{{lower .ModelName}}Where {
	return wft.notLike(val).(*{{lower .ModelName}}Where)
}

func (wft whereFieldText{{.ModelName}}) ILike(val string) *{{lower .ModelName}}Where {
	return wft.ilike(val).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) NotILike(val string) *{{lower .ModelName}}Where {
	return wft.notILike(val).(*{{lower .ModelName}}Where)
}

func (wft whereFieldText{{.ModelName}}) HasPrefix(val string) *{{lower .ModelName}}Where {
	return wft.hasPrefix(val).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) NotHasPrefix(val string) *{{lower .ModelName}}Where {
	return wft.notHasPrefix(val).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) IHasPrefix(val string) *{{lower .ModelName}}Where {
	return wft.ihasPrefix(val).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) NotIHasPrefix(val string) *{{lower .ModelName}}Where {
	return wft.notIHasPrefix(val).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) HasSuffix(val string) *{{lower .ModelName}}Where {
	return wft.hasSuffix(val).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) NotHasSuffix(val string) *{{lower .ModelName}}Where {
	return wft.notHasSuffix(val).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) IHasSuffix(val string) *{{lower .ModelName}}Where {
	return wft.ihasSuffix(val).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) Contains(val string) *{{lower .ModelName}}Where {
	return wft.contains(val).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) NotContains(val string) *{{lower .ModelName}}Where {
	return wft.notContains(val).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) NotIContains(val string) *{{lower .ModelName}}Where {
	return wft.notIContains(val).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) IContains(val string) *{{lower .ModelName}}Where {
	return wft.icontains(val).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) Any(val ...string) *{{lower .ModelName}}Where {
	return wft.any(val...).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) NotAny(val ...string) *{{lower .ModelName}}Where {
	return wft.notAny(val...).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) In(val []string) *{{lower .ModelName}}Where {
	return wft.in(val).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) NotIn(val []string) *{{lower .ModelName}}Where {
	return wft.notIn(val).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) LengthIs(len int) *{{lower .ModelName}}Where {
	return wft.lengthIs(len).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) LengthLessThan(len int) *{{lower .ModelName}}Where {
	return wft.lengthLessThan(len).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) LengthLT(len int) *{{lower .ModelName}}Where {
	return wft.LengthLessThan(len)
}
func (wft whereFieldText{{.ModelName}}) LengthGreaterThan(len int) *{{lower .ModelName}}Where {
	return wft.lengthGreaterThan(len).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) LengthGT(len int) *{{lower .ModelName}}Where {
	return wft.LengthGreaterThan(len)
}
func (wft whereFieldText{{.ModelName}}) LengthGreaterThanOrEqual(len int) *{{lower .ModelName}}Where {
	return wft.lengthGreaterThanOrEqual(len).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) LengthGTE(len int) *{{lower .ModelName}}Where {
	return wft.LengthGreaterThanOrEqual(len)
}
func (wft whereFieldText{{.ModelName}}) LengthLessThanOrEqual(len int) *{{lower .ModelName}}Where {
	return wft.lengthLessThanOrEqual(len).(*{{lower .ModelName}}Where)
}
func (wft whereFieldText{{.ModelName}}) LengthLTE(len int) *{{lower .ModelName}}Where {
	return wft.LengthLessThanOrEqual(len)
}
{{end}}
`

const TextNullableTemplate = `
{{if .}}
type whereFieldTextNullable{{.ModelName}} struct {
	whereFieldText{{.ModelName}}
}

func (wftn whereFieldTextNullable{{.ModelName}}) IsNull() *{{lower .ModelName}}Where {
	return wftn.isNull().(*{{lower .ModelName}}Where)
}
func (wftn whereFieldTextNullable{{.ModelName}}) IsNotNull() *{{lower .ModelName}}Where {
	return wftn.isNotNull().(*{{lower .ModelName}}Where)
}
{{end}}
`

const TextArrayTemplate = `
{{if not .}}
type whereFieldTextArray struct {
	name string
	where modelWhere
}
func(wfta whereFieldTextArray) sqlName() string {
	return wfta.name
}
func(wfta whereFieldTextArray) is(val []string) modelWhere {
	wfta.where.andOr()
	wfta.where.addCond(wfta.sqlName() + " = '"+stringArrayToSqlValue(val)+"'")
	return wfta.where.modelWhere()
}
func(wfta whereFieldTextArray) isNot(val []string) modelWhere {
	wfta.where.andOr()
	wfta.where.addCond(wfta.sqlName() + " <> '"+stringArrayToSqlValue(val)+"'")
	return wfta.where.modelWhere()
}
func(wfta whereFieldTextArray) lessThan(val []string) modelWhere {
	wfta.where.andOr()
	wfta.where.addCond(wfta.sqlName() + " < '"+stringArrayToSqlValue(val)+"'")
	return wfta.where.modelWhere()
}
func(wfta whereFieldTextArray) lessThanOrEqual(val []string) modelWhere {
	wfta.where.andOr()
	wfta.where.addCond(wfta.sqlName() + " <= '"+stringArrayToSqlValue(val)+"'")
	return wfta.where.modelWhere()
}
func(wfta whereFieldTextArray) greaterThan(val []string) modelWhere {
	wfta.where.andOr()
	wfta.where.addCond(wfta.sqlName() + " > '"+stringArrayToSqlValue(val)+"'")
	return wfta.where.modelWhere()
}
func(wfta whereFieldTextArray) greaterThanOrEqual(val []string) modelWhere {
	wfta.where.andOr()
	wfta.where.addCond(wfta.sqlName() + " > '"+stringArrayToSqlValue(val)+"'")
	return wfta.where.modelWhere()
}
func(wfta whereFieldTextArray) contains(val string) modelWhere {
	wfta.where.andOr()
	wfta.where.addCond(wfta.sqlName() + " @> '"+stringArrayToSqlValue([]string{val})+"'")
	return wfta.where.modelWhere()
}
func(wfta whereFieldTextArray) containedBy(val []string) modelWhere {
	wfta.where.andOr()
	wfta.where.addCond(wfta.sqlName() + " <@ '"+stringArrayToSqlValue(val)+"'")
	return wfta.where.modelWhere()
}
func(wfta whereFieldTextArray) overlap(val []string) modelWhere {
	wfta.where.andOr()
	wfta.where.addCond(wfta.sqlName() + " && '"+stringArrayToSqlValue(val)+"'")
	return wfta.where.modelWhere()
}
func(wfta whereFieldTextArray) lengthIs(len int) modelWhere {
	wfta.where.andOr()
	wfta.where.addCond("array_length(" + wfta.sqlName() + ", 1) = '"+fmt.Sprintf("%d", len)+"'")
	return wfta.where.modelWhere()
}
func(wfta whereFieldTextArray) lengthLessThan(len int) modelWhere {
	wfta.where.andOr()
	wfta.where.addCond("array_length(" + wfta.sqlName() + ", 1) < '"+fmt.Sprintf("%d", len)+"'")
	return wfta.where.modelWhere()
}
func(wfta whereFieldTextArray) lengthLessThanOrEqual(len int) modelWhere {
	wfta.where.andOr()
	wfta.where.addCond("array_length(" + wfta.sqlName() + ", 1) <= '"+fmt.Sprintf("%d", len)+"'")
	return wfta.where.modelWhere()
}
func(wfta whereFieldTextArray) lengthGreaterThan(len int) modelWhere {
	wfta.where.andOr()
	wfta.where.addCond("array_length(" + wfta.sqlName() + ", 1) > '"+fmt.Sprintf("%d", len)+"'")
	return wfta.where.modelWhere()
}
func(wfta whereFieldTextArray) lengthGreaterThanOrEqual(len int) modelWhere {
	wfta.where.andOr()
	wfta.where.addCond("array_length(" + wfta.sqlName() + ", 1) >= '"+fmt.Sprintf("%d", len)+"'")
	return wfta.where.modelWhere()
}
{{else}}
type whereFieldTextArray{{.ModelName}} struct {
	whereFieldTextArray
}
func(wfta whereFieldTextArray{{.ModelName}}) Is(val []string) *{{lower .ModelName}}Where {
	return wfta.is(val).(*{{lower .ModelName}}Where)
}
func(wfta whereFieldTextArray{{.ModelName}}) IsNot(val []string) *{{lower .ModelName}}Where {
	return wfta.isNot(val).(*{{lower .ModelName}}Where)
}
func(wfta whereFieldTextArray{{.ModelName}}) LessThan(val []string) *{{lower .ModelName}}Where {
	return wfta.lessThan(val).(*{{lower .ModelName}}Where)
}
func(wfta whereFieldTextArray{{.ModelName}}) LT(val []string) *{{lower .ModelName}}Where {
	return wfta.LessThan(val)
}
func(wfta whereFieldTextArray{{.ModelName}}) LessThanOrEqual(val []string) *{{lower .ModelName}}Where {
	return wfta.lessThanOrEqual(val).(*{{lower .ModelName}}Where)
}
func(wfta whereFieldTextArray{{.ModelName}}) LTE(val []string) *{{lower .ModelName}}Where {
	return wfta.LessThanOrEqual(val)
}
func(wfta whereFieldTextArray{{.ModelName}}) GreaterThan(val []string) *{{lower .ModelName}}Where {
	return wfta.greaterThan(val).(*{{lower .ModelName}}Where)
}
func(wfta whereFieldTextArray{{.ModelName}}) GT(val []string) *{{lower .ModelName}}Where {
	return wfta.GreaterThan(val)
}
func(wfta whereFieldTextArray{{.ModelName}}) GreaterThanOrEqual(val []string) *{{lower .ModelName}}Where {
	return wfta.greaterThanOrEqual(val).(*{{lower .ModelName}}Where)
}
func(wfta whereFieldTextArray{{.ModelName}}) GTE(val []string) *{{lower .ModelName}}Where {
	return wfta.GreaterThanOrEqual(val)
}
func(wfta whereFieldTextArray{{.ModelName}}) Contains(val string) *{{lower .ModelName}}Where {
	return wfta.contains(val).(*{{lower .ModelName}}Where)
}
func(wfta whereFieldTextArray{{.ModelName}}) ContainedBy(val []string) *{{lower .ModelName}}Where {
	return wfta.containedBy(val).(*{{lower .ModelName}}Where)
}
func(wfta whereFieldTextArray{{.ModelName}}) Overlap(val []string) *{{lower .ModelName}}Where {
	return wfta.overlap(val).(*{{lower .ModelName}}Where)
}
func(wfta whereFieldTextArray{{.ModelName}}) LengthIs(len int) *{{lower .ModelName}}Where {
	return wfta.lengthIs(len).(*{{lower .ModelName}}Where)
}
func(wfta whereFieldTextArray{{.ModelName}}) LengthLessThan(len int) *{{lower .ModelName}}Where {
	return wfta.lengthLessThan(len).(*{{lower .ModelName}}Where)
}
func(wfta whereFieldTextArray{{.ModelName}}) LengthLT(len int) *{{lower .ModelName}}Where {
	return wfta.LengthLessThan(len)
}
func(wfta whereFieldTextArray{{.ModelName}}) LengthLessThanOrEqual(len int) *{{lower .ModelName}}Where {
	return wfta.lengthLessThanOrEqual(len).(*{{lower .ModelName}}Where)
}
func(wfta whereFieldTextArray{{.ModelName}}) LengthLTE(len int) *{{lower .ModelName}}Where {
	return wfta.LengthLessThanOrEqual(len)
}
func(wfta whereFieldTextArray{{.ModelName}}) LengthGreaterThan(len int) *{{lower .ModelName}}Where {
	return wfta.lengthGreaterThan(len).(*{{lower .ModelName}}Where)
}
func(wfta whereFieldTextArray{{.ModelName}}) LengthGT(len int) *{{lower .ModelName}}Where {
	return wfta.LengthGreaterThan(len)
}
func(wfta whereFieldTextArray{{.ModelName}}) LengthGreaterThanOrEqual(len int) *{{lower .ModelName}}Where {
	return wfta.lengthGreaterThanOrEqual(len).(*{{lower .ModelName}}Where)
}
func(wfta whereFieldTextArray{{.ModelName}}) LengthGTE(len int) *{{lower .ModelName}}Where {
	return wfta.LengthGreaterThanOrEqual(len)
}
{{end}}
`
