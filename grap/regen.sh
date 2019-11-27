#!/usr/bin/env bash

mv resolver.go resolver.old
go run github.com/99designs/gqlgen 

cp resolver.go resolver.go.new
git merge-file resolver.go resolver.go resolver.old