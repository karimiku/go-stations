package model

type ErrNotFound struct {
	ID int
}

func (e *ErrNotFound) Error() string {
	return "not found"
}
