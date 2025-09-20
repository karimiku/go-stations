package model

type ErrNotFound struct {
	Resource string
	ID       string
}

func (e *ErrNotFound) Error() string {
	if e.ID != "" {
		return e.Resource + " with ID " + e.ID + " not found"
	}
	return e.Resource + " not found"

}
