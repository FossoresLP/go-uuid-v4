package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
)

// UUID represents a Universal Unique Identifier as an array containing 16 bytes
type UUID [16]byte

// Parse parses a string as a UUID returning either the resulting UUID or an error
func Parse(str string) (uuid UUID, err error) {
	if len(str) != 36 {
		return uuid, fmt.Errorf("Invalid length for UUID: %d", len(str))
	}
	if str[8] != '-' || str[13] != '-' || str[18] != '-' || str[23] != '-' {
		return uuid, errors.New("UUID format invalid")
	}
	part1, err := hex.DecodeString(str[:8])
	if err != nil {
		return
	}
	part2, err := hex.DecodeString(str[9:13])
	if err != nil {
		return
	}
	part3, err := hex.DecodeString(str[14:18])
	if err != nil {
		return
	}
	part4, err := hex.DecodeString(str[19:23])
	if err != nil {
		return
	}
	part5, err := hex.DecodeString(str[24:36])
	if err != nil {
		return
	}
	copy(uuid[:4], part1)
	copy(uuid[4:6], part2)
	copy(uuid[6:8], part3)
	copy(uuid[8:10], part4)
	copy(uuid[10:16], part5)
	return
}

// ParseBytes parses a byte slice and returns the contained UUID or an error
func ParseBytes(bytes []byte) (uuid UUID, err error) {
	if len(bytes) != 16 {
		return uuid, fmt.Errorf("Invalid length for binary UUID: %d", len(bytes))
	}
	copy(uuid[:], bytes)
	return
}

// New generates a new UUID v4 using a secure random number generator chosen by GoLang
func New() (uuid UUID, err error) {
	_, err = rand.Read(uuid[:])
	if err != nil {
		return
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant 10
	return
}

// ToString returns the string representation of a UUID
func (uuid *UUID) String() (out string) {
	out += hex.EncodeToString(uuid[:4])
	out += "-"
	out += hex.EncodeToString(uuid[4:6])
	out += "-"
	out += hex.EncodeToString(uuid[6:8])
	out += "-"
	out += hex.EncodeToString(uuid[8:10])
	out += "-"
	out += hex.EncodeToString(uuid[10:16])
	return
}

// NewString generates a new UUID v4 and converts it to a string immediately
func NewString() (string, error) {
	uuid, err := New()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

// IsEmpty returns if the UUID contains only zeros and is therefore empty and invalid
func (uuid *UUID) IsEmpty() bool {
	return *uuid == [16]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
}

// MarshalText provides encoding.TextMarshaler
func (uuid UUID) MarshalText() ([]byte, error) {
	return []byte(uuid.String()), nil
}

// UnmarshalText provides encoding.TextUnmarshaler
func (uuid *UUID) UnmarshalText(in []byte) error {
	id, err := Parse(string(in))
	if err != nil {
		return err
	}
	*uuid = id
	return nil
}

// MarshalBinary provides encoding.BinaryMarshaler
func (uuid UUID) MarshalBinary() ([]byte, error) {
	return uuid[:], nil
}

// UnmarshalBinary provides encoding.BinaryUnmarshaler
func (uuid *UUID) UnmarshalBinary(in []byte) error {
	if len(in) != 16 {
		return fmt.Errorf("Invalid length for binary UUID: %d", len(in))
	}
	copy(uuid[:], in)
	return nil
}
