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
func (o Option) TeamAdd(TI model.TeamInfo) error {
	arr, err := o.SearchAllEntryDNByAttr(Team, "ou", TI.Name)
	if err != nil {
		return err
	}
	if len(arr) != 0 {
		return fmt.Errorf("[FAIL] we find  num %d name team,this version only support one from whole tree", len(arr))
	}

	return o.AddEntryBYKindDN(o.LAI.TopDN, model.EntryInfo{Kind: Team, TI: model.TeamInfo{Name: TI.Name, Description: TI.Description}})
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
