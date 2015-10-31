package rslt


import (
	"testing"

	"errors"
	"reflect"
)


func TestWrapValue(t *testing.T) {

	bfalse      := false
	btrue       := true
	i5          := 5
	sFive       := "five"
	f64FiveZero := 5.0

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
			ExpectedValue: &bfalse,
			ExpectedKind: reflect.Ptr,
			Fn: func() Result {

				fn := func() (*bool, error) {
					return &bfalse, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			ExpectedValue: &bfalse,
			ExpectedKind: reflect.Ptr,
			Fn: func() Result {

				fn := func() (error, *bool) {
					return nil, &bfalse
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
			ExpectedValue: &btrue,
			ExpectedKind: reflect.Ptr,
			Fn: func() Result {

				fn := func() (*bool, error) {
					return &btrue, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			ExpectedValue: &btrue,
			ExpectedKind: reflect.Ptr,
			Fn: func() Result {

				fn := func() (error, *bool) {
					return nil, &btrue
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
			ExpectedValue: &i5,
			ExpectedKind: reflect.Ptr,
			Fn: func() Result {

				fn := func() (*int, error) {
					return &i5, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			ExpectedValue: &i5,
			ExpectedKind: reflect.Ptr,
			Fn: func() Result {

				fn := func() (error, *int) {
					return nil, &i5
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
			ExpectedValue: &sFive,
			ExpectedKind: reflect.Ptr,
			Fn: func() Result {

				fn := func() (*string, error) {
					return &sFive, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			ExpectedValue: &sFive,
			ExpectedKind: reflect.Ptr,
			Fn: func() Result {

				fn := func() (error, *string) {
					return nil, &sFive
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



		{
			ExpectedValue: &f64FiveZero,
			ExpectedKind: reflect.Ptr,
			Fn: func() Result {

				fn := func() (*float64, error) {
					return &f64FiveZero, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			ExpectedValue: &f64FiveZero,
			ExpectedKind: reflect.Ptr,
			Fn: func() Result {

				fn := func() (error, *float64) {
					return nil, &f64FiveZero
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


func TestWrapError(t *testing.T) {


	err1 := errors.New("The error!")


	tests := []struct {
		ExpectedErr  error
		ExpectedKind reflect.Kind
		Fn           func()Result
	}{
		{
			ExpectedErr:  err1,
			ExpectedKind: reflect.Bool,
			Fn: func() Result {

				fn := func() (bool, error) {
					return false, err1
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			ExpectedErr:  err1,
			ExpectedKind: reflect.Bool,
			Fn: func() Result {

				fn := func() (error, bool) {
					return err1, false
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			ExpectedErr:  err1,
			ExpectedKind: reflect.Bool,
			Fn: func() Result {

				fn := func() (bool, error) {
					return true, err1
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			ExpectedErr:  err1,
			ExpectedKind: reflect.Bool,
			Fn: func() Result {

				fn := func() (error, bool) {
					return err1, true
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			ExpectedErr:  err1,
			ExpectedKind: reflect.Int,
			Fn: func() Result {

				fn := func() (int, error) {
					return 5, err1
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			ExpectedErr:  err1,
			ExpectedKind: reflect.Int,
			Fn: func() Result {

				fn := func() (error, int) {
					return err1, 5
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			ExpectedErr:  err1,
			ExpectedKind: reflect.String,
			Fn: func() Result {

				fn := func() (string, error) {
					return "five", err1
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			ExpectedErr:  err1,
			ExpectedKind: reflect.String,
			Fn: func() Result {

				fn := func() (error, string) {
					return err1, "five"
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			ExpectedErr:  err1,
			ExpectedKind: reflect.Float64,
			Fn: func() Result {

				fn := func() (float64, error) {
					return 5.0, err1
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			ExpectedErr:  err1,
			ExpectedKind: reflect.Float64,
			Fn: func() Result {

				fn := func() (error, float64) {
					return err1, 5.0
				}

				result := Wrap(fn())

				return result
			},
		},
	}


	for testNumber, test := range tests {

		result := test.Fn()

		if _, ok := result.(Value); ok {
			t.Errorf("For test #%d, did not expect Result to be a Value but was.", testNumber)
			continue
		}
		if _, ok := result.(Nil); ok {
			t.Errorf("For test #%d, did not expect Result to be Nil but was.", testNumber)
			continue
		}

		if _, ok := result.(Error); !ok {
			t.Errorf("For test #%d, expected Result to be an Error but wasn't; actually was: %T", testNumber, result)
			continue
		} else if _, err, _ := result.Result(); nil == err {
			t.Errorf("For test #%d, expected error returned from Result to not be nil, but actually was %v.", testNumber, err)
			continue
		} else if actualErr := err; test.ExpectedErr != actualErr {
			t.Errorf("For test #%d, expected error returned from Result to be %v, but actually was %v.", testNumber, test.ExpectedErr, actualErr)
		}

	}
}


func TestWrapNil(t *testing.T) {

	tests := []struct {
		Fn func()Result
	}{
		{
			Fn: func() Result {

				fn := func() (*bool, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, *bool) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() (*float32, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, *float32) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() (*float64, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, *float64) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() (*int, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, *int) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() (*int8, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, *int8) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() (*int16, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, *int16) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() (*int32, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, *int32) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() (*int64, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, *int64) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() (*string, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, *string) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() ([]bool, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, []bool) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() ([]float32, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, []float32) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() ([]float64, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, []float64) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() ([]int, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, []int) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() ([]int8, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, []int8) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() ([]int16, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, []int16) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() ([]int32, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, []int32) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() ([]int64, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, []int64) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() (map[string]string, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, map[string]string) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() (map[string]interface{}, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, map[string]interface{}) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() (map[int]string, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, map[int]string) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() (interface{}, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, interface{}) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() (interface{Pow()}, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, interface{Pow()}) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() (func(), error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, func()) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() (func(string)bool, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, func(string)bool) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() (chan struct{}, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, chan struct{}) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() (<-chan int, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, <-chan int) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},



		{
			Fn: func() Result {

				fn := func() (chan<- float64, error) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
		{
			Fn: func() Result {

				fn := func() (error, chan<- float64) {
					return nil, nil
				}

				result := Wrap(fn())

				return result
			},
		},
	}


	for testNumber, test := range tests {

		result := test.Fn()

		if value, ok := result.(Value); ok {
			t.Errorf("For test #%d, did not expect Result to be a Value but was. %T %#v", testNumber, value, value.(*internalValue).value)
			continue
		}

		if _, ok := result.(Error); ok {
			t.Errorf("For test #%d, did not expect Result to be an Error but was.", testNumber)
			continue
		}
		if _, ok := result.(Nil); !ok {
			t.Errorf("For test #%d, expected Result to be Nil but wasn't; actually was: %T", testNumber, result)
			continue
		} else if value, _, warning := result.Result(); nil != value {
			t.Errorf("For test #%d, expected value returned from Result to be nil, but actually was %v.", testNumber, value)
			continue
		} else if nil != warning {
			t.Errorf("For test #%d, expected warning returned from Result to be nil, but actually was %v.", testNumber, warning)
			continue
		}

	}
}
