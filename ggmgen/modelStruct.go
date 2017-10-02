package main

import "github.com/imdario/mergo"

type ModelStruct struct {
	Name      string
	TableName string
	fields    []modelField
	indexes   []modelIndex
	notify    *pgNotify
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
	var one2oneUniqueIndex modelIndex
	one2oneUniqueIndex.isUnique = true
	one2oneUniqueIndex.isCoalesce = true

	var isOne2OneRelation bool
	for _, f := range ms.fields {
		if f.IsForeignKey && f.Relation != nil && f.Relation.isOne2One {
			isOne2OneRelation = true
			one2oneUniqueIndex.modelName = ms.Name
			one2oneUniqueIndex.fieldNames = append(one2oneUniqueIndex.fieldNames, f.Name)
		}
	}

	if isOne2OneRelation {
		//fmt.Println(one2oneUniqueIndex)
		//fmt.Println(one2oneUniqueIndex.CreateIndexSQL())
		ms.indexes = append(ms.indexes, one2oneUniqueIndex)
	}
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
