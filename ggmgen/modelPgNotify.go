package main

type pgNotify struct {
	Name       string
	ModelName  string
	fieldNames []string
	OnInsert   bool
	OnUpdate   bool
	OnDelete   bool

	Model *ModelStruct
}

func (pn pgNotify) Fields() []modelField {
	var fields []modelField
	for _, fieldName := range pn.fieldNames {
		if foundField := pn.Model.GetFieldByName(fieldName); foundField != nil {
			fields = append(fields, *foundField)
		}
	}
	return fields
}
