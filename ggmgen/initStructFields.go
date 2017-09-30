package main

import (
	"go/types"

	"strings"

	"log"

	"fmt"

	"github.com/hhh0pE/ggm/ggmgen/fieldType"
)

func initStructFields(typesInfo *types.Package, modelS *ModelStruct) {

	modelTI := typesInfo.Scope().Lookup(modelS.Name)

	if modelTI == nil {
		log.Println("modelTI == nil :(")
		return
	}

	if modelTI.Type() == nil {
		//analyze()
		log.Println("modelTI.Type() == nil :(")
		return
		//fmt.Println(modelTI == nil)
		//for i := 0; i < typesInfo.Scope().Len(); i++ {
		//	fmt.Println(i, typesInfo.Scope().Len())
		//	fmt.Println(typesInfo.Scope().Child(i))
		//}
	}
	//if modelTI.Type() == nil {
	//	return
	//}
	//fmt.Println("<<", modelTI, reflect.TypeOf(modelTI))
	//fmt.Println("type()", modelTI.Type())
	//fmt.Println("type().underlying()", modelTI.Type().Underlying())

	modelStruct, ok := modelTI.Type().Underlying().(*types.Struct)
	if !ok {
		return
	}

	//fmt.Println(modelS.Name)
	for i := 0; i < modelStruct.NumFields(); i++ {
		foundFields := scanField(modelStruct.Field(i))
		fieldTags := ParseFieldTags(modelStruct.Tag(i))
		for fi, _ := range foundFields {
			foundFields[fi].Tags = fieldTags
			if sqlTag, exist := foundFields[fi].FindTag([]string{"sql", "ggm"}); exist {
				sqlTag = strings.ToLower(sqlTag)
				if sqlTag == "pk" || sqlTag == "primary_key" || sqlTag == "primarykey" {
					foundFields[fi].IsPrimaryKey = true
				}
			}
		}

		for _, ff := range foundFields {
			if ff.IsForeignKey && ff.Relation != nil {
				ff.Model = modelS
				ff.Relation.field = &ff
			}
			modelS.AddField(ff)
			//fmt.Printf("%s: %#v\n", ff.Name, ff.Type)
		}
	}
	//fmt.Println()
}

func scanField(field *types.Var) []modelField {
	//fmt.Println("		scanField", field.Name())
	var newModelFields []modelField
	var newModelField modelField
	newModelField.Name = field.Name()
	if strings.ToLower(field.Name()) == "id" {
		newModelField.IsPrimaryKey = true
	}

	//var isPointer bool

	underlying := field.Type().Underlying()
	if typePointer, ok := underlying.(*types.Pointer); ok {
		newModelField.IsPointer = true
		newModelField.Type.IsNullable = true
		underlying = typePointer.Elem().Underlying()
	}

	var detectBasicType = func(bt *types.Basic, nmf *modelField) {
		kind := bt.Kind()
		switch {
		case types.Int <= kind && kind <= types.Uint64:
			nmf.Type.ConstType = fieldType.IntType
		case types.Float32 == kind || kind == types.Float64:
			nmf.Type.ConstType = fieldType.FloatType
		case kind == types.String:
			nmf.Type.ConstType = fieldType.TextType
		case kind == types.Bool:
			nmf.Type.ConstType = fieldType.BoolType
		default:
			return
		}

	}

	if typeBasic, ok := underlying.(*types.Basic); ok {
		detectBasicType(typeBasic, &newModelField)
		newModelFields = append(newModelFields, newModelField)
		return newModelFields
	}

	if typeSlice, ok := underlying.(*types.Slice); ok {
		if sliceElem, ok2 := typeSlice.Elem().(*types.Named); ok2 {
			if isImplementSqlScannerInterface(sliceElem.Obj()) {
				fmt.Println(sliceElem.Obj().Name())
			}
			if foundModel := pkgS.GetModel(sliceElem.Obj().Name()); foundModel != nil {
				log.Println("Found slice of models. Skipping field..")
				return nil
				//var newRelation tableForeignRelation
				//newRelation.ModelToName = sliceElem.Obj().Name()
				//newRelation.IsManyRelation = true
				//
				//newModelField.IsForeignKey = true
				//newModelField.Relation = &newRelation
				//newModelFields = append(newModelFields, newModelField)
				//return newModelFields
			}

		}
		if basicElem, ok := typeSlice.Elem().(*types.Basic); ok {
			detectBasicType(basicElem, &newModelField)
			newModelField.Type.IsArray = true
			newModelFields = append(newModelFields, newModelField)
			return newModelFields

		}
	}
	if typeStruct, ok := underlying.(*types.Struct); ok {
		if field.Anonymous() { // is embedded
			// if embedded another Model - it's one2one relation
			if foundModel := pkgS.GetModel(field.Name()); foundModel != nil {

				newModelField.Name = field.Name()
				newModelField.Type = fieldType.Integer
				newModelField.IsForeignKey = true
				newModelField.Relation = new(tableForeignRelation)
				newModelField.Relation.isOne2One = true
				newModelField.Relation.field = &newModelField
				newModelField.Relation.modelTo = foundModel

				newModelFields = append(newModelFields, newModelField)
			} else { // if embedded not Model - it's just Go embedding
				for i := 0; i < typeStruct.NumFields(); i++ {
					mField := typeStruct.Field(i)
					scannedFields := scanField(mField)
					if len(scannedFields) > 0 {
						newModelFields = append(newModelFields, scannedFields...)
					}
				}
			}
		} else {
			fType := field.Type()
			var typeNamed *types.Named
			//fmt.Println(reflect.TypeOf(fType))
			if typePointer, ok := fType.(*types.Pointer); ok {
				newModelField.Type.IsNullable = true
				typeNamed = typePointer.Elem().(*types.Named)
			}
			if tNamed, ok := fType.(*types.Named); ok {
				typeNamed = tNamed
			}

			switch typeNamed.String() {
			case "database/sql.NullInt64":
				newModelField.Type = fieldType.IntegerNullable
			case "database/sql.NullFloat64":
				newModelField.Type = fieldType.NumericNullable
			case "database/sql.NullString":
				newModelField.Type = fieldType.TextNullable
			case "database/sql.NullBool":
				newModelField.Type = fieldType.BooleanNullable
			case "time.Time":
				newModelField.Type = fieldType.Date
			case "github.com/lib/pq.NullTime":
				newModelField.Type = fieldType.DateNullable

			//case "github.com/lib/pq.StringArray":
			//	newModelField.Type.ConstType = fieldType.TextType
			//	newModelField.Type.IsArray = true
			//case "github.com/lib/pq.Int64Array":
			//	newModelField.Type.ConstType = fieldType.IntType
			//	newModelField.Type.IsArray = true
			//case "github.com/lib/pq.Float64Array":
			//	newModelField.Type.ConstType = fieldType.FloatType
			//	newModelField.Type.IsArray = true
			//case "github.com/lib/pq.BoolArray":
			//	newModelField.Type.ConstType = fieldType.BoolType
			//	newModelField.Type.IsArray = true
			default:
				isSqlField := isImplementSqlScannerInterface(typeNamed.Obj())
				if isSqlField {
					fmt.Println("isSqlField", typeNamed.Obj().Name())
					if typeStruct.NumFields() == 1 {
						if scannedFields := scanField(typeStruct.Field(0)); len(scannedFields) > 0 {
							scannedField := scannedFields[0]
							scannedField.Name = field.Name()
							newModelFields = append(newModelFields, scannedField)
							return newModelFields
						}
					}
				}

				if foundModel := pkgS.GetModel(typeNamed.Obj().Name()); foundModel != nil {
					newModelField.IsForeignKey = true
					newModelField.Type.ConstType = fieldType.IntType
					var tableFK tableForeignRelation
					//tableFK.field = &newModelField
					tableFK.modelTo = foundModel
					//tableFK.ModelToName = typeNamed.Obj().Name()

					newModelField.Relation = &tableFK
				}

			}

			newModelFields = append(newModelFields, newModelField)
			return newModelFields

		}
		return newModelFields

		//fmt.Println(reflect.TypeOf(typeStruct.Underlying()))
	}

	//fmt.Println(reflect.TypeOf(underlying))
	return nil
}

//func isImplementModelInterface(tnObj *types.TypeName) bool {
//	tnObjP := types.NewPointer(tnObj.Type())
//
//	msp := types.NewMethodSet(tnObjP)
//	tn := msp.Lookup(tnObj.Pkg(), "TableName")
//
//	if tn == nil {
//		return false
//	}
//
//	tnF := tn.Obj().(*types.Func)
//	tnFS := tnF.Type().(*types.Signature)
//	if tnFS.Results().Len() != 1 {
//		return false
//	}
//	if tnFS.Results().At(0).Type().String() != "string" {
//		return false
//	}
//
//	return true
//}

func isImplementSqlScannerInterface(tnObj *types.TypeName) bool {
	tnObjP := types.NewPointer(tnObj.Type())

	msp := types.NewMethodSet(tnObjP)
	sf := msp.Lookup(tnObj.Pkg(), "Scan")

	ms := types.NewMethodSet(tnObj.Type())
	vf := ms.Lookup(tnObj.Pkg(), "Value")

	// if not implemented one of sql.Scanner methods - return false
	if vf == nil || sf == nil {
		return false
	}

	if vf != nil {
		vfFunc := vf.Obj().(*types.Func)
		if vfFuncSign, ok := vfFunc.Type().(*types.Signature); ok {
			if vfFuncSign.Results().Len() != 2 {
				return false
			}
			if vfFuncSign.Results().At(0).Type().String() != "database/sql/driver.Value" {
				return false
			}
			if vfFuncSign.Results().At(1).Type().String() != "error" {
				return false
			}
		}

	}

	if sf != nil {
		sfFunc := sf.Obj().(*types.Func)
		if sfFuncSign, ok := sfFunc.Type().(*types.Signature); ok {
			if sfFuncSign.Results().Len() != 1 {
				return false
			}
			if sfFuncSign.Results().At(0).Type().String() != "error" {
				return false
			}

			if sfFuncSign.Params().Len() != 1 {
				return false
			}
			if sfFuncSign.Params().At(0).Type().String() != "interface{}" {
				return false
			}
		}
	}

	return true
}
