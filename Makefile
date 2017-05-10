.PHONY: e2etests docs

# run `make` to see options
.DEFAULT_GOAL := help

export STPL_REDIS_PORT = 6379

go-get: # get go dependencies
	go get github.com/blang/semver
	go get gopkg.in/yaml.v2
	go get github.com/Sirupsen/logrus
	go get github.com/go-redis/redis


go-run: ## runs the server
	go run cmd/stplsrv/main.go

go-test:
	go test github.com/luebken/stpl/pkg/stpl/stacks

e2etests: ## run the e2etests
	e2etests/e2etests.sh

docker-build: ## builds a docker image (luebken/stpl)

	# build builder image which calls `make buildgo`
	docker build -t luebken/build-stpl -f ./Dockerfile.build .

	# cp binary to local file 
	docker run -t luebken/build-stpl /bin/true
	docker cp `docker ps -q -n=1`:/stplsrv .
	chmod 755 ./stplsrv

	# build final image
	docker build --tag=luebken/stpl .

docker-run: ## runs stpl from docker
	docker run -p 8088:8088  --link stpl-redis:redis --env LIBRARIES_IO_API_KEY=$(LIBRARIES_IO_API_KEY) --env REDIS_URL=redis://redis:6379 luebken/stpl

docker-run-redis: ## runs the redis db
	docker run -p 6379:6379 --rm -v /Users/mdl/workspace/golang/src/github.com/luebken/stpl/data:/data --name stpl-redis redis:3 redis-server --appendonly yes


# get backup: docker exec -it ce65835e8fc3 cat appendonly.aof > appendonly.aof
# docker run -it --link stpl-redis:redis --rm redis redis-cli -h redis -p 6379

docker-push:
# docker login -p ${DOCKER_REGISTRY_TOKEN}  -e unused -u unused docker-registry.stage.engineering.redhat.com
	docker push luebken/stpl

# build go binary
# gets called within docker build
go-build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o stplsrv ./go/src/github.com/luebken/stpl/cmd/stplsrv

# API docs based on https://github.com/yvasiyarov/swagger
docs: ##create docs
	swagger -format=markdown -apiPackage="github.com/luebken/stpl/pkg/stpl/" -mainApiFile="github.com/luebken/stpl/cmd/stplsrv/main.go" -output=docs/API.md

# via http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'