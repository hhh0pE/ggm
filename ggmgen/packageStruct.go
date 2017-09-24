package main

type packageStruct struct {
	Name    string
	DirPath string
	Models  []*ModelStruct
}

func (ps packageStruct) GetModel(name string) *ModelStruct {
	for i, _ := range ps.Models {
		if ps.Models[i].Name == name {
			return ps.Models[i]
		}
	}
	return nil
}

func (ps packageStruct) HasModel(name string) bool {
	for _, m := range ps.Models {
		if m.Name == name {
			return true
		}
	}
	return false
}

func (ps *packageStruct) AddModel(m ModelStruct) {
	for i, _ := range ps.Models {
		if ps.Models[i].Name == m.Name {
			ps.Models[i].TableName = m.TableName
			return
		}
	}
	ps.Models = append(ps.Models, &m)
}

func (ps packageStruct) HasNotify() bool {
	for _, m := range ps.Models {
		if m.notify != nil {
			return true
		}
	}
	return false
}

func (ps packageStruct) Notifies() []pgNotify {
	var notifies []pgNotify
	for _, m := range ps.Models {
		if m.notify != nil {
			notifies = append(notifies, *m.notify)
		}
	}
	return notifies
}
