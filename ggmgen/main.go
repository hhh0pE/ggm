package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"go/importer"
	"sort"

	log "log"

	"time"

	"bytes"
	"io/ioutil"

	"github.com/hhh0pE/ggm/ggmgen/fieldType"
	"github.com/hhh0pE/ggm/ggmgen/templates"
)

func main() {
	if pStruct := analyze(); pStruct != nil {
		generatedContent := generate(*pStruct)

		var oldContent, newContent []byte

		path := pStruct.DirPath + "/ggm.go"
		var ggmFile *os.File

		if file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0666); err != nil {
			if os.IsNotExist(err) {
				if createdFile, creating_err := os.Create(path); creating_err != nil {
					panic("Can't create file " + path + ": " + creating_err.Error())
				} else {
					if _, writing_err := createdFile.Write(generatedContent.Bytes()); writing_err != nil {
						panic("can't write to file " + path + ": " + writing_err.Error())
					} else {
						exec.Command("goimports", "-w", path).Run()
						log.Println("Successfully created file " + path)
						return
					}
				}
			}
		} else {
			ggmFile = file
		}
		if tmpFile, err := ioutil.TempFile(os.TempDir(), "ggmTmpFile"); err != nil {
			panic("Can't create temp file :(")
		} else {
			_, write_err := tmpFile.Write(generatedContent.Bytes())
			tmpFile.Close()
			tmpFile, _ = os.Open(tmpFile.Name())
			if write_err != nil {
				fmt.Println("write err", write_err.Error())
			}
			goimportsOutput, goimports_err := exec.Command("goimports", "-w", tmpFile.Name()).CombinedOutput()
			if goimports_err != nil {
				fmt.Println("goimports error: " + goimports_err.Error())
			}

			log.Println("goimports output: " + string(goimportsOutput))
			var tmp_reading_err error
			newContent, tmp_reading_err = ioutil.ReadAll(tmpFile)
			if tmp_reading_err != nil {
				fmt.Println("tmp_reading_err", tmp_reading_err.Error())
			}

			os.Remove(tmpFile.Name())
		}

		var ggm_read_err error
		oldContent, ggm_read_err = ioutil.ReadAll(ggmFile)
		if ggm_read_err != nil {
			panic("can't read ggm file :(" + ggm_read_err.Error())
		}

		if len(oldContent) == len(newContent) {
			return
		}

		ggmFile.Truncate(0)
		ggmFile.Seek(0, 0)

		_, writing_err := ggmFile.Write(newContent)
		if writing_err != nil {
			panic("cannot write file " + path + ": " + writing_err.Error())
		} else {
			log.Println("successfully writed " + path)
		}
		//fmt.Println("generated content", len(generatedContent.Bytes()))
		//fmt.Println("oldContent", len(oldContent), ggm_read_err)
		//fmt.Println("newContent", len(newContent))
		//dmp := diffmatchpatch.New()
		//
		//diffs := dmp.DiffMain(string(oldContent), string(newContent), false)
		//patches := dmp.PatchMake(string(oldContent), string(newContent))
		//fmt.Println("patches count", len(patches))
		//for _, patch := range patches {
		//	fmt.Println(patch.String())
		//}

		//ggmFile, _ := os.Open(path)
		//oldContentReader := bufio.NewReader(ggmFile)
		//newContentReader := bufio.NewReader(bytes.NewReader(output))
		//
		//var oldContentLine []byte
		//var oldContentErr error
		//oldContentLine, _, oldContentErr = oldContentReader.ReadLine()
		//for oldContentErr == nil {
		//	oldContentLine, _, oldContentErr = oldContentReader.ReadLine()
		//}
		//
		//if ggmFileBytes, ok := ioutil.ReadFile(path); ok {
		//	newContent := string(output)
		//	oldContent := string(ggmFileBytes)
		//
		//}

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

	sort.Slice(pStruct.Models, func(i, j int) bool {
		return pStruct.Models[i].Name < pStruct.Models[j].Name
	})

	time4 := time.Now()
	// for _, m := range pStruct.Models {
	// m.Relations()

	// fmt.Println("model", m.Name)
	// for _, f := range m.AllFields() {
	// 	fmt.Println(f.Name, f.IsPointer, f.Type().Nullable())
	// }
	// fmt.Println()
	//if m.Name != "CurrencyPair" {
	//	continue
	//}

	// fmt.Println(m.Name, len(relations))
	// for _, dr := range relations {
	// 	if dr.ViaModel != nil {
	// 		fmt.Println("\t", dr.RelationType, dr.ModelFrom.Name+"<= "+dr.ViaModel.Name+" =>"+dr.ModelTo.Name)
	// 	} else {
	// 		fmt.Println("\t", dr.RelationType, dr.ModelFrom.Name+"<=>"+dr.ModelTo.Name)
	// 	}

	// }
	// fmt.Println()
	// }
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

func generate(ps packageStruct) *bytes.Buffer {
	var output = bytes.NewBufferString("")
	//ormFile, err := os.Create(ps.DirPath + "/ggm.go")
	//if err != nil {
	//	panic(err)
	//}
	//defer ormFile.Close()

	//fmt.Println(ps.DirPath + "/ggm.go")

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

	generalTmpl.Execute(output, ps)
	//tmpl.ExecuteTemplate(ormFile, "general.tmpl", ps)
	output.WriteString("\n\n")

	fieldTypes := fieldType.GetAllFieldTypes()
	for _, ft := range fieldTypes {
		tmpl, tmpl_err := template.New(ft.Name()).Funcs(funcsMap).Parse(ft.WhereTemplate())
		if tmpl_err != nil {
			panic(tmpl_err)
		}
		if execution_err := tmpl.Execute(output, nil); execution_err != nil {
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
		if executing_err := modelTmpl.Execute(output, m); executing_err != nil {
			panic(executing_err)
		}
		output.WriteString("\n\n")
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

	return output
}
