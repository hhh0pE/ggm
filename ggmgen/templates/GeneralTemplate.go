package templates

const GeneralTemplate = `package {{.Name}}

import (
	"time"
	"strings"
	"fmt"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"log"
	"encoding/json"
	"github.com/hhh0pE/ggm"
)

var ormDB *sql.DB

const DEFAULT_DB_PORT = 5432

var debug bool
var debugSelect bool

func EnableDebug(debugSelect bool) {
	debug = true
	debugSelect = debugSelect
}
func DisableDebug() {
	debug = false
	debugSelect = false
}

func ConnectToDb(userName, dbName, password, host string, port int) {
	if connecting_err := connectToDBAndInit(userName, dbName, password, host, port); connecting_err != nil {
		panic(connecting_err)
	}
}

func connectToDBAndInit(userName, dbName, password, host string, port int) error {
	if sqlConn, connecting_err := sql.Open("postgres", "user="+userName+" dbname="+dbName+" host="+host+" port="+fmt.Sprintf("%d", port)+" password="+password+" sslmode=disable"); connecting_err != nil {
		return errors.New("connectToDb error: "+connecting_err.Error())
	} else {
		ormDB = sqlConn
	}

	if runMigration_err := RunMigration(); runMigration_err != nil {
		return errors.New("RunMigration error: "+runMigration_err.Error())
	}
	{{if .HasNotify}}
	if initPgListener_err := initPgListener("user="+userName+" dbname="+dbName+" host="+host+" port="+fmt.Sprintf("%d", port)+" password="+password+" sslmode=disable"); initPgListener_err != nil {
		return errors.New("initPgListener error: "+initPgListener_err.Error())
	}
	{{end}}
	return nil
}
/*func SetDBAndInit(db *sql.DB) error {
	ormDB = db
	return RunMigration()
}*/

func DB() *sql.DB {
	return ormDB
}

func SetMaxConnections(num int) {
	if ormDB == nil {
		panic("Cannot setMaxConnections for not initialized models.DB. Call SetDBAndInit first.")
	}
	ormDB.SetMaxOpenConns(num)
}

func Query(sql string, args ...interface{}) (*sql.Rows, error) {
	if debugSelect {
		log.Println(sql)
	}
	return ormDB.Query(sql, args...)
}

func QueryRow(sql string, args ...interface{}) *sql.Row {
	if debugSelect {
		log.Println(sql)
	}
	return ormDB.QueryRow(sql, args...)
}

func Exec(sql string, args ...interface{}) (sql.Result, error) {
	if debug {
		log.Println(sql)
	}
	return ormDB.Exec(sql, args...)
}

type modelWhere interface{
    andOr()
    addCond(string)
	addJoinChain(joinChain)
	modelWhere() modelWhere
}

type Model interface {
	TableName() string
}

type join struct {
	from *join
	TableName string
	FieldName string
	Alias string
	JoinType string
}

type joinMapType map[string]map[string]map[string]string
func(jmp joinMapType) sql(fromName, toName, fieldName, fromAlias, toAlias string) string {
	var sql string
	if sqlText, exist := jmp[fromName][toName][fieldName]; exist {
		sql = sqlText
	}
	if fromAlias != "" {
		sql = strings.Replace(sql, "\""+fromName+"\" ON", "\""+fromName+"\" "+fromAlias+" ON", 1)
		sql = strings.Replace(sql, "\""+fromName+"\".\"", fromAlias+".\"", -1)
	}
	if toAlias != "" {
		sql = strings.Replace(sql, "\""+toName+"\" ON", "\""+toName+"\" "+toAlias+" ON", 1)
		sql = strings.Replace(sql, "\""+toName+"\".\"", toAlias+".\"", -1)
	}
	return sql
}
var joinMap joinMapType

type joinChain struct{
	joins []join
}
func(jc *joinChain) AddJoin(tableName, fieldName, alias string) {
	jc.joins = append(jc.joins, join{TableName:tableName,FieldName:fieldName,Alias:alias})
	if len(jc.joins) > 1 {
		jc.joins[len(jc.joins)-1].from = &jc.joins[len(jc.joins)-2]
	}
}
func(jc joinChain) SQL(startTableName string) string {
	jc.joins[0].from = &join{TableName:startTableName}

	var result string
	for _, join := range jc.joins {
		result += joinMap.sql(join.from.TableName, join.TableName, join.FieldName, join.from.Alias, join.Alias)
	}
	return result
}

func init() {
	joinMap = make(map[string]map[string]map[string]string)
	{{range $model := .Models}}
		joinMap["{{$model.TableName}}"] = make(map[string]map[string]string)
		{{range $i, $relation := $model.DirectRelations}}
			
				joinMap["{{$relation.ModelFrom.TableName}}"]["{{$relation.ModelTo.TableName}}"] = make(map[string]string)
			
			{{if $relation.IsOneToXRelation}}
				joinMap["{{$relation.ModelFrom.TableName}}"]["{{$relation.ModelTo.TableName}}"][""] = ` + "`" + `{{($relation.SqlJoin "").SqlString}}` + "`" + `	
				{{range $fk := $relation.ModelFrom.ForeignKeysTo $relation.ModelTo}}
					joinMap["{{$relation.ModelFrom.TableName}}"]["{{$relation.ModelTo.TableName}}"]["{{$fk.Field.TableName}}"] = ` + "`" + `{{($relation.SqlJoin $fk.Field.TableName).SqlString}}` + "`" + `	
				{{end}}
			{{else}}
				joinMap["{{$relation.ModelFrom.TableName}}"]["{{$relation.ModelTo.TableName}}"][""] = ` + "`" + `{{($relation.SqlJoin "").SqlString}}` + "`" + `	
			{{end}}
		{{end}}
	{{end}}
}

func RunMigration() error {
	if creating_table_err := createTableIfNotExist(); creating_table_err != nil {
		return errors.New("RunMigration() createTableIfNotExist error: \n\t"+creating_table_err.Error())
	}
	if creating_columns_err := createTableColumnsIfNotExist(); creating_columns_err != nil {
		return errors.New("RunMigration() createTableColumnsIfNotExist error: \n\t"+creating_columns_err.Error())
	}
	if creating_fk_err := createForeignKeyIfNotExist(); creating_fk_err != nil {
		return errors.New("RunMigration() createForeignKeyIfNotExist error: \n\t"+creating_fk_err.Error())
	}
	if creating_indexes_err := createIndexes(); creating_indexes_err != nil {
		return errors.New("RunMigration() createIndexes error: \n\t"+creating_indexes_err.Error())
	}
	{{if .HasNotify -}}
	if creating_notifies_err := createNotifies(); creating_notifies_err != nil {
		return errors.New("RunMigration() create notifications error: \n\t"+creating_notifies_err.Error())
	}{{end}}
	return nil
}

func createTableIfNotExist() error {
	{{range $model := .Models}}
	if _, err := ormDB.Exec(` + "`" + `{{$model.CreateTableIfNotExistCommand}}` + "`" + `);  err != nil {
		return errors.New("CreateTableIfNotExist error for table \"{{$model.Name}}\": "+err.Error())
	}
	{{end}}
	return nil
}

func createTableColumnsIfNotExist() error {
	var alterTableAddColumnFunc = func(tableName, columnName, columnType string) error {
		_, err := ormDB.Query("SELECT \"" + columnName + "\" FROM \"" + tableName + "\" LIMIT 1;")
		if err != nil {
			if pqErr, canCast := err.(*pq.Error); canCast {
				if pqErr.Code.Name() == "undefined_column" {
					_, exec_err := ormDB.Exec("ALTER TABLE \"" + tableName + "\" ADD COLUMN \"" + columnName + "\" " + columnType + ";")
					if exec_err != nil {
						return errors.New("creatingColumnIfNotExist " + tableName + ".\"" + columnName + "\" error when adding new column: " + exec_err.Error())
					}
					return nil
				} else {
					return errors.New("creatingColumnIfNotExist " + tableName + ".\"" + columnName + "\" error: unexpected error type: " + pqErr.Error())
				}
			}
			return errors.New("creatingColumnIfNotExist " + tableName + ".\"" + columnName + "\" error: not postgres error: " + err.Error())
		}
		return nil
	}

	{{range $model := .Models}}{{range $field := $model.AllFields}}if creating_err := alterTableAddColumnFunc("{{$model.TableName}}", "{{$field.TableName}}", "{{$field.SqlType}}"); creating_err != nil {
			return creating_err
	}
	{{end}}{{end}}
	return nil
}

func createForeignKeyIfNotExist() error {
	var alterTableCreateFKFunc = func(sqlCommand string) error {
		_, exec_err := ormDB.Exec(sqlCommand)
		if exec_err != nil {
			if pqErr, ok := exec_err.(*pq.Error); ok {
				if pqErr.Code.Name() == "foreign_key_violation" {
					return errors.New("Cannot create foreign key constraint: there is a row already in DB that breaks this constraint. \n\tRaw SQL error: \"" + pqErr.Error() + "\"")
				}
				if pqErr.Code.Name() == "duplicate_object" {
					return nil
				}
				return pqErr
			}
		}
		return nil
	}
	{{range $model := .Models}}{{range $fk := $model.ForeignKeys}}if err := alterTableCreateFKFunc(` + "`" + `{{$fk.SqlAlterTable}}` + "`" + `); err != nil {
		return errors.New("createForeignKeyIfNotExist \"{{$fk.ConstraintName}}\" error: \n\t\t"+err.Error())
	}
	{{end}}{{end}}
	return nil
}

func createIndexes() error {
	var alterTableCreateIndexFunc = func(indexName, sqlCommand string) error {
		ormDB.Exec("DROP INDEX \""+indexName+"\";")
		_, exec_err := ormDB.Exec(sqlCommand)
		if exec_err != nil {
			if pqErr, ok := exec_err.(*pq.Error); ok {
				if pqErr.Code.Name() == "unique_violation" {
					return errors.New("Cannot create index: there is a row already in DB that breaks this constraint. \n\tRaw SQL error: \"" + pqErr.Error() + "\"")
				}
				if pqErr.Code.Name() == "duplicate_table" {
					return nil
				}
				return pqErr
			}
		}
		return nil
	}
	{{range $model := .Models}}{{range $index := $model.Indexes}}if err := alterTableCreateIndexFunc("{{$index.Name}}", ` + "`" + `{{$index.CreateIndexSQL}}` + "`" + `); err != nil {
		return errors.New("createIndexes \"{{$index.Name}}\" error: \n\t\t"+err.Error())
	}
	{{end}}{{end}}
	return nil
}

{{define "notifyFuncName"}}{{if .Name}}{{.Name}}{{else}}Notify_{{.ModelName}}{{end}}_ggm{{end}}
{{define "notifyChannelName"}}{{.ModelName}}{{if .OnInsert}}_insert{{end}}{{if .OnUpdate}}_update{{end}}{{if .OnDelete}}_delete{{end}}{{end}}
{{define "notifyChannelValueInsert" -}}
	{{- range $fi, $field := .Fields -}}
		to_json(NEW.{{$field.TableName}})::TEXT
		{{- if IsNotLastElement $fi (len $.Fields)}} || ',' || {{end -}}
	{{- end -}}
{{end}}
{{define "notifyChannelValueUpdate" -}}
	{{- range $fi, $field := .Fields -}}
		to_json(NEW.{{$field.TableName}})::TEXT
		{{- if IsNotLastElement $fi (len $.Fields)}} || ',' || {{end -}}
	{{- end -}}
	{{- if true}} || ';' || {{end -}}
	{{- range $fi, $field := .Fields -}}
		to_json(OLD.{{$field.TableName}})::TEXT
		{{- if IsNotLastElement $fi (len $.Fields)}} || ',' || {{end -}}
	{{- end -}}
{{end}}
{{define "notifyChannelValueDelete" -}}
	{{- range $fi, $field := .Fields -}}
		to_json(OLD.{{$field.TableName}})::TEXT
		{{- if IsNotLastElement $fi (len $.Fields)}} || ',' || {{end -}}
	{{- end -}}
{{end}}
{{define "notifyTriggerEvents" -}}
	{{- if .OnInsert -}}
		INSERT
	{{- end -}}
	{{- if .OnUpdate -}}
		{{- if .OnInsert}} OR {{end -}}
		UPDATE
	{{- end -}}
	{{- if .OnDelete -}}
		{{- if or .OnInsert .OnUpdate}} OR {{end -}}
		DELETE
	{{- end -}}
{{- end}}

func createNotifies() error {
	{{range $model := .Models -}}
		{{- if $model.Notify -}}
			if _, creatingFunc_err := ormDB.Exec(` + "`" + `CREATE OR REPLACE FUNCTION {{template "notifyFuncName" $model.Notify}}() RETURNS trigger AS $$
DECLARE
BEGIN
	{{- if $model.Notify.OnInsert}}
	IF TG_OP = 'INSERT' THEN
		PERFORM pg_notify('{{$model.Name}}_Insert', {{template "notifyChannelValueInsert" $model.Notify}});
	END IF;
	{{- end}}
	{{- if $model.Notify.OnUpdate}}
	IF TG_OP = 'UPDATE' THEN
		PERFORM pg_notify('{{$model.Name}}_Update', {{template "notifyChannelValueUpdate" $model.Notify}});
	END IF;
	{{- end}}
	{{- if $model.Notify.OnDelete}}
	IF TG_OP = 'DELETE' THEN
		PERFORM pg_notify('{{$model.Name}}_Delete', {{template "notifyChannelValueDelete" $model.Notify}});
	END IF;
	{{- end}}
  RETURN NEW;
END
$$ LANGUAGE plpgsql;` + "`" + `); creatingFunc_err != nil {
		return creatingFunc_err
	}

	if _, err := ormDB.Exec(` + "`" + `DROP TRIGGER IF EXISTS "{{template "notifyFuncName" $model.Notify}}" ON "{{$model.TableName}}";` + "`" + `); err != nil {
		return err
	}

	if _, err := ormDB.Exec(` + "`" + `CREATE TRIGGER "{{template "notifyFuncName" $model.Notify}}"
AFTER {{template "notifyTriggerEvents" $model.Notify}} ON "{{$model.TableName}}"
FOR EACH ROW
EXECUTE PROCEDURE {{template "notifyFuncName" $model.Notify}}();` + "`" + `); err != nil {
		return err
	}

		{{end -}}
	{{- end -}}


	return nil
}

{{range $notify := .Notifies}}
type notify{{$notify.Model.Name}} struct {
	{{if $notify.OnInsert}}onInsert []func({{$notify.Model.Name}}){{end}}
	{{if $notify.OnUpdate}}onUpdate []func({{$notify.Model.Name}}, {{$notify.Model.Name}}){{end}}
	{{if $notify.OnDelete}}onDelete []func({{$notify.Model.Name}}){{end}}
}
{{if $notify.OnInsert -}}
func (ne *notify{{$notify.Model.Name}}) OnInsert(callback func({{$notify.Model.Name}})) {
	ne.onInsert = append(ne.onInsert, callback)
}
{{- end}}
{{if $notify.OnUpdate -}}
func (ne *notify{{$notify.Model.Name}}) OnUpdate(callback func({{$notify.Model.Name}}, {{$notify.Model.Name}})) {
	ne.onUpdate = append(ne.onUpdate, callback)
}
{{- end}}
{{if $notify.OnDelete -}}
func (ne *notify{{$notify.Model.Name}}) OnDelete(callback func({{$notify.Model.Name}})) {
	ne.onDelete = append(ne.onDelete, callback)
}
{{- end}}
{{- end -}}

{{if .HasNotify}}
var PgNotify pgNotify
type pgNotify struct {
	{{- range $notify := .Notifies}}
	{{$notify.Model.Name}} notify{{$notify.Model.Name}}
	{{- end}}
}
func pgNotifyOnInsert(newModel Model) {
	switch newModel.TableName() {
		{{- range $notify := .Notifies -}}
		{{- if $notify.OnInsert}}
		case "{{$notify.Model.TableName}}":
			for _, callback := range PgNotify.{{$notify.Model.Name}}.onInsert {
				callback(newModel.({{$notify.Model.Name}}))
			}
		{{- end -}}
		{{- end}}
		default:
			log.Println("PgNotify.OnInsert for model \""+newModel.TableName()+"\" that ggm not generated for. Run ggmgen!")
	}
}
func pgNotifyOnUpdate(oldModel Model, newModel Model) {
	switch newModel.TableName() {
		{{- range $notify := .Notifies -}}
		{{- if $notify.OnUpdate}}
		case "{{$notify.Model.TableName}}":
			for _, callback := range PgNotify.{{$notify.Model.Name}}.onUpdate {
				callback(oldModel.({{$notify.Model.Name}}), newModel.({{$notify.Model.Name}}))
			}
		{{- end -}}
		{{- end}}
		default:
			log.Println("PgNotify.OnUpdate for model \""+newModel.TableName()+"\" that ggm not generated for. Run ggmgen!")
	}
}
func pgNotifyOnDelete(oldModel Model) {
	switch oldModel.TableName() {
		{{- range $notify := .Notifies}}
		{{- if $notify.OnDelete}}
		case "{{$notify.Model.TableName}}":
			for _, callback := range PgNotify.{{$notify.Model.Name}}.onDelete {
				callback(oldModel.({{$notify.Model.Name}}))
			}
		{{- end}}
		{{- end}}
		default:
			log.Println("PgNotify.OnDelete for model \""+oldModel.TableName()+"\" that ggm not generated for. Run ggmgen!")
	}
}

var pqListener *pq.Listener
func initPgListener(name string) error {
	pqListener = pq.NewListener(name, 10*time.Second, time.Minute, func(ev pq.ListenerEventType, err error) {
		if err != nil {
			panic("Error when initializing pq Listener (for NOTIFY/LISTEN): "+err.Error())
		}
	})

	{{range $notify := .Notifies -}}
		{{- if $notify.OnInsert}}
	pqListener.Listen("{{$notify.Model.Name}}_Insert")
		{{- end -}}
		{{- if $notify.OnUpdate}}
	pqListener.Listen("{{$notify.Model.Name}}_Update")
		{{- end -}}
		{{- if $notify.OnDelete}}
	pqListener.Listen("{{$notify.Model.Name}}_Delete")
		{{- end -}}
	{{- end}}

	go func() {
		for {
			select {
				case notify := <-pqListener.Notify:
					{{range $notify := .Notifies}}
						var convertPayloadToModel{{$notify.Model.Name}} = func(payload string) {{$notify.Model.Name}} {
							strFields := strings.Split(payload, ",")
							var new{{$notify.Model.Name}} {{$notify.Model.Name}}
							{{range $fi, $field := $notify.Fields}}
								if decoding_err := json.Unmarshal([]byte(strFields[{{$fi}}]), &new{{$notify.Model.Name}}.{{$field.Name}}); decoding_err != nil {
									log.Println("error when unmarshalling field: "+decoding_err.Error())
								}
							{{end}}
							return new{{$notify.Model.Name}}
						}
						{{if $notify.OnInsert -}}
						if notify.Channel == "{{$notify.Model.Name}}_Insert" {
							pgNotifyOnInsert(convertPayloadToModel{{$notify.Model.Name}}(notify.Extra))
						}
						{{- end -}}
						{{- if $notify.OnUpdate}}
						if notify.Channel == "{{$notify.Model.Name}}_Update" {
							parts := strings.Split(notify.Extra, ";")
							pgNotifyOnUpdate(convertPayloadToModel{{$notify.Model.Name}}(parts[0]), convertPayloadToModel{{$notify.Model.Name}}(parts[1]))
						}
						{{- end -}}
						{{- if $notify.OnDelete}}
						if notify.Channel == "{{$notify.Model.Name}}_Delete" {
							pgNotifyOnDelete(convertPayloadToModel{{$notify.Model.Name}}(notify.Extra))
						}
						{{- end -}}

					{{end}}
			}
		}
	}()
	return nil
}
{{end}}


func boolArrayToSqlValue(arr []bool) string {
	sqlVal, _ := pq.BoolArray(arr).Value()
	return sqlVal.(string)
}
func int64ArrayToSqlValue(arr []int64) string {
	sqlVal, _ := pq.Int64Array(arr).Value()
	return sqlVal.(string)
}
func float64ArrayToSqlValue(arr []float64) string {
	sqlVal, _ := pq.Float64Array(arr).Value()
	return sqlVal.(string)
}
func stringArrayToSqlValue(arr []string) string {
	sqlVal, _ := pq.StringArray(arr).Value()
	return sqlVal.(string)
}
func byteArrayToSqlValue(arr [][]byte) string {
	sqlVal, _ := pq.ByteaArray(arr).Value()
	return sqlVal.(string)
}
func decimalArrayToSqlValue(arr []ggm.Decimal) string {
	var values []string
	for _, dec := range arr {
		values = append(values, dec.String())
	}
	return stringArrayToSqlValue(values)
}


func scannerTypeToBaseType(s interface{}, baseType interface{}) interface{} {
	switch typedVal := s.(type) {
	case pq.NullTime:
		if _, ok := baseType.(*time.Time); ok {
			if !typedVal.Valid {
				return nil
			}
			return &typedVal.Time
		}
		if _, ok := baseType.(pq.NullTime); ok {
			return typedVal
		}
	case sql.NullBool:
		if _, ok := baseType.(*bool); ok {
			if !typedVal.Valid {
				return nil
			}
			return &typedVal.Bool
		}
		if _, ok := baseType.(sql.NullBool); ok {
			return typedVal
		}
	case sql.NullString:
		if _, ok := baseType.(*string); ok {
			if !typedVal.Valid {
				return nil
			}
			return &typedVal.String
		}
		if _, ok := baseType.(sql.NullString); ok {
			return typedVal
		}
	case sql.NullInt64:
		if _, ok := baseType.(*int64); ok {
			if !typedVal.Valid {
				return nil
			}
			return &typedVal.Int64
		}
		if _, ok := baseType.(sql.NullInt64); ok {
			return typedVal
		}
	case sql.NullFloat64:
		if _, ok := baseType.(*float64); ok {
			if !typedVal.Valid {
				return nil
			}
			return &typedVal.Float64
		}
		if _, ok := baseType.(sql.NullFloat64); ok {
			return typedVal
		}
	case pq.StringArray:
		if _, ok := baseType.([]string); ok {
			return []string(typedVal)
		}
		if _, ok := baseType.(pq.StringArray); ok {
			return typedVal
		}
	case *pq.StringArray:
		if _, ok := baseType.(*[]string); ok {
			if typedVal == nil {
				return nil
			}
			val := []string(*typedVal)
			return &val
		}
		if _, ok := baseType.(*pq.StringArray); ok {
			return typedVal
		}
	case pq.BoolArray:
		if _, ok := baseType.([]bool); ok {
			return []bool(typedVal)
		}
		if _, ok := baseType.(pq.BoolArray); ok {
			return typedVal
		}
	case *pq.BoolArray:
		if _, ok := baseType.(*[]bool); ok {
			if typedVal == nil {
				return nil
			}
			val := []bool(*typedVal)
			return &val
		}
		if _, ok := baseType.(*pq.BoolArray); ok {
			return typedVal
		}
	//case pq.ByteaArray:
	//	if _, ok := baseType.([][]byte); ok {
	//		baseType = [][]byte(typedVal)
	//	}
	case pq.Float64Array:
		if _, ok := baseType.([]float64); ok {
			return []float64(typedVal)
		}
		if _, ok := baseType.(*pq.Float64Array); ok {
			return typedVal
		}
	case *pq.Float64Array:
		if _, ok := baseType.(*[]float64); ok {
			if typedVal == nil {
				return nil
			}
			val := []float64(*typedVal)
			return &val
		}
		if _, ok := baseType.(*pq.Float64Array); ok {
			return typedVal
		}
	case pq.Int64Array:
		if _, ok := baseType.([]int64); ok {
			return []int64(typedVal)
		}
	case *pq.Int64Array:
		if _, ok := baseType.(*[]int64); ok {
			if typedVal == nil {
				return nil
			}
			val := []int64(*typedVal)
			return &val
		}
		if _, ok := baseType.(*pq.Int64Array); ok {
			return typedVal
		}
	case ggm.Decimal:
		if _, ok := baseType.(float64); ok {
			f64, _ := typedVal.Float64()
			return f64
		}
		if _, ok := baseType.(ggm.Decimal); ok {
			return typedVal
		}
	case *ggm.Decimal:
		if _, ok := baseType.(*float64); ok {
			f64, _ := typedVal.Float64()
			return &f64
		}
		if _, ok := baseType.(*ggm.Decimal); ok {
			return typedVal
		}
	default:
		return errors.New(fmt.Sprintf("Cannot cast Scanner type to basic go type : %#v ", baseType))
	}

	// never catch that
	return nil
}

func IsEmpty(m Model) bool {
	switch m.TableName() {
	{{range $model := .Models}}
		case "{{$model.TableName}}":
		if casted, ok := m.(*{{$model.Name}}); ok {
			return isEmpty{{$model.Name}}(casted)
		}
		if casted, ok := m.({{$model.Name}}); ok {
			return isEmpty{{$model.Name}}(&casted)
		}
	{{end}}
	}
	log.Println("Cannot run isEmpty for model \""+m.TableName()+"\" - run ggmgen first!")

	return false
}

func Save(m Model) error {
	switch m.TableName() {
	{{range $model := .Models}}
		case "{{$model.TableName}}":
		if casted, ok := m.(*{{$model.Name}}); ok {
			return save{{$model.Name}}(casted)
		} else {
			return errors.New("You can only save *{{$model.Name}} object")
		}
	{{end}}
	}

	return errors.New("Cannot save model \""+m.TableName()+"\" - run ggmgen first!")
}

func Update(m Model) error {
	switch m.TableName() {
	{{range $model := .Models}}
		case "{{$model.TableName}}":
		if casted, ok := m.(*{{$model.Name}}); ok {
			return update{{$model.Name}}(casted)
		}
	{{end}}
	}

	return errors.New("Cannot update model \""+m.TableName()+"\" - run ggmgen first!")
}

func Insert(m Model) error {
	switch m.TableName() {
	{{range $model := .Models}}
		case "{{$model.TableName}}":
		if casted, ok := m.(*{{$model.Name}}); ok {
			return insert{{$model.Name}}(casted)
		}
	{{end}}
	}

	return errors.New("Cannot insert model \""+m.TableName()+"\" - run ggmgen first!")
}

func Delete(m Model) error {
	switch m.TableName() {
	{{range $model := .Models}}
		case "{{$model.TableName}}":
		if casted, ok := m.(*{{$model.Name}}); ok {
			return delete{{$model.Name}}(casted)
		}
		if casted, ok := m.({{$model.Name}}); ok {
			return delete{{$model.Name}}(&casted)
		}
	{{end}}
	}

	return errors.New("Cannot delete model \""+m.TableName()+"\" - run ggmgen first!")
}
`
