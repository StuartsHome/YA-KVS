package store

type store struct {
}

func NewStore() *store {
	return &store{}
}

func (st *store) Get() (Response, error) {
	return Response{}, nil
}
