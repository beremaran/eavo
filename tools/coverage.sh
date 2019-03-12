#!/usr/bin/env bash

PKG_LIST=$(go list ./... | grep -v /vendor/)

mkdir -p cover;

for package in ${PKG_LIST}; do
    go test -covermode=count -coverprofile "cover/${package##*/}.cov" "$package" ;
done

echo "mode: count\n" >> cover/coverage.cov
tail -q -n +2 cover/*.cov >> cover/coverage.cov
go tool cover -func=cover/coverage.cov
rm -rf cover;