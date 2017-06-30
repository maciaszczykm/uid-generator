package main

import "encoding/hex"

// UID is an unique identifier.
type UID []byte

// ToString encodes UID to string.
func (u UID) ToString() string {
	return hex.EncodeToString(u)
}

// FromString decodes encoded string to UID type.
// TODO Validate length.
func FromString(encoded string) UID {
	uid, _ := hex.DecodeString(encoded)
	return uid
}