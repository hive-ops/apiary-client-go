setup-project:
	npm i -g nodemon
	brew install bufbuild/buf/buf
	#	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	#	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

mod-tidy:
	go mod tidy

mod-vendor: mod-tidy
	go mod vendor

vet:
	go vet ./...

dev: start-services
	nodemon

start-services:
	docker-compose down
	docker-compose up -d

compile-proto:
	find ./pb -name "*.pb.go" -exec rm {} +
	buf generate --path=apiary-proto/apiary

pull-submodules:
	git submodule update --remote --merge --recursive
	make compile-proto

tests:
	go test -v `go list ./... | grep -v ./pb` -race -coverprofile=coverage.out; go tool cover -html=coverage.out
