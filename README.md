Go UUIDv4
=========

[![CircleCI](https://img.shields.io/circleci/project/github/FossoresLP/go-uuid-v4/master.svg?style=flat-square)](https://circleci.com/gh/FossoresLP/go-uuid-v4)
[![Coveralls](https://img.shields.io/coveralls/github/FossoresLP/go-uuid-v4/master.svg?style=flat-square)](https://coveralls.io/github/FossoresLP/go-uuid-v4)
[![Codacy](https://img.shields.io/codacy/grade/146eef6b23314b74a0a3a47d106a388d.svg?style=flat-square)](https://www.codacy.com/app/FossoresLP/go-uuid-v4)
[![Licensed under: Boost Software License](https://img.shields.io/badge/style-BSL--1.0-red.svg?longCache=true&style=flat-square&label=License)](https://github.com/FossoresLP/go-uuid-v4/blob/master/LICENSE.md)
[![GoDoc](https://img.shields.io/badge/style-reference-blue.svg?longCache=true&style=flat-square&label=GoDoc)](https://godoc.org/github.com/FossoresLP/go-uuid-v4)

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
