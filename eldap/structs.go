/*
Copyright © 2022 Hao Han <136698493@qq.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package eldap

import (
	"reflect"
)

type Attrs map[string][]string

/**
* Only Used For eldap
 */
func Map(s interface{}) Attrs {
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
