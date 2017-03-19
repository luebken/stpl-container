# Stpl

## Install

clone into $GOPATH/src/github.com/luebken/stpl

    $ make builddocker
    $ docker run luebken/stpl

## Run

### CLI

    $ go run cmd/stplcli/main.go
    $ go install github.com/luebken/stpl/cmd/stplcli

### Server

    $ go run cmd/stplsrv/main.go
    $ go install github.com/luebken/stpl/cmd/stplsrv



## Notes

* Project structure based on https://peter.bourgon.org/go-best-practices-2016/
* Docker build based on https://developer.atlassian.com/blog/2015/07/osx-static-golang-binaries-with-docker/