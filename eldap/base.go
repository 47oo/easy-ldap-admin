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

func (o Option) TypeIs(DN string) (*model.EntryBase, error) {
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
	EBI := model.EntryBase{Kind: Unknown, HasSubordinates: false}
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

/**
* SuperDN like dc=ldap,dc=org  Name like cn,ou,uid,This will return
* Return Example: cn=<Name>,dc=ldap,dc=org
 */
func combineDN(Kind int, SuperDN string, Name string) (string, error) {
	if Kind == Team {
		return fmt.Sprintf("%s=%s,%s", defaultNameAttr[Team], Name, SuperDN), nil
	}
	if Kind == Group {
		return fmt.Sprintf("%s=%s,%s", defaultNameAttr[Group], Name, SuperDN), nil
	}
	if Kind == User {
		return fmt.Sprintf("%s=%s,%s", defaultNameAttr[User], Name, SuperDN), nil
	}
	return "", fmt.Errorf("unknon kind number %d", Kind)
}

// only support one attr
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
func (o Option) SearchAllEntryByKindDN(DN string, Kind int) ([]model.EntryBase, error) {
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
	EBIArr := make([]model.EntryBase, 0)
	for _, entry := range res.Entries {
		ebi := model.EntryBase{}
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
func (o Option) ShowBaseInfoScopeOne(dn string) ([]model.EntryBase, error) {

	conn, err := o.ldapConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	lnsr := ldap.NewSearchRequest(dn, ldap.ScopeSingleLevel,
		ldap.NeverDerefAliases, 0, 0, false, "(objectclass=*)", []string{"uid", "dn", "ou", "cn", "objectClass", "hasSubordinates"}, nil)
	res, err := conn.Search(lnsr)
	if err != nil {
		return nil, err
	}
	EBIArr := make([]model.EntryBase, 0)
	for _, entry := range res.Entries {
		ebi := model.EntryBase{Kind: Unknown, DN: entry.DN}

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

func (o Option) AddEntry(dn string, attrs model.Attrs) error {
	conn, err := o.ldapConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	nar := ldap.NewAddRequest(dn, nil)
	for k, v := range attrs {
		nar.Attribute(k, v)
	}
	return conn.Add(nar)
}

func (o Option) ModifyEntryAttr(dn string, arr []model.AttrVal) error {
	conn, err := o.ldapConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	nmr := ldap.NewModifyRequest(dn, nil)
	for _, v := range arr {
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

func (o Option) DeleteEntry(dn string) error {
	conn, err := o.ldapConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	ndr := ldap.NewDelRequest(dn, nil)
	return conn.Del(ndr)
}

func NewOption() Option {
	pwd := secret.EasyDecrypt(viper.GetString("default.adminpw"), secret.KEY)
	if string(pwd) == "NO" {
		pass, _ := gopass.GetPasswdPrompt("enter admin password: ", true, os.Stdin, os.Stdout)
		pwd = string(pass)
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
