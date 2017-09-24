package main

import (
	"errors"
	"go/ast"
	"strconv"
	"strings"
)

func initPackageStructs(pkg *ast.Package) (packageStruct, error) {
	var ps packageStruct
	ps.Name = pkg.Name

	var indexes []modelIndex
	var pgNotifies []pgNotify

	for fileName, pkgFile := range pkg.Files {
		if strings.HasSuffix(fileName, "ggm.go") {
			continue
		}

		var fkImportName string
		for _, im := range pkgFile.Imports {
			if unquotedPath, err := strconv.Unquote(im.Path.Value); err == nil {
				if unquotedPath == "github.com/hhh0pE/ggm/primaryKey" {
					if im.Name == nil {
						fkImportName = "primaryKey"
						break
					}
					if im.Name.Name == "" {
						fkImportName = "primaryKey"
					} else {
						fkImportName = im.Name.Name
					}

					break
				}
			}
		}
		for _, d := range pkgFile.Decls {
			if genDecl, ok := d.(*ast.GenDecl); ok {
				for _, s := range genDecl.Specs {
					if typeS, typeSOk := s.(*ast.TypeSpec); typeSOk {
						if structInfo, isBased := typeS.Type.(*ast.StructType); isBased {
							var newModel ModelStruct
							newModel.Name = typeS.Name.Name

							if fkImportName != "" {
								for _, f := range structInfo.Fields.List {
									if sel, ok := f.Type.(*ast.SelectorExpr); ok {
										//fmt.Println(sel.Sel.Name, sel.X, fkImportName)
										if xIdent, iok := sel.X.(*ast.Ident); iok {
											if xIdent.Name == fkImportName {
												for _, fieldName := range f.Names {
													var newMF modelField
													newMF.Name = fieldName.Name
													newMF.IsPrimaryKey = true

													newModel.AddField(newMF)
												}
											}

										}

									}
								}
							}

							ps.AddModel(newModel)
						}
					}
				}
			}
			if fDecl, ok := d.(*ast.FuncDecl); ok {
				modelAbbr := fDecl.Recv.List[0].Names[0].Name
				var modelName string

				recvType := fDecl.Recv.List[0].Type
				if starIdent, starOk := recvType.(*ast.StarExpr); starOk {
					recvType = starIdent.X
				}

				if modelIdent, ok := recvType.(*ast.Ident); ok {
					modelName = modelIdent.Name
				}

				//fmt.Println(modelAbbr, modelName)
				if fDecl.Name.Name == "PgNotify" {
					for _, st := range fDecl.Body.List {
						if expr, ok := st.(*ast.ExprStmt); ok {
							if callExpr, ok := expr.X.(*ast.CallExpr); ok {
								var newPgNotify pgNotify
								if err := pgNotifyCallStackRecursive(modelAbbr, callExpr, &newPgNotify); err != nil {
									return ps, err
								}

								newPgNotify.ModelName = modelName

								pgNotifies = append(pgNotifies, newPgNotify)
							}
						}
					}
				}
				if fDecl.Name.Name == "Indexes" {
					for _, st := range fDecl.Body.List {
						if expr, ok := st.(*ast.ExprStmt); ok {
							if callExpr, ok := expr.X.(*ast.CallExpr); ok {
								var newModelIndex modelIndex

								if typeIdent, ok := fDecl.Recv.List[0].Type.(*ast.Ident); ok && len(fDecl.Recv.List) > 0 {
									newModelIndex.modelName = typeIdent.Name
								}
								//fmt.Println()
								indexErr := indexCallStackRecursive(modelAbbr, callExpr, &newModelIndex)
								if indexErr != nil {
									return ps, errors.New("\"" + modelName + "\".Indexes() error: " + indexErr.Error())
								}

								indexes = append(indexes, newModelIndex)
							}
						}
					}
				}
				if fDecl.Name.Name == "TableName" {
					var ms ModelStruct

					for _, st := range fDecl.Body.List {
						if rs, ok := st.(*ast.ReturnStmt); ok {
							for _, re := range rs.Results {
								if bl, ok := re.(*ast.BasicLit); ok {
									ms.TableName = strings.Trim(bl.Value, `"`)
								}
							}
						}
					}
					if fDecl.Recv != nil && fDecl.Recv.List != nil {
						for _, l := range fDecl.Recv.List {
							var ident *ast.Ident
							if i, ok := l.Type.(*ast.Ident); ok {
								ident = i
							}
							if starExpr, ok := l.Type.(*ast.StarExpr); ok {
								if i, ok := starExpr.X.(*ast.Ident); ok {
									ident = i
								}
							}

							ms.Name = ident.Name
						}
					}

					ps.AddModel(ms)
				}
			}
		}
	}

	for ii, index := range indexes {
		if foundModel := ps.GetModel(index.modelName); foundModel != nil {
			foundModel.indexes = append(foundModel.indexes, indexes[ii])
		}
	}

	for pni, notify := range pgNotifies {
		if foundModel := ps.GetModel(notify.ModelName); foundModel != nil {
			foundModel.notify = &pgNotifies[pni]
			foundModel.notify.Model = foundModel
		}
	}

	return ps, nil
}

func pgNotifyCallStackRecursive(modelAbbr string, callExpr *ast.CallExpr, mpn *pgNotify) error {
	if callExprSelector, ok := callExpr.Fun.(*ast.SelectorExpr); ok {

		if callExprSelector.Sel.Name == "PgNotify" {
			for _, arg := range callExpr.Args {
				//fmt.Println("PgNotify arg", arg, reflect.TypeOf(arg))
				if selectorExpr, ok := arg.(*ast.SelectorExpr); ok {
					switch selectorExpr.Sel.Name {
					case "PG_NOTIFY_INSERT":
						mpn.OnInsert = true
					case "PG_NOTIFY_UPDATE":
						mpn.OnUpdate = true
					case "PG_NOTIFY_DELETE":
						mpn.OnDelete = true
					}
					//if selectorExpr.Sel.Name == "PG_NOTIFY_INSERT" {
					//	mpn.OnInsert = true
					//}
					//fmt.Println(selectorExpr.Sel.Name)
					//if ident, identOk := selectorExpr.X.(*ast.Ident); identOk {
					//	fmt.Println(ident.Name, ident.Obj)
					//}

					//fmt.Println(selectorExpr.X, selectorExpr.Sel.Name)
				}
				//fmt.Println(arg, reflect.TypeOf(arg))
			}
		}
		if callExprSelector.Sel.Name == "Name" {
			for _, ar := range callExpr.Args {
				if arBL, ok := ar.(*ast.BasicLit); ok {
					notifyName, _ := strconv.Unquote(arBL.Value)
					mpn.Name = notifyName
				}
			}
			//fmt.Println(callExprSelector.X, reflect.TypeOf(callExprSelector.X))
		}
		if callExprSelector.Sel.Name == "Payload" {
			for _, arg := range callExpr.Args {
				if selectorExpr, ok := arg.(*ast.SelectorExpr); ok {
					if ident, ok := selectorExpr.X.(*ast.Ident); !ok || ident.Name != modelAbbr {
						return errors.New("Cannot pass to PgNotify not model fields!")
					}

					//fmt.Println(selectorExpr.X, selectorExpr.Sel.Name)
					mpn.fieldNames = append(mpn.fieldNames, selectorExpr.Sel.Name)
				}
				//fmt.Println(arg, reflect.TypeOf(arg))
			}
		}
		if callExpr2, ok := callExprSelector.X.(*ast.CallExpr); ok {
			return pgNotifyCallStackRecursive(modelAbbr, callExpr2, mpn)
			//fmt.Println("args2", callExpr2.Args)
		}
	}
	return nil
}

func indexCallStackRecursive(modelAbbr string, callExpr *ast.CallExpr, mi *modelIndex) error {
	if callExprSelector, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
		//buff := bytes.NewBufferString("")
		//if printing_err := printer.Fprint(buff, fset, callExprSelector); printing_err != nil {
		//	log.Fatal(printing_err)
		//} else {
		//	fmt.Println(buff)
		//}
		if callExprSelector.Sel.Name == "Unique" {
			mi.isUnique = true
		}
		if callExprSelector.Sel.Name == "Coalesce" {
			mi.isCoalesce = true
		}
		if callExprSelector.Sel.Name == "Name" {
			for _, ar := range callExpr.Args {
				if arBL, ok := ar.(*ast.BasicLit); ok {
					indexName, _ := strconv.Unquote(arBL.Value)
					mi.name = indexName
				}
			}
			//fmt.Println(callExprSelector.X, reflect.TypeOf(callExprSelector.X))
		}
		if callExprSelector.Sel.Name == "Index" {
			for _, arg := range callExpr.Args {
				if selectorExpr, ok := arg.(*ast.SelectorExpr); ok {
					if ident, ok := selectorExpr.X.(*ast.Ident); !ok || ident.Name != modelAbbr {
						return errors.New("Cannot pass to Index not model fields!")
					}

					//fmt.Println(selectorExpr.X, selectorExpr.Sel.Name)
					mi.fieldNames = append(mi.fieldNames, selectorExpr.Sel.Name)
				}
				//fmt.Println(arg, reflect.TypeOf(arg))
			}
		}

		//fmt.Println(callExprSelector.X, reflect.TypeOf(callExprSelector.X))
		if callExpr2, ok := callExprSelector.X.(*ast.CallExpr); ok {
			return indexCallStackRecursive(modelAbbr, callExpr2, mi)
			//fmt.Println("args2", callExpr2.Args)
		}
		//fmt.Println(reflect.TypeOf(callExprSelector.X))
		//fmt.Println(callExprSelector.Sel)
	}

	return nil
}
