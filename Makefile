build:
	./build.sh

cgen:
	c-for-go --ccincl --ccdefs --nostamp gogc.yml

run-gcfatal: build
	go run gcfatal/main.go

run-enomem: build
	go run enomem/main.go

clean:
	rm -rf ./libgogc.a
