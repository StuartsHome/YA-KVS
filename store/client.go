package store

//go:generate mockery -case=underscore -outpkg mock_store -output ../mock/mock_store -name=Store

type Response struct {
	Message string
}

type Store interface {
	Get(key string) (int, error)
	Put(key string, val int) error
}
