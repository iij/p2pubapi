#! /bin/sh

cd $(dirname $0)
#go generate ./protocol
go build -o p2pub ./cli
( cd example
  for i in *.go ; do
    [ $i = "example.go" ] && continue
    go build $i
  done
)
