package exceptions

type ValidationException struct {
	Message string
	Code    int
}

func NewValidationException(message string) *ValidationException {
	return &ValidationException{
		Message: message,
		Code:    400,
	}
}

func (e *ValidationException) Error() string {
	return e.Message
}
