package rslt


type Error interface {
	Result
	Error()
}


type internalError struct {
	err error
}


func newError(err error) Error {
	result := internalError{
		err: err,
	}

	return &result
}

func (result *internalError) Error() {
	// Nothing here.
}


func (result *internalError)  Result() (interface{}, error, Warning) {
	return nil, result.err, nil
}
