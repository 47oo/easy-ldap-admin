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

/*
*	Add a new team by team name
 */
func (o Option) TeamAdd(t model.TeamEntry) error {
	arr, err := o.SearchAllEntryDNByAttr(Team, "ou", t.Name[0])
	if err != nil {
		return err
	}
	if len(arr) != 0 {
		return fmt.Errorf("[FAIL] we find  num %d name team,this version only support one from whole tree", len(arr))
	}
	t.AssociatedDomain = append(t.AssociatedDomain, o.LAI.TopDN)
	attrs := Map(t)
	dn, _ := combineDN(Team, o.LAI.TopDN, t.Name[0])
	return o.AddEntry(dn, attrs)
}

/**
*  Update team desc
 */
func (o Option) TeamDescUpdate(t model.TeamEntry) error {
	arr, err := o.SearchAllEntryDNByAttr(Team, "ou", t.Name[0])
	if err != nil {
		return err
	}
	if len(arr) != 1 {
		return fmt.Errorf("[FAIL] we find  num %d name team,this version only support one", len(arr))
	}
	DN := arr[0]
	return o.ModifyEntryAttr(DN, []model.AttrVal{
		{Attr: "description", Val: t.Description, AttrOP: Rep},
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

func NewTeamEntry() model.TeamEntry {
	return model.TeamEntry{ObjectClass: defaultLdapOC[Team]}
}
