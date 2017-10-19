package templates

const ModelTemplate = `

// model {{.Name}} ORM part
{{if not .IsTableNameSetByUser}}
func ({{abbr .Name}} {{.Name}}) TableName() string {
	return "{{.TableName}}"
}
{{end}}

func isEmpty{{.Name}}({{abbr .Name}} {{.Name}}) bool {
	{{- range $field := .Fields}}
	if {{if $field.Type.IsNullable}}{{abbr $.Name}}.{{$field.Name}} != nil && {{end -}}
	{{$field.FieldValueName (abbr $.Name)}} != {{$field.DefaultValue}} {
		return false
	}
	{{- end}}
	return true
}

func save{{.Name}}({{abbr .Name}} *{{.Name}}) error {
	if inserting_err := insert{{.Name}}({{abbr .Name}}); inserting_err != nil {
		if pqInsertingErr, ok := inserting_err.(*pq.Error); ok {
			if pqInsertingErr.Code.Name() != "unique_violation" {
				return pqInsertingErr
			}
		}

		var setStatement, whereClause []string

		if {{range $pki, $pk := .PrimaryKeys}}{{abbr $.Name}}.{{$pk.Name}} != {{$pk.DefaultValue}}{{if IsNotLastElement $pki (len $.PrimaryKeys)}} && {{end}}{{end}} {
			{{range $pki, $pk := .PrimaryKeys}}whereClause = append(whereClause, fmt.Sprintf("\"{{$pk.TableName}}\" = '{{$pk.Type.FmtReplacer}}'", {{$pk.FieldValueName (abbr $.Name)}})){{end}}
		} else {
	{{if gt (len .Indexes) 0}}
			if {{range $ii, $index := .Indexes -}}
			{{range $ifi, $indexField := $index.Fields -}}
				{{- if and $indexField.Type.IsNullable (not $indexField.IsForeignKey)}}{{abbr $.Name}}.{{$indexField.Name}} != nil && {{end -}}
				{{- $indexField.FieldValueName (abbr $.Name)}} != {{$indexField.DefaultValue -}}
				{{- if IsNotLastElement $ifi (len $index.Fields)}} && {{end -}}
			{{end}} {
			{{range $ifi, $indexField := $index.Fields -}}
			{{if true}}	{{end}}whereClause = append(whereClause, fmt.Sprintf("\"{{$indexField.TableName}}\" = '{{$indexField.Type.FmtReplacer}}'", {{$indexField.FieldValueName (abbr $.Name)}}))
			{{end -}}
	}{{if IsNotLastElement $ii (len $.Indexes)}} else if {{end -}}
		{{- end -}}
	{{- end}}
		}

		{{range $fi, $field := .AllFields}}{{if not $field.IsPrimaryKey}}
		setStatement = append(setStatement, fmt.Sprintf("\"{{$field.TableName}}\" = '{{$field.Type.FmtReplacer}}'", {{$field.FieldValueName (abbr $.Name)}}))
		{{- end}}{{end}}

		_, err := Exec("UPDATE \"{{.TableName}}\" SET "+strings.Join(setStatement, ", ")+" WHERE "+strings.Join(whereClause, " AND "))
		if err != nil {
			return err
		}

		if {{range $pki, $pk := .PrimaryKeys}}{{abbr $.Name}}.{{$pk.Name}} == {{$pk.DefaultValue}}{{if IsNotLastElement $pki (len $.PrimaryKeys)}} || {{end}}{{end}} {
			var selectPKFields []string
			{{range $pk := .PrimaryKeys}}selectPKFields = append(selectPKFields, "\"{{$pk.TableName}}\""){{end}}
			row := QueryRow("SELECT "+strings.Join(selectPKFields, ", ")+" FROM \"{{.TableName}}\" WHERE " + strings.Join(whereClause, " AND ") + ";")

			if scanning_err := row.Scan({{range $pki, $pk := .PrimaryKeys}}&{{abbr $.Name}}.{{$pk.Name}}{{if IsNotLastElement $pki (len $.PrimaryKeys)}}, {{end}}{{end}}); scanning_err != nil {
				return scanning_err
			}
		}
		return nil

	}
	return nil
}

func insert{{.Name}} ({{abbr .Name}} *{{.Name}}) error {
	{{range $field := .Fields}}{{if $field.IsPrimaryKey}}if {{abbr $.Name}}.{{$field.Name}} != {{$field.Type.DefaultValue}} {
		return errors.New("Cannot insert {{$.Name}} with {{$field.Name}} != {{$field.Type.DefaultValue}}")
	}
	{{end}}{{end}}

	var fieldTableNames, fieldValues []string
	{{range $field := .AllFields -}}
		{{if not $field.IsPrimaryKey}}
	fieldTableNames = append(fieldTableNames, "\"{{$field.TableName}}\"")
	fieldValues = append(fieldValues, fmt.Sprintf("'{{$field.Type.FmtReplacer}}'", {{$field.FieldValueName (abbr $.Name)}}))
		{{- end -}}
	{{- end}}

	result, err := Exec("INSERT INTO \"{{.TableName}}\" ("+strings.Join(fieldTableNames, ", ")+") VALUES ("+strings.Join(fieldValues, ", ")+")")
	if err != nil {
		return err
	}
	if lastID, lastID_err := result.LastInsertId(); lastID_err != nil {
		var selectPKFields []string
		{{range $field := .Fields}}{{if $field.IsPrimaryKey}}selectPKFields = append(selectPKFields, "\"{{$field.TableName}}\""){{end}}{{end}}
		row := QueryRow("SELECT "+strings.Join(selectPKFields, ", ")+" FROM \"{{.TableName}}\" ORDER BY "+strings.Join(selectPKFields, ", ")+" DESC LIMIT 1;")
		if scanning_err := row.Scan(&{{abbr .Name}}.ID); scanning_err != nil {
			return scanning_err
		}
	} else {
		{{abbr .Name}}.ID = lastID
	}
	//fmt.Println(result.LastInsertId())
	//fmt.Println(err)
	return nil
}

func delete{{.Name}}({{abbr .Name}} *{{.Name}}) error {
	var pkFieldWhere []string
	{{range $field := .Fields}}{{if $field.IsPrimaryKey}}pkFieldWhere = append(pkFieldWhere, fmt.Sprintf("\"{{$field.TableName}}\" = '{{$field.Type.FmtReplacer}}'", {{abbr $.Name}}.{{$field.Name}}))
	{{end}}{{end}}
	if _, err := Exec("DELETE FROM \"{{.TableName}}\" WHERE "+strings.Join(pkFieldWhere, " AND ")); err != nil {
		return err
	}
	{{range $field := .Fields}}{{if $field.IsPrimaryKey}}{{abbr $.Name}}.{{$field.Name}} = {{$field.Type.DefaultValue}}
	{{end}}{{end}}
	return nil
}





func {{.Name}}Q() *{{lower .Name}}Query {
	return &{{lower .Name}}Query{}
}

func get{{.Name}}FieldPointers({{abbr .Name}} *{{.Name}}, fieldNames []string) []interface{} {
	var fieldPointers []interface{}
	for i, _ := range fieldNames {
		switch fieldNames[i] {
		{{range $field := .Fields}}case "{{$field.TableName}}":
		    fieldPointers = append(fieldPointers, &{{abbr $.Name}}.{{$field.Name}})
        {{end}}
		}
	}

	return fieldPointers
}

type {{lower .Name}}Query struct {
	{{lower .Name}}Select  {{lower .Name}}Select
	{{lower .Name}}Where   *{{lower .Name}}Where
	joins        []string
	limit           int
	{{lower .Name}}OrderBy *{{lower .Name}}OrderBy
}

func scan{{.Name}}Row({{abbr .Name}} *{{.Name}}, fieldNames []string, row *sql.Rows) error {
	var fieldPointers []interface{}
	{{range $nsf := .NotScannerFields}}
	var {{abbr $nsf.Name}}ForScan {{$nsf.Type.GoScannerType}}
	{{- end}}
	for i, _ := range fieldNames {
		switch fieldNames[i] {
	{{- range $field := .Fields}}
		case "{{$field.TableName}}":
			{{if $field.Type.ImplementScannerInterface -}}
			fieldPointers = append(fieldPointers, &{{abbr $.Name}}.{{$field.Name}})
			{{- else -}}
			fieldPointers = append(fieldPointers, &{{abbr $field.Name}}ForScan)
			{{end -}}
	{{end}}
		}
	}

	scan_err := row.Scan(fieldPointers...)
	if scan_err != nil {
		return scan_err
	}

	for i, _ := range fieldNames {
		switch fieldNames[i] {
	{{range $field := .NotScannerFields}}
		case "{{$field.TableName}}":
			{{abbr $.Name}}.{{$field.Name}} = scannerTypeToBaseType({{abbr $field.Name}}ForScan, {{abbr $.Name}}.{{$field.Name}}).({{$field.Type.GoBaseType}})
	{{end}}
		}
	}

	return nil
}


func ({{abbr .Name}} {{lower .Name}}Query) ALL() []{{.Name}} {
	rows, err := ormDB.Query({{abbr .Name}}.SQL())
	if err != nil {
		if ormDB == nil {
			panic("initialize DB connection first!")
		}
		panic(err)
	}
	var {{lower .Name}}s []{{.Name}}
	for rows.Next() {
		var new{{.Name}} {{.Name}}
		scan_err := scan{{.Name}}Row(&new{{.Name}}, {{abbr .Name}}.{{lower .Name}}Select.fieldNames(), rows)
		//scan_err := rows.Scan(get{{.Name}}FieldPointers(&new{{.Name}}, {{abbr .Name}}.{{lower .Name}}Select.fieldNames())...)
		if scan_err != nil {
			panic("scann err: " + scan_err.Error())
		}
		{{lower .Name}}s = append({{lower .Name}}s, new{{.Name}})
	}
	return {{lower .Name}}s
}

func ({{abbr .Name}} {{lower .Name}}Query) FIRST() *{{.Name}} {
    res := {{abbr .Name}}.LIMIT(1).ALL()
    if len(res) > 0 {
        return &res[0]
    }
    return nil
}
func ({{abbr .Name}} {{lower .Name}}Query) LAST() *{{.Name}} {
    res := {{abbr .Name}}.ORDERBY().Reverse().LIMIT(1).ALL()

    if len(res) > 0 {
        if len({{abbr .Name}}.ORDERBY().ordersBy) == 0 { // if no ORDER - return last in row
            return &res[len(res)-1]
        } else { // if there is order (REVERSED) - return first
            return &res[0]
        }
    }
    return nil
}

/*func ({{abbr .Name}} {{lower .Name}}Query) COUNT() int {

}*/

func ({{abbr .Name}} {{lower .Name}}Query) SQL() string {
    var ordersByStr string
    var whereStr string
    var limitStr string
	var joinStr string
    if {{abbr .Name}}.{{lower .Name}}OrderBy != nil {
        ordersByStr = {{abbr .Name}}.{{lower .Name}}OrderBy.orderBySQL()
    }
    if {{abbr .Name}}.{{lower .Name}}Where != nil {
        whereStr = " WHERE " + {{abbr .Name}}.{{lower .Name}}Where.conditionSQL()
    }

	if {{abbr .Name}}.joins != nil && len({{abbr .Name}}.joins) > 0 {
		joinUniqueMap := make(map[string]bool)
		for _, join := range {{abbr .Name}}.joins {
			if _, alreadyExist := joinUniqueMap[join]; !alreadyExist {
				joinUniqueMap[join] = true
				if joinParts := strings.Split(join, "=>"); len(joinParts) == 2 {
					if joinText, joinExist := joinMap[joinParts[0]][joinParts[1]]; joinExist {
						joinStr += joinText
					} else {
						log.Println("doesn't exist relation in joinMap :( "+join)
					}
				} else {
					log.Println("not exist join pair: "+join)
				}
			}
		}
		joinStr += "\n"
	}
    if {{abbr .Name}}.limit > 0 {
        limitStr = fmt.Sprintf(" LIMIT %d", {{abbr .Name}}.limit)
    }
    return "SELECT " + {{abbr .Name}}.{{lower .Name}}Select.fieldsSQL() + " FROM \"{{.TableName}}\"" + joinStr + whereStr + ordersByStr + limitStr
}

func ({{abbr .Name}} *{{lower .Name}}Query) LIMIT(limit int) *{{lower .Name}}Query {
    {{abbr .Name}}.limit = limit
    return {{abbr .Name}}
}

func ({{abbr .Name}} *{{lower .Name}}Query) SELECT() *{{lower .Name}}Select {
	{{abbr .Name}}.{{lower .Name}}Select.query = {{abbr .Name}}
	return &{{abbr .Name}}.{{lower .Name}}Select
}

func ({{abbr .Name}} *{{lower .Name}}Query) WHERE() *{{lower .Name}}Where {
    if {{abbr .Name}}.{{lower .Name}}Where == nil {
        {{abbr .Name}}.{{lower .Name}}Where = new({{lower .Name}}Where)
        {{range $field := .Fields}}{{abbr $.Name}}.{{lower $.Name}}Where.{{$field.Name}}.where = {{abbr $.Name}}.{{lower $.Name}}Where
        {{abbr $.Name}}.{{lower $.Name}}Where.{{$field.Name}}.name = "\"{{$field.Model.TableName}}\".\"{{$field.TableName}}\""
        {{end}}
        {{abbr .Name}}.{{lower .Name}}Where.query = {{abbr .Name}}

	{{$prefix := print (abbr .Name) "." (lower .Name) "Where"}}
	{{template "relationWhereInitialize" dict "Model" . "Prefix" $prefix "WhereName" $prefix "Joins" "" "ExcludeModelName" ""}}
    }

	return {{abbr .Name}}.{{lower .Name}}Where
}

{{define "stringSliceToText"}}
	{{- range $elem := . -}}
		"{{.}}",
	{{- end -}}
{{end}}

func ({{abbr .Name}} *{{lower .Name}}Query) ORDERBY() *{{lower .Name}}OrderBy {
    if {{abbr .Name}}.{{lower .Name}}OrderBy == nil {
        {{abbr .Name}}.{{lower .Name}}OrderBy = new({{lower .Name}}OrderBy)
    	{{abbr .Name}}.{{lower .Name}}OrderBy.query = {{abbr .Name}}
    }
	return {{abbr .Name}}.{{lower .Name}}OrderBy
}

type {{lower .Name}}OrderBy struct {
	ordersBy []string
	query *{{lower .Name}}Query
}
type {{lower .Name}}OrderBySelectedField struct {
    *{{lower .Name}}OrderBy
}
func ({{abbr .Name}} {{lower .Name}}OrderBySelectedField) ASC() *{{lower .Name}}OrderBy {
    {{abbr .Name}}.ordersBy[len({{abbr .Name}}.ordersBy)-1] = strings.TrimSpace(strings.TrimSuffix(strings.TrimSuffix({{abbr .Name}}.ordersBy[len({{abbr .Name}}.ordersBy)-1], "ASC"), "DESC")) + " ASC"
    return {{abbr .Name}}.{{lower .Name}}OrderBy
}
func ({{abbr .Name}} {{lower .Name}}OrderBySelectedField) DESC() *{{lower .Name}}OrderBy {
    {{abbr .Name}}.ordersBy[len({{abbr .Name}}.ordersBy)-1] = strings.TrimSpace(strings.TrimSuffix(strings.TrimSuffix({{abbr .Name}}.ordersBy[len({{abbr .Name}}.ordersBy)-1], "ASC"), "DESC")) + " DESC"
    return {{abbr .Name}}.{{lower .Name}}OrderBy
}
func ({{abbr $.Name}} *{{lower $.Name}}OrderBy) SELECT() *{{lower $.Name}}Select {
    return {{abbr $.Name}}.query.SELECT()
}
func ({{abbr $.Name}} *{{lower $.Name}}OrderBy) WHERE() *{{lower $.Name}}Where {
    return {{abbr $.Name}}.query.WHERE()
}
func ({{abbr $.Name}} *{{lower $.Name}}OrderBy) SQL() string {
    return {{abbr $.Name}}.query.SQL()
}
func ({{abbr $.Name}} *{{lower $.Name}}OrderBy) orderBySQL() string {
    if len({{abbr $.Name}}.ordersBy) > 0 {
        return " ORDER BY "+strings.Join({{abbr $.Name}}.ordersBy, ", ")
    }
    return ""
}
func ({{abbr .Name}} *{{lower .Name}}OrderBy) ALL() []{{.Name}} {
    return {{abbr .Name}}.query.ALL()
}
func ({{abbr .Name}} *{{lower .Name}}OrderBy) FIRST() *{{.Name}} {
    return {{abbr .Name}}.query.FIRST()
}
func ({{abbr .Name}} *{{lower .Name}}OrderBy) LAST() *{{.Name}} {
    return {{abbr .Name}}.query.LAST()
}
func ({{abbr .Name}} *{{lower .Name}}OrderBy) Reverse() *{{lower .Name}}OrderBy {
	for i, order := range {{abbr .Name}}.ordersBy {
		if strings.HasSuffix(order, "DESC") {
			{{abbr .Name}}.ordersBy[i] = strings.TrimSuffix(order, "DESC") + "ASC"
		} else {
			{{abbr .Name}}.ordersBy[i] = strings.TrimSuffix(order, "ASC") + "DESC"
		}
	}
	return {{abbr .Name}}
}
func ({{abbr .Name}} *{{lower .Name}}OrderBy) LIMIT(limit int) *{{lower .Name}}Query {
    return {{abbr .Name}}.query.LIMIT(limit)
}

{{range $field := .Fields}}
func ({{abbr $.Name}} *{{lower $.Name}}OrderBy) {{$field.Name}}() *{{lower $.Name}}OrderBySelectedField {
    {{abbr $.Name}}.ordersBy = append({{abbr $.Name}}.ordersBy, "\"{{$field.Model.TableName}}\".\"{{.TableName}}\" ASC")
    return &{{lower $.Name}}OrderBySelectedField{ {{abbr $.Name}} }
}
{{end}}

type {{lower .Name}}Select struct {
	query  *{{lower .Name}}Query
	fields []string
}

/*
type {{lower .Name}}SelectField struct {
    *{{lower .Name}}Select

    {{range $field := .Fields}}{{$field.Name}} *{{lower .Name}}Select{{end}}
}*/

func ({{abbr .Name}} {{lower .Name}}Select) fieldNames() []string {
	if len({{abbr .Name}}.fields) == 0 {
		return []string{ {{range $field := $.Fields}}"{{$field.TableName}}",{{end}} }
	}
	return {{abbr .Name}}.fields
}
func ({{abbr .Name}} {{lower .Name}}Select) fieldsSQL() string {
	var sqlFields []string
	for _, f := range {{abbr .Name}}.fieldNames() {
		sqlFields = append(sqlFields, "\"{{.TableName}}\".\""+f+"\"")
	}

	return strings.Join(sqlFields, ", ")
}
func ({{abbr .Name}} {{lower .Name}}Select) SQL() string {
    return {{abbr .Name}}.query.SQL()
}
func ({{abbr .Name}} *{{lower .Name}}Select) ALL() []{{.Name}} {
    return {{abbr .Name}}.query.ALL()
}
func ({{abbr .Name}} *{{lower .Name}}Select) FIRST() *{{.Name}} {
    return {{abbr .Name}}.query.FIRST()
}
func ({{abbr .Name}} *{{lower .Name}}Select) LAST() *{{.Name}} {
    return {{abbr .Name}}.query.LAST()
}

{{range $field := .Fields}}
func ({{abbr $.Name}} *{{lower $.Name}}Select) {{$field.Name}}() *{{lower $.Name}}Select {
    {{abbr $.Name}}.fields = append({{abbr $.Name}}.fields, "{{$field.TableName}}")
    return {{abbr $.Name}}
}
{{end}}

func ({{abbr .Name}} *{{lower .Name}}Select) WHERE() *{{lower .Name}}Where {
    return {{abbr .Name}}.query.WHERE()
}
func ({{abbr .Name}} *{{lower .Name}}Select) ORDERBY() *{{lower .Name}}OrderBy {
    return {{abbr .Name}}.query.ORDERBY()
}
func ({{abbr .Name}} *{{lower .Name}}Select) LIMIT(limit int) *{{lower .Name}}Query {
    return {{abbr .Name}}.query.LIMIT(limit)
}

{{ $model := . }}
{{range $fieldType := GetAllFieldTypes}}
	{{ExecuteFieldTypeTemplate $fieldType $model}}
{{end}}

{{template "modelWhere" .}}

{{template "modelWhereRelation" .}}



{{define "relationWhereInitialize"}}
	
	{{- range $fk := .Model.ForeignKeys}}
	{{if or (and $.ExcludeModelName (not (eq $.ExcludeModelName $fk.ModelTo.Name))) (eq $.ExcludeModelName "")}}
	{{$.Prefix}}.{{$fk.Field.Name}}.originalWhere = {{$.WhereName}}
	{{if $.Joins}}
		{{$joins := print $.Joins ", " "\"" $fk.ModelFrom.Name "=>" $fk.ModelTo.Name "\""}}
		{{$.Prefix}}.{{$fk.Field.Name}}.joins = append({{$.Prefix}}.joins, []string{ {{$joins}} }...)
		
	{{else}}
		{{$joins := print "\"" $fk.ModelFrom.Name "=>" $fk.ModelTo.Name "\""}}
		{{$.Prefix}}.{{$fk.Field.Name}}.joins = []string{ {{$joins}} }
	{{end}}

	
	
		{{- range $field := $fk.ModelTo.Fields}}
	{{$.Prefix}}.{{$fk.Field.Name}}.{{$field.Name}}.name = "\"{{$field.Model.TableName}}\".\"{{$field.TableName}}\""
	{{$.Prefix}}.{{$fk.Field.Name}}.{{$field.Name}}.where = {{$.Prefix}}.{{$fk.Field.Name}}

		{{- $newPrefix := print $.Prefix "." $fk.Field.Name -}}

		{{if $.Joins}}
			{{$joins := print $.Joins ", " "\"" $fk.ModelFrom.Name "=>" $fk.ModelTo.Name "\""}}
			{{template "relationWhereInitialize" dict "Model" $fk.ModelTo "Prefix" $newPrefix "WhereName" $.WhereName "Joins" $joins "ExcludeModelName" $.ExcludeModelName}}
		{{else}}
			{{$joins := print "\"" $fk.ModelFrom.Name "=>" $fk.ModelTo.Name "\""}}
			{{template "relationWhereInitialize" dict "Model" $fk.ModelTo "Prefix" $newPrefix "WhereName" $.WhereName "Joins" $joins "ExcludeModelName" $.ExcludeModelName}}
		{{end}}

		{{- end}}
	
	{{end}}
	{{- end -}}
{{end}}

{{define "modelWhereRelation"}}
{{range $modelRelation := .Relations}}
{{if not (eq $modelRelation.ModelTo.Name $.Name)}}
type whereRelation{{$modelRelation.ModelFrom.Name}}_{{$modelRelation.ModelTo.Name}} struct {
	{{range $field := $modelRelation.ModelTo.Fields}}{{$field.Name}}     whereField{{title $field.Type.Name}}{{$modelRelation.ModelFrom.Name}}
	{{end -}}
	{{range $fk := $modelRelation.ModelTo.ForeignKeys}}
		{{if not (eq $modelRelation.ModelFrom.Name $fk.ModelTo.Name)}}
		{{$fk.Field.Name}}	whereRelation{{$modelRelation.ModelFrom.Name}}_{{$fk.ModelTo.Name}}
		{{end}}
	{{end}}

	joins []string
	originalWhere modelWhere
}

func(wr whereRelation{{$modelRelation.ModelFrom.Name}}_{{$modelRelation.ModelTo.Name}}) modelWhere() modelWhere {
	return wr.originalWhere
}

func(wr whereRelation{{$modelRelation.ModelFrom.Name}}_{{$modelRelation.ModelTo.Name}}) addCond(cond string) {
	wr.originalWhere.addJoin(wr.joins...)
	wr.originalWhere.addCond(cond)
}
func(wr whereRelation{{$modelRelation.ModelFrom.Name}}_{{$modelRelation.ModelTo.Name}}) addJoin(join ...string) {
	wr.originalWhere.addJoin(join...)
}
func(wr whereRelation{{$modelRelation.ModelFrom.Name}}_{{$modelRelation.ModelTo.Name}}) andOr() {
	wr.originalWhere.andOr()
}

{{range $relation := $modelRelation.ModelTo.DirectRelations}}
	{{if not (or (or (eq $relation.RelationType.String "ONE2ONE") (eq $relation.RelationType.String "ONE2MANY")) (eq $relation.ModelTo.Name $modelRelation.ModelFrom.Name) )}}
func(wr whereRelation{{$modelRelation.ModelFrom.Name}}_{{$modelRelation.ModelTo.Name}}) {{$relation.ModelTo.Name}}() whereRelation{{$modelRelation.ModelFrom.Name}}_{{$relation.ModelTo.Name}} {

	var newRelation whereRelation{{$modelRelation.ModelFrom.Name}}_{{$relation.ModelTo.Name}}
	newRelation.originalWhere = wr.originalWhere
	newRelation.joins = append(wr.joins, []string{"{{$modelRelation.ModelTo.Name}}=>{{$relation.ModelTo.Name}}"}...)

    {{range $field := $relation.ModelTo.Fields}}newRelation.{{$field.Name}}.where = newRelation
    newRelation.{{$field.Name}}.name = "\"{{$field.Model.TableName}}\".\"{{$field.TableName}}\""
    {{end}}
        

	{{$prefix := print "newRelation"}}
	{{$whereName := print "wr.originalWhere"}}
	{{$joins := print "\"" $modelRelation.ModelTo.Name "=>" $relation.ModelTo.Name "\"" }}
	{{$excludeModelName := print $modelRelation.ModelFrom.Name }}
	{{template "relationWhereInitialize" dict "Model" $relation.ModelTo "Prefix" $prefix "WhereName" $whereName "Joins" $joins "ExcludeModelName" $excludeModelName}}

	return newRelation
}
	{{end}}
{{end}}

{{end}}
{{end}}

{{end}}
{{define "modelWhere"}}
type {{lower .Name}}Where struct {
	{{range $field := .Fields}}{{$field.Name}}     whereField{{title $field.Type.Name}}{{$.Name}}
	{{end}}
	{{range $relation := .ForeignKeys}}{{$relation.Field.Name}} whereRelation{{$relation.ModelFrom.Name}}_{{$relation.ModelTo.Name}}
	{{end}}
	query          *{{lower .Name}}Query
	conds          string
	nextOperatorOr bool
}
{{range $relation := .DirectRelations}}
	{{if not (or (eq $relation.RelationType.String "ONE2ONE") (eq $relation.RelationType.String "ONE2MANY"))}}
func({{abbr $.Name}} *{{lower $.Name}}Where) {{$relation.ModelTo.Name}}() whereRelation{{$relation.ModelFrom.Name}}_{{$relation.ModelTo.Name}} {


	var newRelation whereRelation{{$relation.ModelFrom.Name}}_{{$relation.ModelTo.Name}}
	newRelation.originalWhere = {{abbr $.Name}}
	newRelation.joins = []string{"{{$relation.ModelFrom.Name}}=>{{$relation.ModelTo.Name}}"}

    {{range $field := $relation.ModelTo.Fields}}
    newRelation.{{$field.Name}}.where = &newRelation
    newRelation.{{$field.Name}}.name = "\"{{$field.Model.TableName}}\".\"{{$field.TableName}}\""
    {{end}}
        

	{{$prefix := print "newRelation"}}
	{{$whereName := print (abbr $.Name)}}
	{{$joins := print "\"" $relation.ModelFrom.Name "=>" $relation.ModelTo.Name "\"" }}
	{{$excludeModelName := print $relation.ModelFrom.Name }}
	{{template "relationWhereInitialize" dict "Model" $relation.ModelTo "Prefix" $prefix "WhereName" $whereName "Joins" $joins "ExcludeModelName" $excludeModelName}}

	return newRelation



	return newRelation
}
	{{end}}
{{end}}

func ({{abbr .Name}} {{lower .Name}}Where) conditionSQL() string {
	return {{abbr .Name}}.conds + ")"
}

func({{abbr .Name}} *{{lower .Name}}Where) modelWhere() modelWhere {
	return {{abbr .Name}}
}

func ({{abbr .Name}} *{{lower .Name}}Where) andOr() {
    if len({{abbr .Name}}.conds) > 0 {
        {{abbr .Name}}.conds += ") "
        if {{abbr .Name}}.nextOperatorOr {
            {{abbr .Name}}.conds += "OR "
        } else {
            {{abbr .Name}}.conds += "AND "
        }
    }
    {{abbr .Name}}.conds += "("
}
func ({{abbr .Name}} *{{lower .Name}}Where) AND() *{{lower .Name}}Where {
	{{abbr .Name}}.nextOperatorOr = false
    return {{abbr .Name}}
}
func ({{abbr .Name}} *{{lower .Name}}Where) OR() *{{lower .Name}}Where {
	{{abbr .Name}}.nextOperatorOr = true
    return {{abbr .Name}}
}

func ({{abbr .Name}} *{{lower .Name}}Where) addCond(cond string) {
	/*if tableNameEnd := strings.Index(cond, "\"."); tableNameEnd != -1 { // auto add join
		tableName := cond[len("\""):tableNameEnd]
		if tableName != "{{.TableName}}" {
			{{abbr .Name}}.addJoin(tableName)
		}
	}*/
    {{abbr .Name}}.conds += cond
}
func ({{abbr .Name}} *{{lower .Name}}Where) addJoin(joins ...string) {
	{{abbr .Name}}.query.joins = append({{abbr .Name}}.query.joins, joins...)
}
func ({{abbr .Name}} *{{lower .Name}}Where) ALL() []{{.Name}} {
    return {{abbr .Name}}.query.ALL()
}
func ({{abbr .Name}} *{{lower .Name}}Where) FIRST() *{{.Name}} {
    return {{abbr .Name}}.query.FIRST()
}
func ({{abbr .Name}} *{{lower .Name}}Where) LAST() *{{.Name}} {
    return {{abbr .Name}}.query.LAST()
}

func ({{abbr .Name}} *{{lower .Name}}Where) SQL() string {
    return {{abbr .Name}}.query.SQL()
}
func ({{abbr .Name}} *{{lower .Name}}Where) ORDERBY() *{{lower .Name}}OrderBy {
    return {{abbr .Name}}.query.ORDERBY()
}
func ({{abbr .Name}} *{{lower .Name}}Where) LIMIT(limit int) *{{lower .Name}}Query {
    return {{abbr .Name}}.query.LIMIT(limit)
}
{{end}}
`
