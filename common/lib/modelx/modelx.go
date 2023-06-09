package modelx

import (
	"fmt"
	"reflect"
	"strings"
)

func RawFieldNames(in interface{}, ignore ...string) []interface{} {
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}

	strIgnore := strings.Join(ignore, ",")

	typ := v.Type()
	// out := make([]interface{}, v.NumField())
	var out []interface{}
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		typFi := typ.Field(i).Tag.Get("db")
		if strIgnore != "" && strings.Contains(strIgnore, typFi) {
			continue
		}
		fi := v.Field(i)
		out = append(out, fi.Interface())
	}

	return out
}
