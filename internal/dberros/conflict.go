// Defines custom error types (ConflictError and NotFoundError) for handling database conflicts and entity not found errors.
//Defines the ConflictError type for handling conflicts during database operations.
package dberrors

type ConflictError struct{}

func (e *ConflictError) Error() string {
	return "attempted to create a record with an existing key"
}
