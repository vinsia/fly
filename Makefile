build:
	go build -o bin/fly main.go

install: build
	$(shell sudo cp bin/fly /usr/local/bin/fly)

default: build
