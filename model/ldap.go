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
package model

// Make Sure entry type is leaf or not
type EntryBase struct {
	HasSubordinates bool   `json:"hasSub"`
	Kind            int    `json:"kind"`
	Name            string `json:"name"`
	DN              string `json:"dn"`
}

type LDAPAuthInfo struct {
	LDAPHost string
	LDAPPort string
	TopDN    string
	Admin    string
	AdminPW  string
}

type Attrs map[string][]string

type AttrVal struct {
	Attr   string
	Val    []string
	AttrOP int
}

type TeamEntry struct {
	Name             []string `eldap:"ou"`
	Description      []string `eldap:"description"`
	ObjectClass      []string `eldap:"objectClass"`
	AssociatedDomain []string `eldap:"associatedDomain"`
}

type GroupEntry struct {
	Name        []string `eldap:"cn"`
	GidNumber   []string `eldap:"gidNumber"`
	Description []string `eldap:"description"`
	MemberUid   []string `eldap:"memberUid"`
	ObjectClass []string `eldap:"objectClass"`
}

type UserEntry struct {
	Name          []string `eldap:"uid"`
	ObjectClass   []string `eldap:"objectClass"`
	LoginShell    []string `eldap:"loginShell"`
	GidNumber     []string `eldap:"gidNumber"`
	UidNumber     []string `eldap:"uidNumber"`
	HomeDirectory []string `eldap:"homeDirectory"`
	UserPassword  []string `eldap:"userPassword"`
	CN            []string `eldap:"cn"`
}
