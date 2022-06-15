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

/**
* Add A new group, like linux cmd groupadd
 */
func (o Option) GroupAdd(teamName string, g model.GroupEntry) error {
	dn := ""
	if teamName == "" {
		dn, _ = combineDN(Group, o.LAI.TopDN, g.Name[0])

	} else {
		arr, err := o.SearchAllEntryDNByAttr(Team, "ou", teamName)
		if err != nil {
			return err
		}
		if len(arr) != 1 {
			return fmt.Errorf("[FAIL] %d num of this team", len(arr))
		}
		dn, _ = combineDN(Group, arr[0], g.Name[0])
	}
	attrs, err := Map(g)
	if err != nil {
		return err
	}
	return o.AddEntry(dn, attrs)

}

/**
* like cmd groupdel
 */

func (o Option) GroupDel(groupName string) error {
	arr, err := o.SearchAllEntryDNByAttr(Group, "cn", groupName)
	if err != nil {
		return err
	}
	if len(arr) != 1 {
		return fmt.Errorf("bad dn number %d", len(arr))
	}
	return o.DeleteEntry(arr[0])
}

/**
* like groupmems ,Add mem or del mem
 */

func (o Option) GroupMems(groupName string, mems []string, attrOP int) error {
	arr, err := o.SearchAllEntryDNByAttr(Group, "cn", groupName)
	if err != nil {
		return err
	}
	if len(arr) != 1 {
		return fmt.Errorf("bad dn number %d", len(arr))
	}
	dn := arr[0]
	return o.ModifyEntryAttr(dn, []model.AttrVal{
		{AttrOP: attrOP, Attr: "memberUid", Val: mems},
	})
}

/**
* like groupmod
 */
func (o Option) GroupMod(groupName string, gidNumber string) error {
	arr, err := o.SearchAllEntryDNByAttr(Group, "cn", groupName)
	if err != nil {
		return err
	}
	if len(arr) != 1 {
		return fmt.Errorf("bad dn number %d", len(arr))
	}
	dn := arr[0]
	return o.ModifyEntryAttr(dn, []model.AttrVal{
		{AttrOP: Rep, Attr: "gidNumber", Val: []string{gidNumber}},
	})
}

func NewGroupEntry() model.GroupEntry {
	return model.GroupEntry{
		ObjectClass: defaultLdapOC[Group],
	}
}
