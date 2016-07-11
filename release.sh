#! /bin/sh


cd $(dirname $0)

ver=${1-$(grep version cli/version.go|cut -f2 -d\")}
[ -z "$ver" ] && exit 1
set -x

# Linux
for arch in amd64 386 arm; do
  GOOS=linux GOARCH=$arch go build -o p2pub-$ver.linux-$arch ./cli
done

# Windows
for arch in amd64 386 ; do
  GOOS=windows GOARCH=$arch go build -o p2pub-$ver.windows-$arch.exe ./cli
done

# Mac
GOOS=darwin GOARCH=amd64 go build -o p2pub-$ver.darwin-amd64 ./cli
