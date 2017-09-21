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

var ormDB *sql.DB

const DEFAULT_DB_PORT = 5432

func ConnectToDBAndInit(userName, dbName, password, host string, port int) error {
	if sqlConn, connecting_err := sql.Open("postgres", "user="+userName+" dbname="+dbName+" host="+host+" port="+fmt.Sprintf("%d", port)+" password="+password+" sslmode=disable"); connecting_err != nil {
		return errors.New("connectToDb error: "+connecting_err.Error())
	} else {
		ormDB = sqlConn
	}

	return RunMigration()
}
func SetDBAndInit(db *sql.DB) error {
	ormDB = db
	return RunMigration()
}

type modelWhere interface{
    andOr()
    addCond(string)
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
`
