.PHONY: build test run docker-build clean

build:
	go build -o bin/main cmd/main.go

test:
	go test ./tests -v

run:
	go run cmd/main.go

docker-build:
	docker build -t weather-aggregator .

clean:
	rm -rf bin/
