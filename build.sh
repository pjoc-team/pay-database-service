#!/usr/bin/env bash
export GO111MODULE=on
CGO_ENABLED=1 GOOS=linux go build -o ./bin/main .
