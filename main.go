package ggm

type modelWhere interface {
	andOr()
	addCond(string)
}

type ModelWithIndexes interface {
	Indexes()
}

//
//type ModelIndex struct {
//	_modelData struct {
//		isUnique   bool
//		model      Model
//		fieldNames []string
//	}
//}
//
//func (mi *ModelIndex) Index(isUnique bool, m Model, fieldNames ...string) {
//	mi._modelData.isUnique = isUnique
//	mi._modelData.model = m
//	mi._modelData.fieldNames = fieldNames
//}
//
//func (mi ModelIndex) CheckFields() error {
//	return nil
//}
//
//func (mi ModelIndex) Run(db *sql.DB) error {
//
//}
//
//type modelFK struct {
//	_modelData struct {
//		modelFrom     Model
//		fieldFromName string
//		modelTo       Model
//		fieldToName   string
//	}
//}
//
//func (mfk *modelFK) ForeignKey(modelFrom Model, fieldFromName string, modelTo Model, fieldToName string) {
//	mfk._modelData.modelFrom = modelFrom
//	mfk._modelData.fieldFromName = fieldFromName
//	mfk._modelData.modelTo = modelTo
//	mfk._modelData.fieldToName = fieldToName
//}
