package eldap

import (
	"reflect"
)

func Map(s interface{}) map[string][]string {
	rt := reflect.TypeOf(s)
	rv := reflect.ValueOf(s)
	mapv := map[string][]string{}
	manlen := rt.NumField()
	for i := 0; i < manlen; i++ {
		ft := rt.Field(i)
		key := ft.Tag.Get("eldap")
		// drop attr if nil or "" or len is zero
		val := rv.Field(i).Interface().([]string)
		if len(val) == 0 {
			continue
		}
		mapv[key] = val
	}
	return mapv
}
