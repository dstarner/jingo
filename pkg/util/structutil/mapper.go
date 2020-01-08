package structutil

import "reflect"

//ToMap will help you to convert your object from struct to map[string]interface{} based on the tag in its definition
func ToMap(item interface{}, tagName string) map[string]interface{} {

	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get(tagName)
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = ToMap(field, tagName)
			} else {
				res[tag] = field
			}
		}
	}
	return res
}
