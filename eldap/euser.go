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
