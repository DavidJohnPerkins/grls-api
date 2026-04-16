package store

import (
	"fmt"
)

type RecordCreationError struct {
	ErrorText string
}

type RecordNotFoundError struct {
	Id int
}

func (e *RecordCreationError) Error() string {
	return fmt.Sprintf("Error during model insertion: %v", e.ErrorText)
}

func (e *RecordNotFoundError) Error() string {
	return fmt.Sprintf("Record not found with id: %v", e.Id)
}
