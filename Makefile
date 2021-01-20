.PONEY: all test install clean

all: build

build: go-test-ignore

go-test-ignore: *.go
	go build .

test:
	go test -cover -tags GO_TEST ./...

install: build
	install go-test-ignore /usr/bin/ 

clean:
	rm -f go-test-ignore
