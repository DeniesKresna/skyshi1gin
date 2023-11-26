package helper

import (
	"errors"
	"reflect"
)

func WrapPercentOnStructString(data interface{}) (err error) {
	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		err = errors.New("Input must be a pointer to a struct")
		return
	}

	structVal := val.Elem()

	for i := 0; i < structVal.NumField(); i++ {
		field := structVal.Field(i)

		if field.Kind() == reflect.String {
			currentValue := field.String()

			if currentValue != "" {
				uppercaseValue := WrapString(currentValue, "%")
				field.SetString(uppercaseValue)
			}
		}
	}

	return
}
