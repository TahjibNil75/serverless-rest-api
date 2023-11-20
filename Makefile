.PHONY: build clean deploy remove

build:
	go get ./...
	go mod vendor
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/create_author api/create_author/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get_authors api/get_authors/main.go


clean:
	rm -rf ./bin

deploy: build
	sls deploy

remove:
	sls remove

