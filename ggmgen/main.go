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

	"go/importer"

	log "log"

	"time"

	"github.com/hhh0pE/ggm/ggmgen/fieldType"
	"github.com/hhh0pE/ggm/ggmgen/templates"
)

func main() {
	if pStruct := analyze(); pStruct != nil {
		generate(*pStruct)
	}
}

var pkgS *packageStruct
var fset *token.FileSet

func analyze() *packageStruct {

	time1 := time.Now()

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

	time2 := time.Now()
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

	for i, _ := range pkgS.Models {
		if pkgS.Models[i].TableName == "" {
			pkgS.Models[i].TableName = generateTableName(pkgS.Models[i].Name)
		}

	}

	pStruct.DirPath = dir

	time3 := time.Now()
	for i, _ := range pStruct.Models {
		m := pStruct.Models[i]
		initStructFields(pkgTypesInfo, m)
	}

	for i := 0; i < len(pStruct.Models); i++ {
		for fi, _ := range pStruct.Models[i].fields {
			pStruct.Models[i].fields[fi].Model = pStruct.Models[i]
		}
		if !pStruct.Models[i].HasPrimaryKey() {
			pStruct.Models = append(pStruct.Models[:i], pStruct.Models[i+1:]...)
			i--
		}
	}

	time4 := time.Now()
	for _, m := range pStruct.Models {
		relations := m.Relations()
		//if m.Name != "CurrencyPair" {
		//	continue
		//}

		fmt.Println(m.Name, len(relations))
		for _, dr := range relations {
			if dr.ViaModel != nil {
				fmt.Println("\t", dr.RelationType, dr.ModelFrom.Name+"<= "+dr.ViaModel.Name+" =>"+dr.ModelTo.Name)
			} else {
				fmt.Println("\t", dr.RelationType, dr.ModelFrom.Name+"<=>"+dr.ModelTo.Name)
			}

		}
		fmt.Println()
	}
	fmt.Println("building relations execution time: ", time.Now().Sub(time4))
	fmt.Println("initPackageStruct time", time3.Sub(time2))
	fmt.Println("initStructFields time", time4.Sub(time3))
	fmt.Println("Total time: ", time.Now().Sub(time1))
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

	fieldTypes := fieldType.GetAllFieldTypes()
	for _, ft := range fieldTypes {
		tmpl, tmpl_err := template.New(ft.Name()).Funcs(funcsMap).Parse(ft.WhereTemplate())
		if tmpl_err != nil {
			panic(tmpl_err)
		}
		if execution_err := tmpl.Execute(ormFile, nil); execution_err != nil {
			panic(execution_err)
		}
	}

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
