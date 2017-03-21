default: builddocker

.PHONY: e2etests

setup:
    #go get golang.org/x/oauth2
    #go get golang.org/x/oauth2/jwt
    #go get google.golang.org/api/analytics/v3

# build go binary
# gets called within docker build
buildgo:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o stplsrv ./go/src/github.com/luebken/stpl/cmd/stplsrv

# build docker 
builddocker:

	# build builder image which calls `make buildgo`
	docker build --no-cache -t luebken/build-stpl -f ./Dockerfile.build .

	# cp binary to local file 
	docker run -t luebken/build-stpl /bin/true
	docker cp `docker ps -q -n=1`:/stplsrv .
	chmod 755 ./stplsrv

	# build final image
	docker build --rm=true --tag=luebken/stpl  .

rundocker:
	docker run luebken/stpl

# written with https://github.com/sstephenson/bats
e2etests:
	cd e2etests && bats e2etests.sh