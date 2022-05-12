.PHONY: build-all build-linux build-win install uninstall

BINDIR := /usr/bin
EXE := seribund

build-all: build-linux build-win

build-linux:
	@echo '### Building Linux Executable... ###'
	GOOS=linux go build -o ${EXE} main.go

build-win:
	@echo '### Building Windows Executable... ###'
	GOOS=windows go build -o ${EXE}.exe main.go

install: build
	@echo '### Adding Symlink to ${BINDIR}... ###'
	sudo ln -s $(shell pwd)/${EXE} ${BINDIR}/${EXE}

uninstall:
	@echo '### Removing Symlink from ${BINDIR}... ###'
	sudo rm ${BINDIR}/${EXE}