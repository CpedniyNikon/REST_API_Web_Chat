package errors

type UserAlreadyExistsError struct {
	Message string
}

func (err UserAlreadyExistsError) Error() string {
	return err.Message
}
