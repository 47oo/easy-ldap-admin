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
	"ela/secret"
	"fmt"
	"os"

	"github.com/go-ldap/ldap/v3"
	"github.com/howeyc/gopass"
	"github.com/spf13/viper"
)

const (
	User  int = iota // Entry has posixAccount
	Group            // Entry has posixGroup
	Team             // Entry has organizationalUnit
	Unknown
)

const (
	Add int = iota // Add Entry Attr
	Del            // Del Entry Attr
	Rep            //Replace Entry Attr If Attr not one ,you can use del ,then add ,rep will replace all attr about this name
)

var defaultKindOC = [...]string{
	"posixAccount",
	"posixGroup",
	"organizationalUnit",
}
var defaultNameAttr = [...]string{
	"uid",
	"cn",
	"ou",
}

var defaultLdapOC = [...][]string{
	{"posixAccount", "top", "shadowAccount", "account"},
	{"posixGroup", "top"},
	{"top", "organizationalUnit", "domainRelatedObject"},
}

// Make Sure entry type is leaf or not
type Option struct {
	LAI model.LDAPAuthInfo
}

/**
connect to ldap server like slapd and manage by your admin account
*/
func (o Option) ldapConn() (*ldap.Conn, error) {
	conn, err := ldap.DialURL(fmt.Sprintf("ldap://%s:%s", o.LAI.LDAPHost, o.LAI.LDAPPort))
	if err != nil {
		return nil, err
	}
	err = conn.Bind(fmt.Sprintf("cn=%s,%s", o.LAI.Admin, o.LAI.TopDN), o.LAI.AdminPW)
	if err != nil {
		return nil, err
	}
	return conn, nil

}

func (o Option) TypeIs(DN string) (*model.EntryBaseInfo, error) {
	conn, err := o.ldapConn()

	if err != nil {
		return nil, err
	}
	defer conn.Close()
	nsr := ldap.NewSearchRequest(DN, ldap.ScopeBaseObject, ldap.NeverDerefAliases,
		0, 0, false, "(objectclass=*)", []string{"hasSubordinates", "objectClass"}, nil)
	sr, err := conn.Search(nsr)
	if err != nil {
		return nil, err
	}
	EBI := model.EntryBaseInfo{Kind: Unknown, HasSubordinates: false}
	sr.Print()
	for _, entry := range sr.Entries {
		for _, v := range entry.GetAttributeValues("objectClass") {
			if v == defaultKindOC[Group] {
				EBI.Kind = Group
			} else if v == defaultKindOC[User] {
				EBI.Kind = User
			} else if v == defaultKindOC[Team] {
				EBI.Kind = Team
			}
		}
		if entry.GetAttributeValue("hasSubordinates") == "TRUE" {
			EBI.HasSubordinates = true
		}
	}
	return &EBI, nil
}
func (o Option) SearchAllEntryDNByAttr(Kind int, Attr string, Val string) ([]string, error) {
	conn, err := o.ldapConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	nsr := ldap.NewSearchRequest(o.LAI.TopDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases,
		0, 0, false, "", []string{"uid", "ou", "cn", "hasSubordinates"}, nil)
	nsr.Filter = fmt.Sprintf("(&(objectclass=%s)(%s=%s))", defaultKindOC[Kind], Attr, Val)
	res, err := conn.Search(nsr)
	if err != nil {
		return nil, err
	}
	DNArr := make([]string, 0)
	for _, v := range res.Entries {
		DNArr = append(DNArr, v.DN)
	}
	return DNArr, nil
}

/**
* DN is the Base Search location,and Kind support User,Group,Team range 0-2
 */
func (o Option) SearchAllEntryByKindDN(DN string, Kind int) ([]model.EntryBaseInfo, error) {
	conn, err := o.ldapConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	nsr := ldap.NewSearchRequest(DN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases,
		0, 0, false, "", []string{"uid", "ou", "cn", "hasSubordinates"}, nil)
	if Kind == Team {
		nsr.Filter = fmt.Sprintf("(objectclass=%s)", defaultKindOC[Team])
	} else if Kind == Group {
		nsr.Filter = fmt.Sprintf("(objectclass=%s)", defaultKindOC[Group])
	} else if Kind == User {
		nsr.Filter = fmt.Sprintf("(objectclass=%s)", defaultKindOC[User])
	} else {
		return nil, fmt.Errorf("unknown kind: %v", Kind)
	}

	res, err := conn.Search(nsr)
	if err != nil {
		return nil, err
	}
	res.Print()
	EBIArr := make([]model.EntryBaseInfo, 0)
	for _, entry := range res.Entries {
		ebi := model.EntryBaseInfo{}
		ebi.HasSubordinates = false
		if entry.GetAttributeValue("hasSubordinates") == "TRUE" {
			ebi.HasSubordinates = true
		}
		ebi.Name = entry.GetAttributeValue(defaultNameAttr[Kind])
		ebi.Kind = Team
		EBIArr = append(EBIArr, ebi)
	}
	return EBIArr, nil
}

/**
Only return one layer entry by the domain you input
*/
func (o Option) ShowBaseInfoScopeOne(DN string) ([]model.EntryBaseInfo, error) {

	conn, err := o.ldapConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	lnsr := ldap.NewSearchRequest(DN, ldap.ScopeSingleLevel,
		ldap.NeverDerefAliases, 0, 0, false, "(objectclass=*)", []string{"uid", "dn", "ou", "cn", "objectClass", "hasSubordinates"}, nil)
	res, err := conn.Search(lnsr)
	if err != nil {
		return nil, err
	}
	EBIArr := make([]model.EntryBaseInfo, 0)
	for _, entry := range res.Entries {
		ebi := model.EntryBaseInfo{Kind: Unknown, DN: entry.DN}

		for _, v := range entry.GetAttributeValues("objectClass") {
			if v == defaultKindOC[Group] {
				ebi.HasSubordinates = false
				ebi.Kind = Group
				ebi.Name = entry.GetAttributeValue("cn")
				break
			} else if v == defaultKindOC[User] {
				ebi.HasSubordinates = false
				ebi.Kind = User
				ebi.Name = entry.GetAttributeValue("uid")
				break
			} else if v == defaultKindOC[Team] {
				if entry.GetAttributeValue("hasSubordinates") == "TRUE" {

					ebi.HasSubordinates = true
				} else {
					ebi.HasSubordinates = false
				}
				ebi.Kind = Team
				ebi.Name = entry.GetAttributeValue("ou")
			}

		}
		EBIArr = append(EBIArr, ebi)
	}
	return EBIArr, nil
}

/*
* Only support User Group Team
 */
func (o Option) AddEntryBYKindDN(SuperDN string, EI model.EntryInfo) error {
	conn, err := o.ldapConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	nar := ldap.NewAddRequest("", nil)
	if EI.Kind == Team {
		nar.DN = fmt.Sprintf("ou=%s,%s", EI.TI.Name, SuperDN)
		nar.Attribute("objectClass", defaultLdapOC[Team])
		nar.Attribute("ou", []string{EI.TI.Name})
		nar.Attribute("associatedDomain", []string{SuperDN})
		nar.Attribute("description", []string{EI.TI.Description})
	} else if EI.Kind == Group {
		nar.DN = fmt.Sprintf("cn=%s,%s", EI.GI.Name, SuperDN)
		nar.Attribute("objectClass", defaultLdapOC[Group])
		nar.Attribute("cn", []string{EI.GI.Name})
		nar.Attribute("gidNumber", []string{EI.GI.GidNumber})
		nar.Attribute("description", []string{EI.GI.Description})
	} else if EI.Kind == User {
		nar.DN = fmt.Sprintf("uid=%s,%s", EI.UI.Name, SuperDN)
		nar.Attribute("objectClass", defaultLdapOC[User])
		nar.Attribute("cn", []string{EI.UI.Name})
		nar.Attribute("uid", []string{EI.UI.Name})            // user username
		nar.Attribute("uidNumber", []string{EI.UI.UidNumber}) //user uid
		nar.Attribute("gidNumber", []string{EI.UI.GidNumber}) // This is primary group
		nar.Attribute("homeDirectory", []string{EI.UI.HomeDirectory})
		nar.Attribute("userPassword", []string{EI.UI.UserPassword})
	} else {
		return fmt.Errorf("unknow kind: %d", EI.Kind)
	}
	return conn.Add(nar)
}

func (o Option) ModifyEntryAttr(DN string, Arr []model.AttrVal) error {
	conn, err := o.ldapConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	nmr := ldap.NewModifyRequest(DN, nil)
	for _, v := range Arr {
		if v.AttrOP == Add {
			nmr.Add(v.Attr, v.Val)
		} else if v.AttrOP == Del {
			nmr.Delete(v.Attr, v.Val)
		} else if v.AttrOP == Rep {
			nmr.Replace(v.Attr, v.Val)
		} else {
			return fmt.Errorf("unknown OP %d", v.AttrOP)
		}
	}
	return conn.Modify(nmr)
}

func (o Option) DeleteEntry(DN string) error {
	conn, err := o.ldapConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	ndr := ldap.NewDelRequest(DN, nil)
	return conn.Del(ndr)
}

func NewOption() Option {
	pwd, _ := secret.DecryptAES([]byte(viper.GetString("default.adminpw")), secret.KEY)
	if string(pwd) == "NO" {
		pass, _ := gopass.GetPasswdPrompt("enter admin password: ", true, os.Stdin, os.Stdout)
		pwd = pass
	}

	return Option{
		LAI: model.LDAPAuthInfo{
			LDAPHost: viper.GetString("default.ldaphost"),
			LDAPPort: viper.GetString("default.ldapport"),
			Admin:    viper.GetString("default.admin"),
			AdminPW:  string(pwd),
			TopDN:    viper.GetString("default.topdn"),
		},
	}
}
