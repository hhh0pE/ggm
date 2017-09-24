package ggm

type indexParams struct {
}

func (ip *indexParams) Unique() *indexUnique {
	return &indexUnique{}
}
func (ip *indexParams) Name(name string) *indexParams {
	return ip
}

type indexUnique struct {
}

func (iu *indexUnique) Coalesce() {

}

func Index(fields ...interface{}) *indexParams {
	return &indexParams{}
}
