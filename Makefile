BUILD_TAGS?=
BUILD_FLAGS = -ldflags "-X github.com/BinacsLee/server/version.GitCommit=`git rev-parse HEAD`"

default: clean build

clean:
	rm -rf bin

build:
	go build $(BUILD_FLAGS) -tags '$(BUILD_TAGS)' -o bin/server ./cmd

mock:
	cd gateway && go generate; cd -
	cd service && go generate; cd -

docker:
	docker build -t binacslee/binacs-cn:latest . 

test:
	go test ./... -cover

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

.PHONY: mock test