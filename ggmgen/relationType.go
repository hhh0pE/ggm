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

func (mr modelRelation) SqlJoin() ForeignRelationSlice {
	if mr.ViaModel != nil {
		return []string{
			`INNER JOIN "` + mr.ViaModel.TableName + `" ON "` + mr.ModelFrom.TableName + `"."` + mr.Field.TableName() + `" = "` + mr.ViaModel.TableName + `"."` + mr.ViaModel.PrimaryKey().TableName() + `"`,
			`INNER JOIN "` + mr.ModelTo.TableName + `" ON "` + mr.ViaModel.TableName + `"."` + mr.ViaModel.PrimaryKey().TableName() + `" = "` + mr.ModelTo.TableName + `"."` + mr.ModelTo.PrimaryKey().TableName() + `"`,
		}
	} else {
		return []string{
			`INNER JOIN "` + mr.ModelTo.TableName + `" ON "` + mr.ModelFrom.TableName + `"."` + mr.Field.TableName() + `" = "` + mr.ModelTo.TableName + `"."` + mr.ModelTo.PrimaryKey().TableName() + `"`,
		}
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
