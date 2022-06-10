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
type EntryBaseInfo struct {
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

type EntryInfo struct {
	TI   TeamInfo
	GI   GroupInfo
	UI   UserInfo
	Kind int //User Group Team must exist
}
type TeamInfo struct {
	Name        string
	Description string
}

type GroupInfo struct {
	Name        string
	GidNumber   string
	Description string
	MemberUid   []string
	TeamName    string
}

type UserInfo struct {
	LoginShell    string
	GidNumber     string
	UidNumber     string
	Name          string
	HomeDirectory string
	UserPassword  string
	TeamName      string
}

type AttrVal struct {
	Attr   string
	Val    []string
	AttrOP int
}
