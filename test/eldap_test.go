package test

import (
	"ela/eldap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	tests := []map[string][]string{
		{
			"ou":               []string{"47oo"},
			"associatedDomain": []string{"dc=nudt,dc=org"},
			"objectClass":      []string{"top", "organizationalUnit", "domainRelatedObject"},
		},
	}
	nt := eldap.CreateNewTeamEntry()
	nt.Name = append(nt.Name, "47oo")
	nt.AssociatedDomain = append(nt.AssociatedDomain, "dc=nudt,dc=org")
	mapnt := eldap.Map(nt)
	t.Logf("%v\n", mapnt)
	for _, tt := range tests {
		assert.Equal(t, tt, mapnt)
	}
}
