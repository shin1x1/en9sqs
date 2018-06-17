.PHONY: all
all: dep build

.PHONY: dep
dep:
	dep ensure -vendor-only

.PHONY: build
build:
	go build -o ./bin/en9sqs en9sqs.go
