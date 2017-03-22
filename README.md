# Stpl

## Install

clone into $GOPATH/src/github.com/luebken/stpl

    $ make builddocker
    $ docker run luebken/stpl

## Run

### Server

    $ go run cmd/stplsrv/main.go
    $ go install github.com/luebken/stpl/cmd/stplsrv

### Test

    $ go test github.com/luebken/stpl/pkg/stpl/stacks

### MiniKube

    $ minikube start --vm-driver=xhyve
    $ eval $(minikube docker-env)
    $ make builddocker
    $ kubectl create -f k8s/deployment.yaml
    $ kubectl create -f k8s/service.yaml
    $ curl -X POST -d @e2etests/example-1-pom.xml $(minikube service stpl --url)/analytics

### MiniShift

    $ eval $(minishift docker-env)
    $ make builddocker
    $ oc create -f k8s/deployment.yaml
    $ oc create -f k8s/service.yaml
    $ oc create -f k8s/route.yaml
    $ curl -X POST -d @e2etests/example-1-pom.xml $(oc get route -o=jsonpath='{.items[0].spec.host}')/analytics


## Notes

* Project structure based on https://peter.bourgon.org/go-best-practices-2016/
* Docker build based on https://developer.atlassian.com/blog/2015/07/osx-static-golang-binaries-with-docker/