package main

import (
	"strings"

	"github.com/hhh0pE/ggm/ggmgen/fieldType"
)

type modelIndex struct {
	name       string
	modelName  string
	fieldNames []string
	isUnique   bool
	isCoalesce bool
}

func (mi modelIndex) Name() string {
	if mi.name != "" {
		return mi.name
	}
	if mi.isUnique {
		return mi.modelName + `__` + strings.Join(mi.fieldNames, `_`) + `__ui`
	} else {
		return mi.modelName + `__` + strings.Join(mi.fieldNames, `_`) + `__i`
	}
}

func (mi modelIndex) Model() *ModelStruct {
	return pkgS.GetModel(mi.modelName)
}

func (mi modelIndex) Fields() []modelField {
	model := mi.Model()
	var fields []modelField
	for _, fieldName := range mi.fieldNames {
		if foundField := model.GetFieldByName(fieldName); foundField != nil {
			fields = append(fields, *foundField)
		}
	}

	return fields
}

//func (mi modelIndex) Check() error {
//	model := pkgS.GetModel(mi.modelName)
//	if model == nil {
//		return errors.New("Cannot find model with name \"" + mi.modelName + "\"")
//	}
//	return nil
//}

func (mi modelIndex) CreateIndexSQL() string {
	model := pkgS.GetModel(mi.modelName)
	if model == nil {
		return ""
	}
	var tableFieldNames []string
	for _, modelFieldName := range mi.fieldNames {
		if foundField := model.GetFieldByName(modelFieldName); foundField != nil {
			//fmt.Println(foundField.Name, foundField.Type.Name, foundField.Type.IsNullable)
			var fieldNameQuoted = foundField.TableName()
			if mi.isCoalesce && foundField.IsPointer {
				fieldNameQuoted = coalesce(*foundField)
				//fieldNameQuoted = `COALESCE("` + fieldNameQuoted + `", `
				//switch foundField.Type.ConstType {
				//case fieldType.IntType:
				//	fieldNameQuoted += `'0'`
				//case fieldType.FloatType:
				//	fieldNameQuoted += `'0.0'`
				//case fieldType.TextType:
				//	fieldNameQuoted += `''`
				//case fieldType.BoolType:
				//	fieldNameQuoted += `'FALSE'`
				//case fieldType.DateType:
				//	fieldNameQuoted += `'-infinity'`
				//}
				//fieldNameQuoted += `)`
			} else {
				//if strings.HasPrefix(fieldNameQuoted, `(`) && strings.HasSuffix(fieldNameQuoted, `)`) {
				//	fieldNameQuoted = fieldNameQuoted
				//} else {
				fieldNameQuoted = `"` + fieldNameQuoted + `"`
				//}
			}

			tableFieldNames = append(tableFieldNames, fieldNameQuoted)
		}
	}
	var isUnique string
	if mi.isUnique {
		isUnique = " UNIQUE"
	}

	return `CREATE` + isUnique + ` INDEX "` + mi.Name() + `" on "` + model.TableName + `" (` + strings.Join(tableFieldNames, `, `) + `);`
}

func coalesce(field modelField) string {
	fieldNameCoalesce := `COALESCE("` + field.TableName() + `", `
	switch field.ConstType {
	case fieldType.IntType:
		fieldNameCoalesce += `'0'`
	case fieldType.FloatType:
		fieldNameCoalesce += `'0.0'`
	case fieldType.TextType:
		fieldNameCoalesce += `''`
	case fieldType.BoolType:
		fieldNameCoalesce += `'FALSE'`
	case fieldType.DateType:
		fieldNameCoalesce += `'-infinity'`
	default:
		if field.IsForeignKey {
			fieldNameCoalesce += `'0'`
		}
	}
	fieldNameCoalesce += `)`

	return fieldNameCoalesce
}
