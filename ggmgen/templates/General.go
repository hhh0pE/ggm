package templates

const GeneralTemplate = `package {{.Name}}

import (
	"time"
	"strings"
	"fmt"
	"database/sql"
	"errors"
	"github.com/lib/pq"
)

var DB *sql.DB

type modelWhere interface{
    andOr()
    addCond(string)
}

type Model interface {
	TableName() string
}

type modelIndex struct {
	_modelData struct {
		isUnique bool
		model Model
		fieldNames []string
	}
}

func(mi *modelIndex) Index(isUnique bool, m Model, fieldNames ...string) {
	mi._modelData.isUnique = isUnique
	mi._modelData.model = m
	mi._modelData.fieldNames = fieldNames
}

type modelFK struct {
	_modelData struct {
		modelFrom Model
		fieldFromName string
		modelTo Model
		fieldToName string
	}
}

func(mfk *modelFK) ForeignKey(modelFrom Model, fieldFromName string, modelTo Model, fieldToName string) {
	mfk._modelData.modelFrom = modelFrom
	mfk._modelData.fieldFromName = fieldFromName
	mfk._modelData.modelTo = modelTo
	mfk._modelData.fieldToName = fieldToName
}

func RunMigration() error {
	if creating_table_err := createTableIfNotExist(); creating_table_err != nil {
		return errors.New("RunMigration() error: "+creating_table_err.Error())
	}
	if creating_columns_err := createTableColumnsIfNotExist(); creating_columns_err != nil {
		return errors.New("RunMigration() error: "+creating_columns_err.Error())
	}
	return nil
}

func createTableIfNotExist() error {
	var createTableSQL string
	{{range $model := .Models}}
	createTableSQL += ` + "`" + `CREATE TABLE IF NOT EXISTS "{{$model.TableName}}" (
	{{range $fi, $field := $model.Fields}}	"{{$field.TableName}}" {{$field.SqlType}}{{if not (IsLastElement $fi (len $model.Fields))}},{{end}}
	{{end}});
` + "`" + `{{end}}

	if _, err := DB.Exec(createTableSQL); err != nil {
		return err
	}
	return nil
}

func createTableColumnsIfNotExist() error {
	var alterTableAddColumnFunc = func(tableName, columnName, columnType string) error {
		_, err := DB.Query("SELECT \"" + columnName + "\" FROM \"" + tableName + "\" LIMIT 1;")
		if err != nil {
			if pqErr, canCast := err.(*pq.Error); canCast {
				if pqErr.Code.Name() == "undefined_column" {
					_, exec_err := DB.Exec("ALTER TABLE \"" + tableName + "\" ADD COLUMN \"" + columnName + " " + columnName + ";")
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

	{{range $model := .Models}}{{range $field := $model.Fields}}if creating_err := alterTableAddColumnFunc("{{$model.TableName}}", "{{$field.TableName}}", "{{$field.SqlType}}"); creating_err != nil {
			return creating_err
	}
	{{end}}{{end}}
	return nil
}
`
