default: builddocker

setup:
    #go get golang.org/x/oauth2
    #go get golang.org/x/oauth2/jwt
    #go get google.golang.org/api/analytics/v3

# build go binary
# gets called within docker build
buildgo:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o stplcli ./go/src/github.com/luebken/stpl/cmd/stplcli

# build docker 
builddocker:
	# build builder image which calls `make buildgo`
	docker build --no-cache -t luebken/build-stpl -f ./Dockerfile.build .
	# cp binary to local file 
	docker run -t luebken/build-stpl /bin/true
	docker cp `docker ps -q -n=1`:/stplcli .
	chmod 755 ./stplcli
	# build final image
	docker build --rm=true --tag=luebken/stpl  .

run:
	docker run luebken/stpl