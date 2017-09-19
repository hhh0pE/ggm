package fields

const FloatTemplate = `
type whereFieldFloat{{.ModelName}} struct {
	name  string
	where *{{lower .ModelName}}Where
}

func (wfi whereFieldFloat{{.ModelName}}) Is(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" = '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) Eq(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" = '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) Equal(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" = '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) IsNot(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" <> '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) GreaterThan(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" > '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) GT(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" > '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) GreaterThanOrEqual(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" >= '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) GTE(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" >= '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) LessThan(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" < '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) LT(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" < '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) LessThanOrEqual(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" <= '%f'", val))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) LTE(val float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" <= '%f'", val))
	return wfi.where
}

func (wfi whereFieldFloat{{.ModelName}}) Between(left, right float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" BETWEEN '%f' AND '%f'", left, right))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) NotBetween(left, right float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	wfi.where.addCond(fmt.Sprintf("\""+wfi.name+"\" NOT BETWEEN '%f' AND '%f'", left, right))
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) In(nums []float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	var cond string
	cond += "\"" + wfi.name + "\" IN ("
	for i, num := range nums {
		cond += fmt.Sprintf("'%f'", num)
		if i != len(nums) {
			cond += ", "
		}
	}
	cond += ")"
	wfi.where.addCond(cond)
	return wfi.where
}
func (wfi whereFieldFloat{{.ModelName}}) Any(nums ...float64) *{{lower .ModelName}}Where {
	wfi.where.andOr()
	var cond string
	cond += "\"" + wfi.name + "\" IN ("
	for i, num := range nums {
		cond += fmt.Sprintf("'%f'", num)
		if i != len(nums) {
			cond += ", "
		}
	}
	cond += ")"
	wfi.where.addCond(cond)
	return wfi.where
}
`
