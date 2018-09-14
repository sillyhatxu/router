#!/usr/bin/env bash

echo clean start

cd api_a
go clean
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
cd ..

cd api_b
go clean
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
cd ..

cd api_c
go clean
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
cd ..
echo build end

echo docker-compose build start
docker-compose -f docker-compose.yml build
echo docker-compose build end

docker stack deploy -c docker-compose.yml dev