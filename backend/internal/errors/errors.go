package errors

type HttpError struct {
	Msg string
	Status int
}

func NewHttpError(code int, msg string) *HttpError {
	return &HttpError{
		Msg: msg,
		Status: code,
	}
}

func (err HttpError) Error() string {
	return err.Msg
}
