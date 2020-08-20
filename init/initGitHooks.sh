#!/bin/sh

#installing dependency
go get -u golang.org/x/lint/golint
go get golang.org/x/tools/cmd/goimports
go get -u golang.org/x/tools/cmd/gopls

git config core.hooksPath .githooks