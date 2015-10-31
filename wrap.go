package rslt


import (
	"reflect"
)


// Wrap returns a Result.
func Wrap(args ...interface{}) Result {

	var value   interface{}
	var warning Warning

	for _, arg := range args {
		if err, ok := arg.(error); ok {
			return newError(err)
		} else if warn, ok := arg.(Warning); ok {
			warning = warn
		} else if nil == value {
//@TODO: Is there no better way to check if something is nil (regardless of the type)?
			reflectedValue := reflect.ValueOf(arg)

			switch reflectedValue.Kind() {
			case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice:
				if !reflectedValue.IsNil() {
					value = arg
				}
			default:
				value = arg
			}
		}
	}

	if nil == value {
		return newNil()
	}

	return newValue(value, warning)
}
