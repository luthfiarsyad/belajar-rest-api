package exception

type InternalServerError struct {
	Error error
}

func NewInternalServerError(err error) InternalServerError {
	return InternalServerError{err}
}
