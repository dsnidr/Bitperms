package bitperms

import (
	"strconv"
)

// Permissions is a helper type for holding a map of PermissionValues.
type Permissions map[string]PermissionValue

// PermissionValue is an int64 with some additional functions attached.
// Note that since PermissionValue is built on a 64 bit integer, you should support a maximum of
// 64 different permission flags.
type PermissionValue int64

// Serialize returns the string version of the PermissionValue it was called on.
func (pv PermissionValue) Serialize() string {
	return strconv.FormatInt(int64(pv), 10)
}

// Deserialize creates and returns a PermissionValue from a valid string.
// If the provided string can not be parsed into an int64, an error is returned.
func Deserialize(serialized string) (PermissionValue, error) {
	res, err := strconv.ParseInt(serialized, 10, 64)
	if err != nil {
		return 0, err
	}

	return PermissionValue(res), nil
}

// HasFlag takes in an int64 permission flag and runs a logical AND operation on the PermissionValue and the flag to
// check if the target flags are set. If they are, true is returned.
//
// To check against multiple flags, you can either use the HasFlags function or do an OR operation between all
// flags you want to test as the flag argument.
func (pv PermissionValue) HasFlag(flag int64) bool {
	res := int64(pv) & flag

	return res == flag
}

// HasFlags takes in multiple int64 permission flags and combines them using bitwise OR operations before calling
// HasFlag using the constructed flag.
//
// See HasFlag for more information on how permission checking works.
func (pv PermissionValue) HasFlags(flags ...int64) bool {
	flag := int64(0)

	for _, f := range flags {
		flag = flag | f
	}

	return pv.HasFlag(flag)
}
