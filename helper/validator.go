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
	ValidationTypeURI      ValidationType = "uri"
	ValidationTypeForm     ValidationType = "form"
)

func ValidateStruct(data any) error {
	dataType := getOriginalStructType(data)

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		validationTags := strings.Split(field.Tag.Get("validation"), ",")

		jsonTag := field.Tag.Get("json")
		uriTag := field.Tag.Get("uri")
		formTag := field.Tag.Get("form")
		tag := ""

		if jsonTag != "" {
			tag = jsonTag
		} else if uriTag != "" {
			tag = uriTag
			validationTags = append(validationTags, string(ValidationTypeURI))
		} else if formTag != "" {
			tag = formTag
		}

		err := validateField(tag, reflect.ValueOf(data).Elem().Field(i), validationTags)
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

func validateField(tag string, value reflect.Value, validationTags []string) error {
	for _, validation := range validationTags {
		switch ValidationType(validation) {
		case ValidationTypeRequired:
			if isEmptyValue(value) {
				return fmt.Errorf("%s is required", tag)
			}
		case ValidationTypeEmail:
			if value.Kind() == reflect.String {
				if !isValidEmail(value.String()) {
					return fmt.Errorf("%s format is invalid", tag)
				}
			}
		case ValidationTypeURI:
			if value.Kind() == reflect.String {
				if value.String() == fmt.Sprintf(":%s", tag) {
					return fmt.Errorf("parameter %s is required", tag)
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
