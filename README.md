# Stpl

Stpl (short for stapel, loosely German for stack) is a simplistic prototype to test and evaluate stack recommendations.

## Getting started

### Prerequisites

    * Get an API Key from https://libraries.io/ and set it as LIBRARIES_IO_API_KEY

### Run pre-build Docker Image

    # run container:
    $ docker run -p 8088:8088 docker-registry.stage.engineering.redhat.com/luebken/stpl

    # run e2etests:
    $ git clone git@gitlab.cee.redhat.com:mluebken/stpl.git
    $ cd stpl
    $ make e2etests

### Run from source
    # clone into $GOPATH/src/github.com/luebken/stpl
    $ make go-get
    $ make go-run-server
    # open new terminal
    $ make e2etests

## Deployment

### Docker Registry

    $ make docker-build
    $ make docker-push

### MiniShift

    $ oc new-project mdl
    $ oc create -f k8s/deployment-redis.yaml && oc create -f k8s/service-redis.yaml
    $ oc create -f k8s/deployment.yaml && oc create -f k8s/service.yaml
    $ oc create -f k8s/route.yaml
    $ curl -X POST -d @e2etests/example-1-pom.xml $(oc get route -o=jsonpath='{.items[0].spec.host}')/analytics

### online
    # get online accound and login
    $ oc new-project mdl

### MiniKube

    $ minikube start --vm-driver=xhyve --insecure-registry=docker-registry.stage.engineering.redhat.com
    $ eval $(minikube docker-env)
    $ make builddocker
    $ kubectl create -f k8s/deployment.yaml
    $ kubectl create -f k8s/service.yaml
    $ curl -X POST -d @e2etests/example-1-pom.xml $(minikube service stpl --url)/analytics

## Notes

* Project structure based on https://peter.bourgon.org/go-best-practices-2016/
* Docker build based on https://developer.atlassian.com/blog/2015/07/osx-static-golang-binaries-with-docker/