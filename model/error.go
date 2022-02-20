package model

type error interface {
	Error() string
}

type DBError struct {
	Message string
}

func (db DBError) Error() string {
	return ""
}
