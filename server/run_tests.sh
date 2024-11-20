#!/usr/bin/env sh

docker compose build server redis db
docker compose up -d redis db
export $(grep -v '^#' ../.env | xargs)
export POSTGRES_HOST=127.0.0.1 REDIS_HOST=127.0.0.1
go mod tidy && ginkgo -coverpkg=./... -r -coverprofile=cover.out ./... && go tool cover -func=cover.out &&go tool cover -html cover.out -o cover.html
docker compose stop redis db
