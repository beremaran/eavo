#!/usr/bin/env bash

PKG_LIST=$(go list ./... | grep -v /vendor/)

mkdir -p cover;

for package in ${PKG_LIST}; do
    COV_PATH=$(echo "${package}" | tr / _)
    COV_PATH="cover/${COV_PATH}.cov"
    go test -covermode=count -coverprofile ${COV_PATH} "$package" ;
done

echo "mode: count" >> cover/coverage.cov
tail -q -n +2 cover/*.cov >> cover/coverage.cov
go tool cover -func=cover/coverage.cov
rm -rf cover;