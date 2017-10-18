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
	Field     *modelField
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
	return tfr.Field.Model
}

func (tfr tableForeignRelation) ConstraintName() string {
	fieldFrom := tfr.Field
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

func (tfr tableForeignRelation) SqlJoin() ForeignRelationSlice {
	return []string{
		`INNER JOIN "` + tfr.ModelTo().TableName + `" ON "` + tfr.ModelFrom().TableName + `"."` + tfr.Field.TableName() + `" = "` + tfr.ModelTo().TableName + `"."` + tfr.ModelTo().PrimaryKey().TableName() + `"`,
	}

}

func (tfr tableForeignRelation) SqlCreateTable() string {

	return fmt.Sprintf("CONSTRAINT \"%s\" FOREIGN KEY (\"%s\") REFERENCES \"%s\" (\"%s\")",
		tfr.ConstraintName(),
		tfr.Field.TableName(),
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
	Name string

	IsPrimaryKey bool
	IsForeignKey bool

	IsPointer    bool
	IsGoBaseType bool
	IsArray      bool
	//Type fieldType.FieldType
	ConstType fieldType.ConstFieldType
	Relation  *tableForeignRelation

	//Type         fieldType.FieldType

	//ForeignKeys  []modelFieldFK
	Tags []modelFieldTag

	Model *ModelStruct

	fieldType fieldType.FieldType
}

func (mf modelField) TableName() string {
	tableName := snaker.CamelToSnake(mf.Name)
	if mf.IsForeignKey && mf.Relation != nil {
		tableName = strings.TrimSuffix(tableName, "s")
		return tableName + "_id"
	}
	return tableName
}

func (mf modelField) Type() fieldType.FieldType {
	if mf.fieldType == nil {
		mf.fieldType = mf.ConstType.BaseType(mf.IsPointer, mf.IsArray, mf.IsGoBaseType)
	}
	return mf.fieldType
}

//func (mf modelField) IsHiddenField() bool {
//	if mf.IsForeignKey && mf.Relation != nil && mf.Relation.IsManyRelation {
//		return true
//	}
//	return false
//}
func (mf modelField) FieldValueName(objName string) string {
	if mf.IsForeignKey {
		return mf.Relation.modelTo.PrimaryKey().FieldValueName(objName)
	}

	var fieldName = objName + "." + mf.Name
	if mf.ConstType == fieldType.DateType {
		return fieldName + ".Unix()"
	}

	if mf.IsPointer {
		fieldName = "*" + fieldName
	}

	return fieldName
}

func (mf modelField) DefaultValue() string {
	//if mf.IsForeignKey && mf.Type.ConstType == 0 {
	//	mf.Type = fieldType.Integer
	//}
	//fmt.Println(mf.Model.Name, mf.Name, "defaultValue", mf.Type)
	return mf.Type().DefaultValue()
}

func (mf modelField) FindTag(names ...string) (string, bool) {
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

	return mf.Type().SqlType()
}

func (mf modelField) ExecuteTemplate() string {
	//if mf.ConstType.Template == "" {
	//
	//	fmt.Println(mf)
	//	fmt.Println(mf.ConstType)
	//}
	var result = bytes.NewBufferString("")
	tmpl, parsing_err := template.New("modelField").Funcs(funcsMap).Parse(mf.Type().WhereTemplate())
	if parsing_err != nil {
		panic(parsing_err)
	}

	executing_err := tmpl.Execute(result, struct {
		ModelName      string
		ModelTableName string
		FieldAbbr      string
		FieldName      string
	}{mf.Model.Name, mf.Model.TableName, templateAbbrFunc(mf.Name), mf.Name})
	if executing_err != nil {
		panic(executing_err)
	}

	return result.String()
}

func (mf modelField) IsUnique() bool {
	for _, i := range mf.Model.indexes {
		if len(i.fieldNames) == 1 && i.fieldNames[0] == mf.Name {
			return true
		}
	}
	return false
}
