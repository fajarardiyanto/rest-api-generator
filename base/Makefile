help:
	@go run ./cmd/main.go help
tidy:
	@go mod tidy
download:
	@go mod download
test:
	@go test -v ./...
run: tidy
	@go run ./cmd/main.go api
scan:
	@script/gosec.sh
generate:
	@go run ./cmd/main.go config
build:
	@go build \
		-ldflags "-X main.buildName=health-service -X main.buildVersion=`git rev-parse --short HEAD`" \
		-o health-service.app cmd/api/main.go
docker-build:
	@docker build -f Dockerfile -t "health_service:1.0" --build-arg BUILD_DATE="$(date -u +"%Y-%m-%dT%H:%M:%SZ")" .
docker-run-container:
	@docker container run -ti --init --rm \
		--name health_service \
		--memory="1g" --cpus="1" \
		--net host \
		-v $$HOME/go/docker:/go \
		-v $(PWD):/data \
		-w /data \
		redhat/ubi8-micro:8.5-437 ./health-service.app

