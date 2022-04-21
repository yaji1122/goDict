#!/bin/bash

set -e

GO111MODULE=off go get github.com/githubnemo/CompileDaemon

CompileDaemon --build="go build -o ./../goDict ./../cmd/*.go" --command=./../goDict