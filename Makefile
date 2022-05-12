.PHONY: build install

BINDIR := /usr/bin
EXE := seribund

build:
	@echo '### Building Executable... ###'
	go build -o ${EXE} main.go

install: build
	@echo '### Adding Symlink to ${BINDIR}... ###'
	sudo ln -s $(shell pwd)/${EXE} ${BINDIR}/${EXE}

uninstall:
	@echo '### Removing Symlink from ${BINDIR}... ###'
	sudo rm ${BINDIR}/${EXE}