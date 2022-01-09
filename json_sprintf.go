package json_sprintf

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func Exec(target interface{}) string {
	format := getFormat(target)

	structAndType := getStructAndType(target)
	var a string
	for _, v := range structAndType {
		a += "\n" + v.(string)
	}

	var result = "fmt.Sprintf(`"
	result += format + "`,"
	result += a
	result += "\n" + ")"

	fmt.Println(result)
	return result
}

func getFormat(target interface{}) string {
	j, _ := json.Marshal(target)
	sj := string(j)

	replaced := strings.ReplaceAll(sj, ":\"\"", ":\"%s\"")   // string
	replaced = strings.ReplaceAll(replaced, ":0", ":%d")     // int, uint, float
	replaced = strings.ReplaceAll(replaced, ":false", ":%t") // bool
	replaced = strings.ReplaceAll(replaced, ":null", ":%v")  // pointer

	return replaced
}

func getStructAndType(target interface{}) []interface{} {
	var vs []interface{}
	v := reflect.Indirect(reflect.ValueOf(target))
	t := v.Type()

	switch t.Kind() {
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			ft := t.Field(i)
			fv := v.FieldByName(ft.Name)

			if ft.Type.Kind() == reflect.Struct {
				sl := getStructAndType(fv.Interface())
				vs = append(vs, sl...)
			} else if ft.Type.Kind() == reflect.Ptr && ft.Type.Elem().Kind() == reflect.Struct {
				sl := getStructAndType(fv.Interface())
				vs = append(vs, sl...)
			} else if ft.Type.Kind() == reflect.Slice && ft.Type.Elem().Kind() == reflect.Struct {
				sl := getStructAndType(fv.Interface())
				vs = append(vs, sl...)
			} else {
				vs = append(vs, t.Name()+"."+ft.Name+",//"+t.Name()+"."+ft.Name)
			}

		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			e := v.Index(i)
			sl := getStructAndType(e.Addr().Interface())
			vs = append(vs, sl...)
		}
	}
	return vs
}
