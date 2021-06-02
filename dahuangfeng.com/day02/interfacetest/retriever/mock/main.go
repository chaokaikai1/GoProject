package mock

type Retriever struct {
	Contents string
}

func (r Retriever) Get(ur string) string {
	return r.Contents
}
