package rslt


import (
	"testing"

	"reflect"
)


func TestWrapValue(t *testing.T) {

	tests := []struct {
		ExpectedValue interface{}
		ExpectedKind  reflect.Kind
		Fn            func()Result
	}{
		{
			ExpectedValue: false,
			ExpectedKind: reflect.Bool,
			Fn: func() Result {

				fn := func() (bool, error) {
					return false, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			ExpectedValue: false,
			ExpectedKind: reflect.Bool,
			Fn: func() Result {

				fn := func() (error, bool) {
					return nil, false
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			ExpectedValue: true,
			ExpectedKind: reflect.Bool,
			Fn: func() Result {

				fn := func() (bool, error) {
					return true, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			ExpectedValue: true,
			ExpectedKind: reflect.Bool,
			Fn: func() Result {

				fn := func() (error, bool) {
					return nil, true
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			ExpectedValue: 5,
			ExpectedKind: reflect.Int,
			Fn: func() Result {

				fn := func() (int, error) {
					return 5, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			ExpectedValue: 5,
			ExpectedKind: reflect.Int,
			Fn: func() Result {

				fn := func() (error, int) {
					return nil, 5
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			ExpectedValue: "five",
			ExpectedKind: reflect.String,
			Fn: func() Result {

				fn := func() (string, error) {
					return "five", nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			ExpectedValue: "five",
			ExpectedKind: reflect.String,
			Fn: func() Result {

				fn := func() (error, string) {
					return nil, "five"
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			ExpectedValue: 5.0,
			ExpectedKind: reflect.Float64,
			Fn: func() Result {

				fn := func() (float64, error) {
					return 5.0, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			ExpectedValue: 5.0,
			ExpectedKind: reflect.Float64,
			Fn: func() Result {

				fn := func() (error, float64) {
					return nil, 5.0
				}

				result := Wrap(fn())

				return result
			},
		},
 	}


	for testNumber, test := range tests {

		result := test.Fn()

		if _, ok := result.(Error); ok {
			t.Errorf("For test #%d, did not expect Result to be an Error but was.", testNumber)
			continue
		}
		if _, ok := result.(Nil); ok {
			t.Errorf("For test #%d, did not expect Result to be Nil but was.", testNumber)
			continue
		}

		if _, ok := result.(Value); !ok {
			t.Errorf("For test #%d, expected Result to be a Value but wasn't; actually was: %T", testNumber, result)
			continue
		} else if value, err, warning := result.Result(); nil != err {
			t.Errorf("For test #%d, expected error returned from Result to be nil, but actually was %v.", testNumber, err)
			continue
		} else if nil != warning {
			t.Errorf("For test #%d, expected warning returned from Result to be nil, but actually was %v.", testNumber, warning)
			continue
		} else if reflectedType := reflect.TypeOf(value); test.ExpectedKind != reflectedType.Kind() {
			t.Errorf("For test #%d, expected value returned from Result to be type %v, but actually was type %T.", testNumber, test.ExpectedKind, value)
			continue
		} else if actualValue := value; test.ExpectedValue != actualValue {
			t.Errorf("For test #%d, expected value returned from Result to be %v, but actually was %v.", testNumber, test.ExpectedValue, actualValue)
		}

	}
}
