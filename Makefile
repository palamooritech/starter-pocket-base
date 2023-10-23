build:
	go build -o bin/app

run: build
		./bin/app

test:
	go test -v ./...

serve: build
		./bin/app serve