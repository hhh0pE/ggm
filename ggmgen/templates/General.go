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
)

var ormDB *sql.DB

const DEFAULT_DB_PORT = 5432

var debug bool

func EnableDebug() {
	debug = true
}
func DisableDebug() {
	debug = false
}

func ConnectToDBAndInit(userName, dbName, password, host string, port int) error {
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
	return ormDB.Query(sql, args...)
}

func QueryRow(sql string, args ...interface{}) *sql.Row {
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
}

type model interface {
	tableName() string
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
func pgNotifyOnInsert(newModel model) {
	switch newModel.tableName() {
		{{- range $notify := .Notifies -}}
		{{- if $notify.OnInsert}}
		case "{{$notify.Model.TableName}}":
			for _, callback := range PgNotify.{{$notify.Model.Name}}.onInsert {
				callback(newModel.({{$notify.Model.Name}}))
			}
		{{- end -}}
		{{- end}}
		default:
			log.Println("PgNotify.OnInsert for model \""+newModel.tableName()+"\" that ggm not generated for. Run ggmgen!")
	}
}
func pgNotifyOnUpdate(oldModel model, newModel model) {
	switch newModel.tableName() {
		{{- range $notify := .Notifies -}}
		{{- if $notify.OnUpdate}}
		case "{{$notify.Model.TableName}}":
			for _, callback := range PgNotify.{{$notify.Model.Name}}.onUpdate {
				callback(oldModel.({{$notify.Model.Name}}), newModel.({{$notify.Model.Name}}))
			}
		{{- end -}}
		{{- end}}
		default:
			log.Println("PgNotify.OnUpdate for model \""+newModel.tableName()+"\" that ggm not generated for. Run ggmgen!")
	}
}
func pgNotifyOnDelete(oldModel model) {
	switch oldModel.tableName() {
		{{- range $notify := .Notifies}}
		{{- if $notify.OnDelete}}
		case "{{$notify.Model.TableName}}":
			for _, callback := range PgNotify.{{$notify.Model.Name}}.onDelete {
				callback(oldModel.({{$notify.Model.Name}}))
			}
		{{- end}}
		{{- end}}
		default:
			log.Println("PgNotify.OnDelete for model \""+oldModel.tableName()+"\" that ggm not generated for. Run ggmgen!")
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
`
