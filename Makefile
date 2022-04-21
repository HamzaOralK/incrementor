.PHONY: build dockerize test

build:
	rm -rf ./incrementor && \
	env GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o ./incrementor ./ && \
	upx -9 ./incrementor
