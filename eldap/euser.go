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
	"ela/model"
	"fmt"
)

var UserKVMap = map[string]string{
	"home": "homeDirectory",
}

/**
* like cmd useradd
 */
func (o Option) UserAdd(UI model.UserInfo) error {
	SuperDN := ""
	if UI.TeamName == "" {
		SuperDN = o.LAI.TopDN
	} else {
		arr, err := o.SearchAllEntryDNByAttr(Team, "ou", UI.TeamName)
		if err != nil {
			return err
		}
		if len(arr) != 1 {
			return fmt.Errorf("bad dn number %d", len(arr))
		}
		SuperDN = arr[0]
	}
	return o.AddEntryBYKindDN(SuperDN, model.EntryInfo{Kind: User, UI: UI})
}

/**
like cmd userdel
*/
func (o Option) UserDel(UserName string) error {
	arr, err := o.SearchAllEntryDNByAttr(User, "uid", UserName)
	if err != nil {
		return err
	}
	if len(arr) != 1 {
		return fmt.Errorf("bad dn number %d", len(arr))
	}
	DN := arr[0]
	return o.DeleteEntry(DN)
}
