package util

import (
	"errors"
	"reflect"
)

// MapToStruct map 转为 struct
func MapToStruct(param map[string]interface{}, s interface{}) (err error) {
	sType := reflect.TypeOf(s)
	if sType.Kind() != reflect.Ptr {
		err = errors.New("参数s请传入struct指针")
		return
	}
	sType = sType.Elem()
	if sType.Kind() != reflect.Struct {
		err = errors.New("参数s请传入struct")
		return
	}
	sValue := reflect.ValueOf(s)
	for i := 0; i < sType.NumField(); i++ {
		field := sType.Field(i)
		fieldKind := field.Type.Kind()
		if fieldKind == reflect.Struct {
			innerFieldType := field.Type
			innerFieldValue := sValue.Elem().Field(i)
			for j := 0; j < innerFieldType.NumField(); j++ {
				innerField := innerFieldType.Field(j)
				innerFieldName := innerField.Name
				if innerField.Type.Kind() == reflect.Ptr {
					if v, ok := param[innerFieldName]; ok {
						vv := reflect.ValueOf(v)
						innerFieldValue.Field(j).Set(vv)
					}
				} else {
					if v, ok := param[innerFieldName]; ok {
						vv := reflect.ValueOf(v)
						innerFieldValue.Field(j).Set(vv)
					}
				}
			}
		} else {
			fieldName := field.Name
			if v, ok := param[fieldName]; ok {
				vv := reflect.ValueOf(v)
				sValue.Elem().Field(i).Set(vv)
			}
		}
	}
	return
}
