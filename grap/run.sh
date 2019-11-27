#!/usr/bin/env bash

export PORT=8888
export DB_URL="user=postgres password=1234 dbname=chatg sslmode=disable"
export DB_URL_CREATE="user=postgres password=1234 sslmode=disable"

go run ./server/server.go
