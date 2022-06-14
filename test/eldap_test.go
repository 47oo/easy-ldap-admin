package test

import (
	"ela/eldap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	tests := []eldap.Attrs{
		{
			"ou":               []string{"47oo"},
			"associatedDomain": []string{"dc=nudt,dc=org"},
			"objectClass":      []string{"top", "organizationalUnit", "domainRelatedObject"},
		},
	}
	nt := eldap.NewTeamEntry()
	nt.Name = append(nt.Name, "47oo")
	nt.AssociatedDomain = append(nt.AssociatedDomain, "dc=nudt,dc=org")
	mapnt := eldap.Map(nt)
	t.Logf("%v\n", mapnt)
	for _, tt := range tests {
		assert.Equal(t, tt, mapnt)
	}
}
