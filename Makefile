.PHONY: build dockerize test

name = incrementor-0-1-0-linux-amd64

build:
	rm -rf ./$(name) && \
	env GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o ./$(name) ./ && \
	upx -9 ./$(name)
