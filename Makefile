.PHONY: env  clean lint build

all: env  clean lint build

env:
	@echo "=========install goimports ==========="
	GOPROXY=https://goproxy.cn/,direct go install -v github.com/incu6us/goimports-reviser/v3@latest
	@echo "=========install gofmt ==========="
	GOPROXY=https://goproxy.cn/,direct go install mvdan.cc/gofumpt@latest

build:
	go mod tidy
	gofumpt -l -w .
	CGO_ENABLED=0 go build  -ldflags "$(LDFLAGS)" -v .
	
clean:
	go clean -i .

run:
	go mod tidy
	gofumpt -l -w .
	CGO_ENABLED=0 go build  -ldflags "$(LDFLAGS)" -v .
	./gg