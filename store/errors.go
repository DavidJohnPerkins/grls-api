package store

import (
	"fmt"
)

type DuplicateKeyError struct {
	Id int64
}

type RecordNotFoundError struct {
	Id int64
}

func (e *DuplicateKeyError) Error() string {
	return fmt.Sprintf("Duplicate model id: %v", e.Id)
}

func (e *RecordNotFoundError) Error() string {
	return fmt.Sprintf("Record not found with id: %v", e.Id)
}
