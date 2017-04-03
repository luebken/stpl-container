.PHONY: e2etests

# run `make` to see options
.DEFAULT_GOAL := help

go-get: ## get go dependencies
	go get github.com/blang/semver
	go get gopkg.in/yaml.v2

go-run-server: ## runs the server
	go run cmd/stplsrv/main.go

go-test:
	go test github.com/luebken/stpl/pkg/stpl/stacks

e2etests: ## run the e2etests
	e2etests/e2etests.sh

docker-build: ## builds a docker image (luebken/stpl)

	# build builder image which calls `make buildgo`
	docker build --no-cache -t luebken/build-stpl -f ./Dockerfile.build .

	# cp binary to local file 
	docker run -t luebken/build-stpl /bin/true
	docker cp `docker ps -q -n=1`:/stplsrv .
	chmod 755 ./stplsrv

	# build final image
	docker build --tag=docker-registry.stage.engineering.redhat.com/luebken/stpl .

docker-run: docker-build ## runs stpl from docker
	docker run -p 8088:8088 docker-registry.stage.engineering.redhat.com/luebken/stpl

# build go binary
# gets called within docker build
go-build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o stplsrv ./go/src/github.com/luebken/stpl/cmd/stplsrv


# via http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'