#!/bin/sh

godoc -http=:6060
start http://localhost:6060/pkg/gosemantic/set/
