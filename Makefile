.PHONY: all build clean

all: build

build:
	go build -o server .

clean:
	rm -f server
