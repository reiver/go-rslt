package rslt


type Value interface {
	Result
	Value()
}


type internalValue struct {
	value interface{}
	warning  Warning
}


func newValue(value interface{}, warning Warning) Value {
	result := internalValue{
		value: value,
		warning: warning,
	}

	return &result
}


func (result *internalValue) Value() {
	// Nothing here.
}

func (result *internalValue) Result() (interface{}, error, Warning) {
	return result.value, nil, result.warning
}
