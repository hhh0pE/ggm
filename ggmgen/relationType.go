package main

type modelRelation struct {
	Field    *modelField
	ModelFrom    *ModelStruct
	ModelTo      *ModelStruct
	RelationType RelationType
	ViaModel     *ModelStruct
}

func (mr modelRelation) String() string {
	return mr.RelationType.String()+":"+mr.ModelFrom.Name+"->"+mr.ModelTo.Name
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
