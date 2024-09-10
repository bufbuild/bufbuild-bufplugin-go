bufplugin-go
==============

[![Build](https://github.com/bufbuild/bufplugin-go/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/bufbuild/bufplugin-go/actions/workflows/ci.yaml)
[![Report Card](https://goreportcard.com/badge/buf.build/go/bufplugin)](https://goreportcard.com/report/buf.build/go/bufplugin)
[![GoDoc](https://pkg.go.dev/badge/buf.build/go/bufplugin.svg)](https://pkg.go.dev/buf.build/go/bufplugin)
[![Slack](https://img.shields.io/badge/slack-buf-%23e01563)](https://buf.build/links/slack)

This is the Golang SDK for [bufplugin](https://github.com/bufbuild/bufplugin).

```go
import "buf.build/go/bufplugin/check"
```

This is very early, but see the [example](check/internal/example) for how this works in practice.

## Status: Alpha

Bufplugin is as early as it gets - [buf](https://github.com/bufbuild/buf) doesn't actually support
plugins yet! We're publishing this publicly to get early feedback as we approach stability.

## Legal

Offered under the [Apache 2 license](https://github.com/bufbuild/bufplugin-go/blob/main/LICENSE).
