package options

import (
	"fmt"
	"reflect"
	"strconv"
)

func defaultStruct(s interface{}) error {
	val := reflect.ValueOf(s)
	var typ reflect.Type
	if val.Kind() == reflect.Ptr {
		typ = val.Elem().Type()
	} else {
		typ = val.Type()
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldV := reflect.Indirect(val).Field(i)

		// If the field is a ptr, recurse
		if field.Type.Kind() == reflect.Ptr {
			if fieldV.IsNil() {
				fieldV.Set(reflect.New(fieldV.Type().Elem()))
			}
			err := defaultStruct(fieldV.Interface())
			if err != nil {
				return fmt.Errorf("error defaulting %s: %v", field.Name, err)
			}
			continue
		}

		// If the field is a ptr, recurse
		if field.Type.Kind() == reflect.Struct {
			newField := reflect.New(reflect.TypeOf(fieldV.Interface()))
			err := defaultStruct(newField.Interface())
			if err != nil {
				return fmt.Errorf("error defaulting %s: %v", field.Name, err)
			}
			fieldV.Set(newField.Elem())
			continue
		}

		defaultVal := field.Tag.Get("default")

		if defaultVal != "" {
			switch fieldV.Kind() {
			case reflect.String:
				fieldV.Set(reflect.ValueOf(defaultVal).Convert(field.Type))
			case reflect.Bool:
				boolVal, err := strconv.ParseBool(defaultVal)
				if err != nil {
					return fmt.Errorf("expected default value (%s) to be bool: %v", defaultVal, err)
				}
				fieldV.Set(reflect.ValueOf(boolVal).Convert(field.Type))
			case reflect.Int64:
				intVal, err := strconv.ParseInt(defaultVal, 10, 64)
				if err != nil {
					return fmt.Errorf("expected default value (%s) to be int: %v", defaultVal, err)
				}
				fieldV.Set(reflect.ValueOf(intVal).Convert(field.Type))
			}

		}
	}
	return nil
}
