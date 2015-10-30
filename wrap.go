package rslt


// Wrap returns a Result.
func Wrap(args ...interface{}) Result {

	var value   interface{}
	var warning Warning

	for _, arg := range args {
		if err, ok := arg.(error); ok {
			return newError(err)
		} else if warn, ok := arg.(Warning); ok {
			warning = warn
		} else {
			if nil != arg {
				value = arg
			}
		}
	}

	if nil == value {
		return newNil()
	}

	return newValue(value, warning)
}
