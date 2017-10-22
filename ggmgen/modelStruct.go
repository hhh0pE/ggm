package main

import (
	"fmt"

	"github.com/imdario/mergo"
)

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
		if i+1 < len(fks) && fks[i].Field.Name != fks[i+1].Field.Name {
			return true
		}
	}
	return false
}

func (ms ModelStruct) IsMany2ManyTableWithoutData() bool {
	for fi, f := range ms.fields {
		if !f.IsForeignKey && !f.IsPrimaryKey {
			return false
		}
		if fi > 0 && (ms.fields[fi-1].Relation != nil && ms.fields[fi].Relation != nil &&
			ms.fields[fi-1].Relation.modelTo.Name == ms.fields[fi].Relation.modelTo.Name) {
			return false
		}
	}
	return true
}

func (ms *ModelStruct) DirectRelations() []modelRelation {
	if ms.directRelationsCache == nil {
		if ms.IsMany2ManyTableWithoutData() {
			fmt.Println("Many2ManyWithoutData", ms.Name)
			return nil
		}

		var relationExistMap = make(map[string]bool)

		var modelRelations []modelRelation
		for _, f := range ms.fields {
			if f.Relation != nil {
				newRelation := modelRelation{ModelFrom: ms, ModelTo: f.Relation.modelTo}
				if f.IsUnique() {
					newRelation.RelationType = ONE2ONE
				} else {
					newRelation.RelationType = ONE2MANY
				}

				if _, exist := relationExistMap[newRelation.Name()]; !exist {
					relationExistMap[newRelation.Name()] = true
					modelRelations = append(modelRelations, newRelation)
				}
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
						// newRelation.Field = fk.Field
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
}

func (ms ModelStruct) CreateTableIfNotExistCommand() string {
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

func (ms ModelStruct) ForeignKeysTo(modelTo *ModelStruct) []tableForeignRelation {
	var relationsTo []tableForeignRelation
	for _, fk := range ms.ForeignKeys() {
		if fk.ModelTo().Name == modelTo.Name {
			relationsTo = append(relationsTo, fk)
		}
	}
	return relationsTo
}

func (ms ModelStruct) Indexes() []modelIndex {
	return ms.indexes
}
func (ms ModelStruct) MustBeUniqueFields() []modelField {
	var fields []modelField
	for _, index := range ms.Indexes() {
		if index.IsRealCoalesce() {
			for _, indexField := range index.Fields() {
				var alreadyExist bool
				for _, alreadyExistField := range fields {
					if IsModelFieldEqual(indexField, alreadyExistField) {
						alreadyExist = true
						break
					}
				}
				if !alreadyExist {
					fields = append(fields, indexField)
				}
			}
		}
	}

	return fields
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
