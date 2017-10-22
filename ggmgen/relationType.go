package main

import "strings"

type modelRelation struct {
	ModelFrom    *ModelStruct
	ModelTo      *ModelStruct
	RelationType RelationType
	ViaModel     *ModelStruct
}

func (mr modelRelation) Name() string {
	return mr.RelationType.String() + ":" + mr.ModelFrom.Name + "->" + mr.ModelTo.Name
}

// func (mr modelRelation) Name() string {
// 	if mr.Field != nil {
// 		return mr.RelationType.String() + ":" + mr.ModelFrom.Name + "." + mr.Field.Name + "->" + mr.ModelTo.Name
// 	}
// 	return mr.ShortName()
// }

func (mr modelRelation) IsOneToXRelation() bool {
	return (mr.RelationType == ONE2ONE || mr.RelationType == ONE2MANY)
}

type ForeignRelationSlice []string

func (frs ForeignRelationSlice) String() string {
	var result string
	for i, text := range frs {
		if i != 0 {
			result += ", "
		}
		result += `"` + text + `"`
	}
	return result
}

func (frs ForeignRelationSlice) SqlString() string {
	var result string
	for _, text := range frs {
		result += "\n\t" + text
	}
	return result
}

// func (mr modelRelation) FieldCondition() string {
// 	return "\"" + mr.ModelFrom.TableName + "\".\"" + mr.Field.TableName() + "\"=\"" + mr.ModelTo.TableName + "\".\"" + mr.ModelTo.PrimaryKey().TableName() + "\""
// }

func (mr modelRelation) SqlJoin(fieldName string) ForeignRelationSlice {
	var modelFrom, modelTo *ModelStruct
	modelFrom = mr.ModelFrom
	modelTo = mr.ModelTo

	var fieldFrom, fieldTo *modelField

	if mr.ViaModel != nil {
		modelFromFKs := mr.ViaModel.ForeignKeysTo(modelFrom)
		modelToFKs := mr.ViaModel.ForeignKeysTo(modelTo)
		if len(modelFromFKs) > 1 || len(modelToFKs) > 1 {
			panic("cannot build sql join for join: more than 1 FK relation " + mr.Name())
		}
		if len(modelFromFKs) == 0 || len(modelToFKs) == 0 {
			panic("cannot build sql join for join: no FK relation " + mr.Name())
		}
		fieldFrom = modelFromFKs[0].Field
		fieldTo = modelToFKs[0].Field

		return []string{
			`INNER JOIN "` + mr.ViaModel.TableName + `" ON "` + mr.ViaModel.TableName + `"."` + fieldFrom.TableName() + `" = "` + modelFrom.TableName + `"."` + modelTo.PrimaryKey().TableName() + `"`,
			`INNER JOIN "` + modelTo.TableName + `" ON "` + modelTo.TableName + `"."` + modelTo.PrimaryKey().TableName() + `" = "` + mr.ViaModel.TableName + `"."` + fieldTo.TableName() + `"`,
		}
		// return []string{
		// 	`INNER JOIN "` + mr.ViaModel.TableName + `" ON "` + modelFrom.TableName + `"."` + fromFieldName + `" = "` + mr.ViaModel.TableName + `"."` + mr.ViaModel.PrimaryKey().TableName() + `"`,
		// 	`INNER JOIN "` + modelTo.TableName + `" ON "` + mr.ViaModel.TableName + `"."` + mr.ViaModel.PrimaryKey().TableName() + `" = "` + modelTo.TableName + `"."` + toFieldName + `"`,
		// }
	} else {
		var conditions []string

		if mr.RelationType == ONE2MANY {
			modelFrom = mr.ModelFrom

			modelTo = mr.ModelTo
			fieldTo = modelTo.PrimaryKey()

			modelFromFKs := mr.ModelFrom.ForeignKeysTo(modelTo)
			if len(modelFromFKs) > 1 && fieldName != "" {
				for _, fk := range modelFromFKs {
					if fk.Field.TableName() == fieldName {
						modelFromFKs = []tableForeignRelation{fk}
						break
					}
				}
			}
			for _, fk := range modelFromFKs {
				conditions = append(conditions, `"`+modelFrom.TableName+`"."`+fk.Field.TableName()+`" = "`+modelTo.TableName+`"."`+fieldTo.TableName()+`"`)
			}
		} else {
			modelFrom = mr.ModelFrom
			fieldFrom = mr.ModelFrom.PrimaryKey()

			modelTo = mr.ModelTo

			modelToFKs := mr.ModelTo.ForeignKeysTo(modelFrom)
			if len(modelToFKs) > 1 && fieldName != "" {
				for _, fk := range modelToFKs {
					if fk.Field.TableName() == fieldName {
						modelToFKs = []tableForeignRelation{fk}
						break
					}
				}
			}
			for _, fk := range modelToFKs {
				conditions = append(conditions, `"`+modelFrom.TableName+`"."`+fieldFrom.TableName()+`" = "`+modelTo.TableName+`"."`+fk.Field.TableName()+`"`)
			}
		}

		// if mr.RelationType == MANY2ONE {
		// 	return []string{
		// 		`INNER JOIN "` + mr.ModelTo.TableName + `" ON "` + mr.ModelTo.TableName + `"."` + mr.Field.TableName() + `" = "` + mr.ModelFrom.TableName + `"."` + mr.ModelTo.PrimaryKey().TableName() + `"`,
		// 	}
		// } else {

		return []string{
			`INNER JOIN "` + modelTo.TableName + `" ON ` + strings.Join(conditions, " OR "),
		}
		// }

	}
}

type RelationType int8

const (
	ONE2ONE RelationType = iota + 1
	ONE2MANY
	MANY2ONE
	MANY2MANY
	LONG
)

//
//func (rt RelationType) IsRelation() bool {
//	if rt == UNRELATED {
//		return false
//	}
//	return true
//}
//
//func (rt RelationType) IsDirectRelation() bool {
//	if rt == LONG_RELATION || rt == UNRELATED {
//		return false
//	}
//	return true
//}
//
func (rt RelationType) String() string {
	switch rt {
	case ONE2ONE:
		return "ONE2ONE"
	case ONE2MANY:
		return "ONE2MANY"
	case MANY2ONE:
		return "MANY2ONE"
	case MANY2MANY:
		return "MANY2MANY"
	case LONG:
		return "LONG"
	}
	return ""
}
