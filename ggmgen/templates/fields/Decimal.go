package fields

const DecimalTemplate = `
type whereFieldDecimal{{.ModelName}} struct {
	name  string
	where *{{lower .ModelName}}Where
}

func (wfi whereFieldDecimal{{.ModelName}}) Is(val ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" = '"+val.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) Eq(val ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" = '"+val.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) Equal(val ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" = '"+val.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) IsNot(val ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" <> '"+val.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) GreaterThan(val ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" > '"+val.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) GT(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.GreaterThan(val)
}
func (wfi whereFieldDecimal{{.ModelName}}) GreaterThanOrEqual(val ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" >= '"+val.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) GTE(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.GreaterThanOrEqual(val)
}
func (wfi whereFieldDecimal{{.ModelName}}) LessThan(val ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" < '"+val.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) LT(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.LessThan(val)
}
func (wfi whereFieldDecimal{{.ModelName}}) LessThanOrEqual(val ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" <= '"+val.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) LTE(val ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.LessThanOrEqual(val)
}

func (wfi whereFieldDecimal{{.ModelName}}) Between(left, right ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" BETWEEN '"+left.String()+"' AND '"+right.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) NotBetween(left, right ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond("\""+wfi.name+"\" NOT BETWEEN '"+left.String()+"' AND '"+right.String()+"'")
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) In(nums []ggm.Decimal) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	var cond string
	cond += "\"" + wfi.name + "\" IN ("
	for i, num := range nums {
		cond += "'"+num.String()+"'"
		if i != len(nums) {
			cond += ", "
		}
	}
	cond += ")"
	wfi.where.addCond(cond)
	return wfi.where
}
func (wfi whereFieldDecimal{{.ModelName}}) Any(nums ...ggm.Decimal) *{{lower .ModelName}}Where {
	return wfi.In(nums)
}
`
