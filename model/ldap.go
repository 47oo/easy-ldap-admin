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
