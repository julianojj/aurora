package exceptions

type NotFoundException struct {
	Message string
	Code    int
}

func NewNotFoundException(message string) *NotFoundException {
	return &NotFoundException{
		Message: message,
		Code:    404,
	}
}

func (e *NotFoundException) Error() string {
	return e.Message
}
