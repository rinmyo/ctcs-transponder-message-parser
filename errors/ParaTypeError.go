package errors

type ParaTypeError struct {
	info string
}

func NewParaTypeError(info string) *ParaTypeError {
	return &ParaTypeError{info: info}
}

func (err ParaTypeError) Error() string {
	return err.info
}
