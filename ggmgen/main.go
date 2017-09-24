package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"go/importer"

	log "log"

	"github.com/hhh0pE/ggm/ggmgen/templates"
	"github.com/nullbio/inflect"
	"github.com/serenize/snaker"
)

func main() {
	if pStruct := analyze(); pStruct != nil {
		generate(*pStruct)
	}
}

var pkgS *packageStruct
var fset *token.FileSet

func analyze() *packageStruct {
	fset = token.NewFileSet()
	dir, _ := os.Getwd()

	pkgs, _ := parser.ParseDir(fset, dir, nil, parser.ParseComments)

	astf := make([]*ast.File, 0)

	for _, pkg := range pkgs {

		for fileName, pkgFile := range pkg.Files {
			if strings.HasSuffix(fileName, "ggm.go") {
				continue
			}
			astf = append(astf, pkgFile)
		}
	}

	conf := types.Config{Importer: importer.For("source", nil), IgnoreFuncBodies: true}
	conf.Error = func(err error) {
		log.Println("Error in package: " + err.Error())
	}

	var pkgTypesInfo *types.Package
	if typesInfo, checking_err := conf.Check(dir, fset, astf, nil); checking_err != nil {
		return nil
	} else {
		pkgTypesInfo = typesInfo
	}

	var pStruct packageStruct

	pkgS = &pStruct
	for _, pkg := range pkgs {
		var initPkgErr error
		pStruct, initPkgErr = initPackageStructs(pkg)
		if initPkgErr != nil {
			log.Println(initPkgErr.Error())
			return nil
		}
		break
	}

	for _, m := range pkgS.Models {
		if m.TableName == "" {
			m.TableName = generateTableName(m.Name)
		}
		//fmt.Println(m.Name, m.TableName)
	}

	pStruct.DirPath = dir

	for i, _ := range pStruct.Models {
		m := pStruct.Models[i]
		initStructFields(pkgTypesInfo, m)
	}

	for i := 0; i < len(pStruct.Models); i++ {
		if !pStruct.Models[i].HasPrimaryKey() {
			pStruct.Models = append(pStruct.Models[:i], pStruct.Models[i+1:]...)
			i--
		}
	}
	//
	//for _, m := range pStruct.Models {
	//
	//	if m.notify != nil {
	//		fmt.Println(m.Name, m.TableName)
	//		fmt.Println(m.notify)
	//		fmt.Println()
	//	}
	//
	//}

	return &pStruct
}

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
}

func generate(ps packageStruct) {
	ormFile, err := os.Create(ps.DirPath + "/ggm.go")
	if err != nil {
		panic(err)
	}
	defer ormFile.Close()

	fmt.Println(ps.DirPath + "/ggm.go")

	//var dirPath = "./templates"
	//var tmplFiles []string
	//filepath.Walk("./templates", func(path string, info os.FileInfo, err error) error {
	//	if info.IsDir() {
	//		return nil
	//	}
	//	tmplFiles = append(tmplFiles, path)
	//	return nil
	//})

	generalTmpl, generalTmpl_err := template.New("generalTmpl").Funcs(funcsMap).Parse(templates.GeneralTemplate)
	if generalTmpl_err != nil {
		panic(generalTmpl_err)
	}

	modelTmpl, modelTmpl_err := template.New("modelTmpl").Funcs(funcsMap).Parse(templates.ModelTemplate)
	if modelTmpl_err != nil {
		panic(modelTmpl_err)
	}

	//var templateParsing_err error
	//tmpl, templateParsing_err = template.New("ormTpl").Funcs(funcsMap).ParseFiles(tmplFiles...)
	//if templateParsing_err != nil {
	//	panic(templateParsing_err)
	//}

	generalTmpl.Execute(ormFile, ps)
	//tmpl.ExecuteTemplate(ormFile, "general.tmpl", ps)
	ormFile.WriteString("\n\n")

	//if fieldTypeFiles, err := ioutil.ReadDir("./templates/fields"); err != nil {
	//	panic(err)
	//} else {
	//	for _, ftf := range fieldTypeFiles {
	//		if exe_err := tmpl.ExecuteTemplate(ormFile, ftf.Name(), ps); exe_err != nil {
	//			panic(exe_err)
	//		}
	//		ormFile.WriteString("\n\n")
	//	}
	//}

	for _, m := range ps.Models {
		//fmt.Println("model!", m)
		//m.Fields
		if executing_err := modelTmpl.Execute(ormFile, m); executing_err != nil {
			panic(executing_err)
		}
		ormFile.WriteString("\n\n")
	}

	//fmt.Println(tmplFiles)
	//files, err := filepath.Glob("*.tmpl")
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(files)

	//
	//for _, file := range files {
	//	if file.IsDir() {
	//		ioutil.ReadDir(file.Name())
	//	}
	//	fmt.Println(file.Name())
	//}

	//mainTmpl := template.New("templates/Model.tmpl")
	//
	//executing_err := mainTmpl.Execute(ormFile, nil)
}
