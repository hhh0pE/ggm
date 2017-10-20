package main

type modelRelation struct {
	Field        *modelField
	ModelFrom    *ModelStruct
	ModelTo      *ModelStruct
	RelationType RelationType
	ViaModel     *ModelStruct
}

func (mr modelRelation) String() string {
	return mr.RelationType.String() + ":" + mr.ModelFrom.Name + "->" + mr.ModelTo.Name
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

func (mr modelRelation) SqlJoin() ForeignRelationSlice {
	var modelFrom, modelTo *ModelStruct
	modelFrom = mr.ModelFrom
	modelTo = mr.ModelTo

	var fieldFrom, fieldTo *modelField

	if mr.ViaModel != nil {
		directRelations := mr.ViaModel.DirectRelations()
		for dri, dr := range directRelations {
			if dr.ModelTo.Name == modelFrom.Name {
				fieldFrom = directRelations[dri].Field
			}
			if dr.ModelTo.Name == modelTo.Name {
				fieldTo = directRelations[dri].Field
			}
		}
		return []string{
			`INNER JOIN "` + mr.ViaModel.TableName + `" ON "` + mr.ViaModel.TableName + `"."` + fieldFrom.TableName() + `" = "` + modelFrom.TableName + `"."` + modelTo.PrimaryKey().TableName() + `"`,
			`INNER JOIN "` + modelTo.TableName + `" ON "` + modelTo.TableName + `"."` + modelTo.PrimaryKey().TableName() + `" = "` + mr.ViaModel.TableName + `"."` + fieldTo.TableName() + `"`,
		}
		// return []string{
		// 	`INNER JOIN "` + mr.ViaModel.TableName + `" ON "` + modelFrom.TableName + `"."` + fromFieldName + `" = "` + mr.ViaModel.TableName + `"."` + mr.ViaModel.PrimaryKey().TableName() + `"`,
		// 	`INNER JOIN "` + modelTo.TableName + `" ON "` + mr.ViaModel.TableName + `"."` + mr.ViaModel.PrimaryKey().TableName() + `" = "` + modelTo.TableName + `"."` + toFieldName + `"`,
		// }
	} else {
		if mr.RelationType == ONE2MANY {
			modelFrom = mr.ModelFrom
			fieldFrom = mr.Field

			modelTo = mr.ModelTo
			fieldTo = modelTo.PrimaryKey()
		} else {
			modelFrom = mr.ModelFrom
			fieldFrom = mr.ModelFrom.PrimaryKey()

			modelTo = mr.ModelTo
			fieldTo = mr.Field
		}
		// if mr.RelationType == MANY2ONE {
		// 	return []string{
		// 		`INNER JOIN "` + mr.ModelTo.TableName + `" ON "` + mr.ModelTo.TableName + `"."` + mr.Field.TableName() + `" = "` + mr.ModelFrom.TableName + `"."` + mr.ModelTo.PrimaryKey().TableName() + `"`,
		// 	}
		// } else {

		return []string{
			`INNER JOIN "` + modelTo.TableName + `" ON "` + modelFrom.TableName + `"."` + fieldFrom.TableName() + `" = "` + modelTo.TableName + `"."` + fieldTo.TableName() + `"`,
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
