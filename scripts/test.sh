#!/bin/sh

go test -v -cover -coverprofile=coverage.out -bench=. ./...
go tool cover -html=coverage.out
rm coverage.out
