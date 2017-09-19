package fields

const IntegerTemplate = `
type whereFieldInteger{{.ModelName}} struct {
	name  string
	where *{{lower .ModelName}}Where
}

func (wfi whereFieldInteger{{.ModelName}}) Is(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" = '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) Eq(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" = '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) Equal(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" = '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) IsNot(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" <> '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) GreaterThan(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" > '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) GT(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" > '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) GreaterThanOrEqual(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" >= '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) GTE(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" >= '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) LessThan(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" < '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) LT(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" < '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) LessThanOrEqual(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" <= '%d'", val))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) LTE(val int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" <= '%d'", val))
	return wfi.where
}

func (wfi whereFieldInteger{{.ModelName}}) Between(left, right int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" BETWEEN '%d' AND '%d'", left, right))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) NotBetween(left, right int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" NOT BETWEEN '%d' AND '%d'", left, right))
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) In(nums []int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	var cond string
	cond += "\"" + wfi.name + "\" IN ("
	for i, num := range nums {
		cond += fmt.Sprintf("'%d'", num)
		if i != len(nums) {
			cond += ", "
		}
	}
	cond += ")"
	wfi.where.addCond(cond)
	return wfi.where
}
func (wfi whereFieldInteger{{.ModelName}}) Any(nums ...int) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	var cond string
	cond += "\"" + wfi.name + "\" IN ("
	for i, num := range nums {
		cond += fmt.Sprintf("'%d'", num)
		if i != len(nums) {
			cond += ", "
		}
	}
	cond += ")"
	wfi.where.addCond(cond)
	return wfi.where
}
`
