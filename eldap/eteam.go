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

// func (o Option) combinationDN(Name string, SuperDN string) string {
// 	return fmt.Sprintf("%s,%s", Name, SuperDN)
// }

/*
	Add a new team by team name
*/
func tranformTE(t model.TeamEntry) []model.AttrVal {
	tmpmap := Map(t)
	Attrs := []model.AttrVal{}
	for k, v := range tmpmap {
		Attrs = append(Attrs, model.AttrVal{Attr: k, Val: v})
	}
	return Attrs
}

func (o Option) TeamAdd(t model.TeamInfo) error {
	arr, err := o.SearchAllEntryDNByAttr(Team, "ou", t.Name)
	if err != nil {
		return err
	}
	if len(arr) != 0 {
		return fmt.Errorf("[FAIL] we find  num %d name team,this version only support one from whole tree", len(arr))
	}
	attrs := Map(t)
	dn, err := combineDN(Team, o.LAI.TopDN, t.Name)
	if err != nil {
		return err
	}
	return o.AddEntry(dn, attrs)
}

/**
*  Update team desc
 */
func (o Option) TeamDescUpdate(TI model.TeamInfo) error {
	arr, err := o.SearchAllEntryDNByAttr(Team, "ou", TI.Name)
	if err != nil {
		return err
	}
	if len(arr) != 1 {
		return fmt.Errorf("[FAIL] we find  num %d name team,this version only support one", len(arr))
	}
	DN := arr[0]
	return o.ModifyEntryAttr(DN, []model.AttrVal{
		{Attr: "description", Val: []string{TI.Description}, AttrOP: Rep},
	})
}

/**
* Del team and the team must has no leaf ,or delete err
 */

func (o Option) TeamDelete(TeamName string) error {
	arr, err := o.SearchAllEntryDNByAttr(Team, "ou", TeamName)
	if err != nil {
		return err
	}
	if len(arr) != 1 {
		return fmt.Errorf("[FAIL] we find  num %d name team,this version only support one", len(arr))
	}
	DN := arr[0]
	return o.DeleteEntry(DN)
}

func CreateNewTeamEntry() model.TeamEntry {
	return model.TeamEntry{ObjectClass: defaultLdapOC[Team]}
}
