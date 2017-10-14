package main

import "github.com/imdario/mergo"

type ModelStruct struct {
	Name                 string
	TableName            string
	fields               []modelField
	indexes              []modelIndex
	notify               *pgNotify
	relations            []modelRelation
	IsTableNameSetByUser bool

	directRelationsCache []modelRelation
	longRelationCache    []modelRelation
}

func (ms ModelStruct) HasPrimaryKey() bool {
	for _, f := range ms.fields {
		if f.IsPrimaryKey {
			return true
		}
	}
	return false
}

func (ms ModelStruct) PrimaryKey() *modelField {
	for fi, f := range ms.Fields() {
		if f.IsPrimaryKey {
			return &ms.Fields()[fi]
		}
	}
	return nil
}

func (ms ModelStruct) PrimaryKeys() []modelField {
	var keys []modelField
	for fi, f := range ms.Fields() {
		if f.IsPrimaryKey {
			keys = append(keys, ms.Fields()[fi])
		}
	}
	return keys
}

func (ms ModelStruct) IsMany2ManyTable() bool {
	fks := ms.ForeignKeys()
	if len(fks) < 2 {
		return false
	}
	for i := 0; i < len(fks); i++ {
		if i+1 < len(fks) && fks[i].field.Name != fks[i+1].field.Name {
			return true
		}
	}
	return false
}

func (ms ModelStruct) IsMany2ManyTableWithoutData() bool {
	for _, f := range ms.fields {
		if !f.IsForeignKey && !f.IsPrimaryKey {
			return false
		}
	}
	return true
}

func (ms *ModelStruct) DirectRelations() []modelRelation {
	if ms.directRelationsCache == nil {
		if ms.IsMany2ManyTableWithoutData() {
			return nil
		}

		var modelRelations []modelRelation
		for _, f := range ms.fields {
			if f.Relation != nil {
				newRelation := modelRelation{ModelFrom: ms, ModelTo: f.Relation.modelTo}
				if f.IsUnique() {
					newRelation.RelationType = ONE2ONE
				} else {
					newRelation.RelationType = ONE2MANY
				}
				modelRelations = append(modelRelations, newRelation)
			}
		}

		for _, m := range pkgS.Models {
			if m.Name == ms.Name {
				continue
			}
			if !m.IsMany2ManyTableWithoutData() {
				for _, f := range m.fields { // adding reverse relations
					if f.Relation != nil && f.Relation.modelTo.Name == ms.Name {
						var newReverseRelation modelRelation
						newReverseRelation.ModelFrom = ms
						newReverseRelation.ModelTo = m
						if f.IsUnique() {
							newReverseRelation.RelationType = ONE2ONE
						} else {
							newReverseRelation.RelationType = MANY2ONE
						}

						modelRelations = append(modelRelations, newReverseRelation)
						break
					}
				}
			}

			if m.IsMany2ManyTable() { // adding many2many relations
				var isCurrentTableRelation bool
				for _, fk := range m.ForeignKeys() {
					if fk.modelTo.Name == ms.Name {
						isCurrentTableRelation = true
						break
					}
				}
				if isCurrentTableRelation {
					for _, fk := range m.ForeignKeys() {
						if fk.modelTo.Name == ms.Name {
							continue
						}
						var newRelation modelRelation
						newRelation.ModelFrom = ms
						newRelation.ViaModel = m
						newRelation.ModelTo = fk.modelTo
						newRelation.RelationType = MANY2MANY
						modelRelations = append(modelRelations, newRelation)
					}
				}

			}
		}

		ms.directRelationsCache = modelRelations
	}

	return ms.directRelationsCache
}

func (ms *ModelStruct) LongRelations() []modelRelation {
	if ms.longRelationCache == nil {
		directRelations := ms.DirectRelations()
		var longRelationNames = make(map[string]bool)
		for _, dr := range directRelations {
			longRelationNames[dr.ModelTo.Name] = true
		}

		for _, dr := range directRelations {
			ms.longRelationCache = append(ms.longRelationCache, directRelationRecursiveParse(ms, longRelationNames, dr)...)
		}
	}

	return ms.longRelationCache
}

func directRelationRecursiveParse(ms *ModelStruct, longRelationNames map[string]bool, relation modelRelation) []modelRelation {
	var longRelations []modelRelation

	//for _, dr := range relations {
	for _, longRelation := range relation.ModelTo.DirectRelations() {
		if longRelation.ModelTo.Name == ms.Name {
			continue
		}
		if _, exist := longRelationNames[longRelation.ModelTo.Name]; !exist {
			longRelationNames[longRelation.ModelTo.Name] = true

			var newLongRelation modelRelation
			newLongRelation.ModelFrom = ms
			newLongRelation.RelationType = LONG
			newLongRelation.ModelTo = longRelation.ModelTo
			//fmt.Println(newLongRelation)
			longRelations = append(longRelations, newLongRelation)
			if longLong := directRelationRecursiveParse(ms, longRelationNames, longRelation); len(longLong) > 0 {
				longRelations = append(longRelations, longLong...)
			}

		}
	}
	//}
	return longRelations
}

func (ms ModelStruct) Relations() []modelRelation {
	return append(ms.DirectRelations(), ms.LongRelations()...)
	//var allRelations []modelRelation
	//directRelations := ms.DirectRelations()
	//allRelations = append(allRelations, directRelations...)
	//
	//return allRelations
	//for _, relation := range directRelations {
	//	relation.ModelTo.DirectRelations()
	//}
}

//func (ms ModelStruct) RelatedModels() []ModelStruct {
//	var relatedModels []ModelStruct
//	for _, m := range pkgS.Models {
//
//	}
//}
//
//func isInRelation(model1 *ModelStruct, model2 *ModelStruct) bool {
//
//}

//func (ms ModelStruct) TableFields() []modelField {
//	fields := ms.fields
//	for i:=0; i<len(fields); i++ {
//		if fields[i].Relation != nil && fields[i].Relation.
//	}
//	return ms.fields
//}

func (ms ModelStruct) CreateTableIfNotExistCommand() string {

	//var lines []string
	//for _, f := range ms.AllFields() {
	//	lines = append(lines, `"`+f.TableName()+`" `+f.SqlType())
	//}

	//for _, fk := range ms.ForeignKeys() {
	//	lines = append(lines, fk.SqlCreateTable())
	//}

	//return `CREATE TABLE IF NOT EXISTS "` + ms.TableName + `" (` + "\n\t\t" + strings.Join(lines, ",\n\t\t") + "\n\t\t);"
	return `CREATE TABLE IF NOT EXISTS "` + ms.TableName + `" ();`
}

func (ms ModelStruct) Fields() []modelField {
	var fields []modelField
	for fi, f := range ms.fields {
		if !f.IsForeignKey {
			fields = append(fields, ms.fields[fi])
		}
	}
	return fields
}

func (ms ModelStruct) AllFields() []modelField {
	return ms.fields
}

func (ms ModelStruct) PrimaryFields() []modelField {
	var primaryFields []modelField
	for pi, pf := range ms.fields {
		if pf.IsPrimaryKey {
			primaryFields = append(primaryFields, ms.fields[pi])
		}
	}
	return primaryFields
}

func (ms ModelStruct) NotPrimaryFields() []modelField {
	var notPrimaryFields []modelField
	for pi, pf := range ms.fields {
		if !pf.IsPrimaryKey {
			notPrimaryFields = append(notPrimaryFields, ms.fields[pi])
		}
	}
	return notPrimaryFields
}

func (ms ModelStruct) ForeignKeys() []tableForeignRelation {
	var relations []tableForeignRelation
	for _, f := range ms.AllFields() {
		if f.Relation != nil {
			relations = append(relations, *f.Relation)
		}
	}

	return relations
}

func (ms ModelStruct) Indexes() []modelIndex {
	return ms.indexes
}

func (ms *ModelStruct) AddFields(mfs []modelField) {
	if mfs == nil {
		return
	}
	for _, mf := range mfs {
		ms.AddField(mf)
	}
}

func (ms *ModelStruct) AddField(mf modelField) {
	for fi, f := range ms.fields {
		if f.Name == mf.Name {
			mergo.Merge(&ms.fields[fi], &mf)
			return
		}
	}
	mf.Model = ms

	ms.fields = append(ms.fields, mf)
}

func (ms *ModelStruct) GetFieldByName(fieldName string) *modelField {
	for fi, f := range ms.fields {
		if f.Name == fieldName {
			return &ms.fields[fi]
		}
	}
	return nil
}

func (ms ModelStruct) Notify() *pgNotify {
	return ms.notify
}

func (ms ModelStruct) UniqueTypeFields() []modelField {
	var fTypes []modelField
	for _, f := range ms.Fields() {
		appendFieldWithUniqueType(&fTypes, f)
		//if f.ConstType.ConstType == fieldType.DateType {
		//	fmt.Println(f.Name, f.ConstType, f.ConstType.IsNullable)
		//}

		if f.IsPointer {
			f2 := f
			f2.IsPointer = false
			appendFieldWithUniqueType(&fTypes, f2)
		}
	}

	//fmt.Println(ms.Name)
	//for _, ft := range fTypes {
	//	fmt.Println(ft.Type.TemplateName(), ft.Type.IsNullable)
	//}
	//fmt.Println()

	return fTypes
}

func (ms ModelStruct) NotScannerFields() []modelField {
	var nsFields []modelField
	for _, f := range ms.Fields() {
		if !f.Type().ImplementScannerInterface() {
			nsFields = append(nsFields, f)
		}
	}
	return nsFields
}

func appendFieldWithUniqueType(fields *[]modelField, newField modelField) {
	var exist bool
	for _, f := range *fields {
		if f.Type().Name() == newField.Type().Name() {
			exist = true
			break
		}
	}
	if !exist {
		*fields = append(*fields, newField)
	}

}
