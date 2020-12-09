build:
	export GO111MODULE="on"; \
	go mod download; \
	go mod vendor; \
	GO111MODULE="on" CGO_ENABLED=0 go build -mod=vendor -a -ldflags '-s' -installsuffix cgo -o main cmd/main.go

install:
	go mod download

run:
	go run cmd/main.go

test:
	go test ./...