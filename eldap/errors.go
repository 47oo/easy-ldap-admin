package eldap

import (
	"errors"
)

var ErrGroupAdd = errors.New("[FAIL] group name exist")
var ErrGroupDel = errors.New("[FAIL] group name not found  or more than one same name")
var ErrGroupMod = errors.New("[FAIL] modify attr for group fail")

var ErrTeamAdd = errors.New("[FAIL] team name exist")
var ErrTeamDel = errors.New("[FAIL] team name not found or more than one same name")
var ErrTeamMod = errors.New("[FAIL] modify attr for team fail")

var ErrUserAdd = errors.New("[FAIL] user name exist")
var ErrUserDel = errors.New("[FAIL] user name not dount or more than one same name")
var ErrUserMod = errors.New("[FAIL] modify attr for user fail")
