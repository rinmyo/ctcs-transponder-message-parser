package errors

type ParaLengthError struct {
	info string
}

func NewParaLengthError(info string) *ParaLengthError {
	return &ParaLengthError{info: info}
}

func (err ParaLengthError) Error() string {
	return err.info
}
