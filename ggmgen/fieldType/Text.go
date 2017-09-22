package fieldType

var Text = FieldType{ConstType: TextType, IsNullable: false}
var TextNullable = FieldType{ConstType: TextType, IsNullable: true}

var VarChar = FieldType{ConstType: TextType, IsNullable: false, MaxSize: 255}
var VarCharNullable = FieldType{ConstType: TextType, IsNullable: false, MaxSize: 255}
