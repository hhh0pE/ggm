package main

import (
	"bytes"

	"text/template"

	"fmt"

	"strings"

	"github.com/hhh0pE/ggm/ggmgen/fieldType"
	"github.com/serenize/snaker"
)

type tableForeignRelation struct {
	field     *modelField
	modelTo   *ModelStruct
	isOne2One bool
	//TableToName    string
	//ModelToName    string
	//
	//modelTo   *ModelStruct
}

func (tfr tableForeignRelation) ModelTo() *ModelStruct {
	return tfr.modelTo
}
func (tfr tableForeignRelation) ModelFrom() *ModelStruct {
	return tfr.field.Model
}

func (tfr tableForeignRelation) ConstraintName() string {
	fieldFrom := tfr.field
	fieldTo := tfr.modelTo.PrimaryKey()
	//if fieldTo == nil {
	return fmt.Sprintf("%s_%s__%s_%s__fk",
		snaker.SnakeToCamel(fieldFrom.Model.TableName),
		fieldFrom.Name,
		snaker.SnakeToCamel(tfr.ModelTo().TableName),
		fieldTo.Name,
	)

	//return fieldFrom.Model.TableName + "__" + tfr.modelTo.TableName + "__fk"
	//}
}

func (tfr tableForeignRelation) SqlCreateTable() string {

	return fmt.Sprintf("CONSTRAINT \"%s\" FOREIGN KEY (\"%s\") REFERENCES \"%s\" (\"%s\")",
		tfr.ConstraintName(),
		tfr.field.TableName(),
		tfr.modelTo.TableName,
		tfr.modelTo.PrimaryKey().TableName(),
	)
	//return ""
}

func (tfr tableForeignRelation) SqlAlterTable() string {
	return fmt.Sprintf("ALTER TABLE \"%s\" ADD %s", tfr.ModelFrom().TableName, tfr.SqlCreateTable())
}

//func (tfr tableForeignRelation) IsManyToMany() bool {
//	if !tfr.IsManyRelation {
//		return false
//	}
//	if relationModel := pkgS.GetModel(tfr.ModelToName); relationModel != nil {
//		for _, f := range relationModel.fields {
//			if f.Relation != nil && f.Relation.IsManyRelation {
//				if relatedModel := f.Relation.ModelTo(); relationModel != nil {
//
//				}
//			}
//			//if f.Relation != nil && f.Relation.IsManyRelation && f.Relation.ModelToName ==
//		}
//	}
//}

type modelField struct {
	Name         string
	IsPointer    bool
	IsPrimaryKey bool
	IsForeignKey bool
	Relation     *tableForeignRelation
	Type         fieldType.FieldType

	//ForeignKeys  []modelFieldFK
	Tags []modelFieldTag

	Model *ModelStruct
}

func (mf modelField) TableName() string {
	tableName := snaker.CamelToSnake(mf.Name)
	if mf.IsForeignKey && mf.Relation != nil {
		tableName = strings.TrimSuffix(tableName, "s")
		return tableName + "_id"
	}
	return tableName
}

//func (mf modelField) IsHiddenField() bool {
//	if mf.IsForeignKey && mf.Relation != nil && mf.Relation.IsManyRelation {
//		return true
//	}
//	return false
//}
func (mf modelField) FieldValueName() string {
	if mf.Type.ConstType == fieldType.DateType {
		return mf.Name + ".Unix()"
	}
	if mf.IsForeignKey {
		return mf.Name + "." + mf.Relation.modelTo.PrimaryKey().FieldValueName()
	}
	return mf.Name
}

func (mf modelField) DefaultValue() string {
	//if mf.IsForeignKey && mf.Type.ConstType == 0 {
	//	mf.Type = fieldType.Integer
	//}
	//fmt.Println(mf.Model.Name, mf.Name, "defaultValue", mf.Type)
	return mf.Type.DefaultValue()
}

func (mf modelField) FindTag(names []string) (string, bool) {
	for _, name := range names {
		if val, exist := mf.GetTag(name); exist {
			return val, exist
		}
	}
	return "", false
}

func (mf modelField) GetTag(name string) (string, bool) {
	for _, t := range mf.Tags {
		if t.Name == name {
			return t.Value, true
		}
	}
	return "", false
}

func (mf modelField) SqlType() string {
	if mf.IsForeignKey {
		return "BIGINT"
	}
	if mf.IsPrimaryKey {
		return "SERIAL PRIMARY KEY"
	}
	var sqlType string
	switch mf.Type.ConstType {
	case fieldType.BoolType:
		sqlType = "BOOLEAN"
	case fieldType.IntType:
		sqlType = "BIGINT"
	case fieldType.FloatType:
		sqlType = "REAL"
	case fieldType.TextType:
		if mf.Type.MaxSize > 0 && mf.Type.MaxSize < 255 {
			sqlType = fmt.Sprintf("VARCHAR(%d)", mf.Type.MaxSize)
		} else {
			sqlType = "TEXT"
		}
	case fieldType.DateType:
		sqlType = "TIMESTAMP WITH TIME ZONE"
	default:
		return ""
	}

	if mf.Type.IsNullable {
		sqlType += " NULL"
	} else {
		sqlType += " NOT NULL"
	}
	return sqlType
	if mf.IsPrimaryKey {
		return "SERIAL"
	}
	return mf.Type.SqlType()
}

func (mf modelField) ExecuteTemplate() string {
	//if mf.ConstType.Template == "" {
	//
	//	fmt.Println(mf)
	//	fmt.Println(mf.ConstType)
	//}
	var result = bytes.NewBufferString("")
	tmpl, parsing_err := template.New("modelField").Funcs(funcsMap).Parse(mf.Type.Template())
	if parsing_err != nil {
		panic(parsing_err)
	}

	executing_err := tmpl.Execute(result, struct {
		ModelName string
		FieldAbbr string
		FieldName string
	}{mf.Model.Name, templateAbbrFunc(mf.Name), mf.Name})
	if executing_err != nil {
		panic(executing_err)
	}

	return result.String()
}