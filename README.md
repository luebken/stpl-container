# Stpl

Stpl (short for stapel, loosely German for stack) is a simplistic prototype to test and evaluate stack analysis.

## Getting started

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

### MiniKube

    $ minikube start --vm-driver=xhyve --insecure-registry=docker-registry.stage.engineering.redhat.com
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