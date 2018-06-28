Go UUIDv4
=========

[![CircleCI branch](https://img.shields.io/circleci/project/github/FossoresLP/go-uuid-v4/master.svg?style=flat-square)](https://circleci.com/gh/FossoresLP/go-uuid-v4)

This package contains a *minimal* implementation of UUIDv4 in Go.

Consider using either [Google's implementation](https://github.com/google/uuid) or [this one by Satori](https://github.com/satori/go.uuid) if you need additional features.

*This package can only generate version 4 UUIDs as defined in [RFC 4122](http://tools.ietf.org/html/rfc4122)*

Create a new UUID: `uuid.New() (UUID, error)`

Convert an UUID to a string: `UUID.String() string`

Do both in one step: `uuid.NewString() (string, error)`

Convert a string to an UUID: `uuid.Parse(string) (UUID, error)`

Convert a byte slice to an UUID: `uuid.ParseBytes([]byte) (UUID, error)`

Check if UUID contains only zeros: `UUID.IsEmpty() bool`

Both `encoding.Text(Un)Marshaler` and `encoding.Binary(Un)Marshaler` are supported to ensure compatibility with databases and data exchange formats.
