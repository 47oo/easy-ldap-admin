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
)

/**
* like cmd useradd
 */
func (o Option) UserAdd(teamName string, u model.UserEntry) error {
	dn := ""
	if teamName == "" {
		dn, _ = combineDN(User, o.LAI.TopDN, u.Name[0])
	} else {
		arr, err := o.SearchAllEntryDNByAttr(Team, "ou", teamName)
		if err != nil {
			return err
		}
		if len(arr) != 1 {
			return ErrUserAdd
		}
		dn, _ = combineDN(User, arr[0], u.Name[0])
	}
	attrs, err := Map(u)
	if err != nil {
		return err
	}
	return o.AddEntry(dn, attrs)
}

/**
like cmd userdel
*/
func (o Option) UserDel(userName string) error {
	arr, err := o.SearchAllEntryDNByAttr(User, "uid", userName)
	if err != nil {
		return err
	}
	if len(arr) != 1 {
		return ErrUserDel
	}
	dn := arr[0]
	return o.DeleteEntry(dn)
}

func (o Option) UserMod(u model.UserEntry) error {
	arr, err := o.SearchAllEntryDNByAttr(User, "uid", u.Name[0])
	if err != nil {
		return err
	}
	if len(arr) != 1 {
		return ErrUserMod
	}
	dn := arr[0]
	um, err := Map(u)
	if err != nil {
		return err
	}
	delete(um, "uid")
	attrs := []model.AttrVal{}
	for k, v := range um {
		attrs = append(attrs, model.AttrVal{Attr: k, Val: v, AttrOP: Rep})
	}
	return o.ModifyEntryAttr(dn, attrs)
}

func NewUserEntry() model.UserEntry {
	return model.UserEntry{
		ObjectClass: defaultLdapOC[User],
	}
}
