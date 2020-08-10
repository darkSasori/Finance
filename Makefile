test:
	go test -v -race ./...

build-cmd:
	go build -o bin/cli ./cmd/...

mongodb:
	docker run -d --rm --name mongodb -p 27017:27017 mongo
