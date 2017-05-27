#!/bin/bash
CGO_ENABLE=1 GOOS=linux GOARCH=amd64 go build .