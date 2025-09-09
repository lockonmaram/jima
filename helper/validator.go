package helper

import (
	"fmt"
	"reflect"
	"strings"
)

type ValidationType string

const (
	ValidationTypeRequired ValidationType = "required"
	ValidationTypeEmail    ValidationType = "email"
)

func ValidateStruct(data any) error {
	dataType := getOriginalStructType(data)

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		jsonTag := field.Tag.Get("json")
		validationTags := strings.Split(field.Tag.Get("validation"), ",")

		err := validateField(jsonTag, reflect.ValueOf(data).Elem().Field(i), validationTags)
		if err != nil {
			return err
		}
	}

	return nil
}

func getOriginalStructType(data any) reflect.Type {
	dataType := reflect.TypeOf(data)
	if dataType.Kind() == reflect.Ptr {
		dataType = dataType.Elem()
	}
	return dataType
}

func validateField(jsonTag string, value reflect.Value, validationTags []string) error {
	for _, tag := range validationTags {
		switch ValidationType(tag) {
		case ValidationTypeRequired:
			if isEmptyValue(value) {
				return fmt.Errorf("%s is required", jsonTag)
			}
		case ValidationTypeEmail:
			if value.Kind() == reflect.String {
				if !isValidEmail(value.String()) {
					return fmt.Errorf("%s format is invalid", jsonTag)
				}
			}
		}
	}
	return nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

func isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}
