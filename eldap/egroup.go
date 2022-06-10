package eldap

import (
	"ela/model"
	"fmt"
)

/**
* Add A new group, like linux cmd groupadd
 */
func (o Option) GroupAdd(GI model.GroupInfo) error {
	SuperDN := ""
	if GI.TeamName == "" {
		SuperDN = o.LAI.TopDN
	} else {
		arr, err := o.SearchAllEntryDNByAttr(Team, "ou", GI.TeamName)
		if err != nil {
			return err
		}
		if len(arr) != 1 {
			return fmt.Errorf("bad dn number %d", len(arr))
		}
		SuperDN = arr[0]
	}
	return o.AddEntryBYKindDN(SuperDN, model.EntryInfo{Kind: Group, GI: GI})
}

/**
* like cmd groupdel
 */

func (o Option) GroupDel(GroupName string) error {
	arr, err := o.SearchAllEntryDNByAttr(Group, "cn", GroupName)
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

func (o Option) GroupMems(GroupName string, Memes []string, AttrOP int) error {
	arr, err := o.SearchAllEntryDNByAttr(Group, "cn", GroupName)
	if err != nil {
		return err
	}
	if len(arr) != 1 {
		return fmt.Errorf("bad dn number %d", len(arr))
	}
	DN := arr[0]
	return o.ModifyEntryAttr(DN, []model.AttrVal{
		{AttrOP: AttrOP, Attr: "memberUid", Val: Memes},
	})
}

/**
* like groupmod
 */
func (o Option) GroupMod(GroupName string, GidNumber string) error {
	arr, err := o.SearchAllEntryDNByAttr(Group, "cn", GroupName)
	if err != nil {
		return err
	}
	if len(arr) != 1 {
		return fmt.Errorf("bad dn number %d", len(arr))
	}
	DN := arr[0]
	return o.ModifyEntryAttr(DN, []model.AttrVal{
		{AttrOP: Rep, Attr: "gidNumber", Val: []string{GidNumber}},
	})
}
