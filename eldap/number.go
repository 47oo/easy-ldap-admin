package eldap

import (
	"fmt"
	"strconv"
)

var MinNumber = 10000
var MaxNumber = 100000

func NewGidNumber(min int, max int) (int, error) {
	o := NewOption()
	for i := min; i < max; i++ {
		arr, err := o.SearchAllEntryDNByAttr(Group, "gidNumber", strconv.Itoa(i))
		if err != nil {
			return -1, err
		}
		if len(arr) != 0 {
			continue
		}
		return i, nil
	}
	return -1, fmt.Errorf("not found unique gidNumber in %d and %d", min, max)

}

func NewPrivateGidNumber(min int, max int, gidNumber int) (int, error) {
	o := NewOption()
	arr, err := o.SearchAllEntryDNByAttr(Group, "gidNumber", strconv.Itoa(gidNumber))
	if err != nil {
		return -1, err
	}
	if len(arr) != 0 {
		return NewGidNumber(min, max)
	}
	return gidNumber, nil
}

func NewUidNumber(min int, max int) (int, error) {
	o := NewOption()
	for i := min; i < max; i++ {
		arr, err := o.SearchAllEntryDNByAttr(User, "uidNumber", strconv.Itoa(i))
		if err != nil {
			return -1, err
		}
		if len(arr) != 0 {
			continue
		}
		return i, nil
	}
	return -1, fmt.Errorf("not found unique uidNumber in %d and %d", min, max)

}
