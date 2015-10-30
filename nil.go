package rslt


type Nil interface {
	Result
	Nil()
}


type internalNil struct{}


func newNil() Nil {
	result := internalNil{}

	return &result
}


func (result *internalNil) Nil() {
	// Nothing here.
}


func (result *internalNil)  Result() (interface{}, error, Warning) {
	return nil, nil, nil
}

