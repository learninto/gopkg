package structlib

import (
	"kingstar-go/commons/namedlib"
	"reflect"
)

// https://www.sozhidao.com/articles/260
// ToMap 转map
func ToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)  //使用反射,k根据typeof拿到field
	v := reflect.ValueOf(obj) //v根据valueof拿到interface

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		fileName := t.Field(i).Name

		fileName = namedlib.UnMarshal(fileName)

		data[fileName] = v.Field(i).Interface()
	}
	return data
}

// FindJsonTag 查询json标签
func FindJsonTag(param interface{}) (res []string) {
	t := reflect.TypeOf(param).Elem()

	for i := 0; i < t.NumField(); i++ {
		jStr := t.Field(i).Tag.Get("json")
		res = append(res, jStr)
	}

	return res
}
