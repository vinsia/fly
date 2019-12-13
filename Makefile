build:
	./script/setup.sh

install: build
	$(shell sudo cp bin/fly /usr/local/bin/fly)

default: build
