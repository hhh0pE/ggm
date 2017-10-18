package main

import (
	"bytes"
	"errors"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"github.com/hhh0pE/ggm/ggmgen/fieldType"
	"github.com/nullbio/inflect"
	"github.com/serenize/snaker"
)

func templateAbbrFunc(name string) string {
	var abbr string
	for _, symb := range name {
		if unicode.IsUpper(symb) {
			abbr += strings.ToLower(string(symb))
		}
	}
	return abbr
}

func generateTableName(modelName string) string {
	snakeCase := snaker.CamelToSnake(modelName)
	return inflect.Pluralize(snakeCase)
}

func templateLowerFunc(name string) string {
	firstSymb, firstSymbSize := utf8.DecodeRuneInString(name)
	return strings.ToLower(string(firstSymb)) + name[firstSymbSize:]
}

func executeFieldTypeTemplate(ft fieldType.FieldType, model ModelStruct) string {
	var result = bytes.NewBufferString("")
	tmpl, parsing_err := template.New("fieldTypeTemplate").Funcs(baseFuncMap).Parse(ft.WhereTemplate())
	if parsing_err != nil {
		panic(parsing_err)
	}

	executing_err := tmpl.Execute(result, struct {
		ModelName      string
		ModelTableName string
		FieldAbbr      string
		FieldName      string
	}{model.Name, model.TableName, templateAbbrFunc(model.Name), model.Name})
	if executing_err != nil {
		panic(executing_err)
	}

	return result.String()
}

func dict(values ...interface{}) (map[string]interface{}, error) {
		if len(values)%2 != 0 {
			return nil, errors.New("invalid dict call")
		}
		dict := make(map[string]interface{}, len(values)/2)
		for i := 0; i < len(values); i += 2 {
			key, ok := values[i].(string)
			if !ok {
				return nil, errors.New("dict keys must be strings")
			}
			dict[key] = values[i+1]
		}
		return dict, nil
}

func allFiledTypes() []fieldType.FieldType {
	return fieldType.GetAllFieldTypes()
}

var baseFuncMap = template.FuncMap{
	"lower": templateLowerFunc,
	"title": func(str string) string {
		return strings.Title(str)
	},
	"abbr": templateAbbrFunc,
	"IsLastElement": func(index, elementCounts int) bool {
		if index == elementCounts-1 {
			return true
		}
		return false
	},
	"IsNotLastElement": func(index, elementCounts int) bool {
		if index == elementCounts-1 {
			return false
		}
		return true
	},
	"dict":dict,
}

var funcsMap = template.FuncMap{
	"lower": templateLowerFunc,
	"title": func(str string) string {
		return strings.Title(str)
	},
	"abbr": templateAbbrFunc,
	"IsLastElement": func(index, elementCounts int) bool {
		if index == elementCounts-1 {
			return true
		}
		return false
	},
	"IsNotLastElement": func(index, elementCounts int) bool {
		if index == elementCounts-1 {
			return false
		}
		return true
	},
	"dict":dict,
	"ExecuteFieldTypeTemplate": executeFieldTypeTemplate,
	"GetAllFieldTypes":         allFiledTypes,
}
