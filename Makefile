build:
	/usr/local/go/bin/go build -o main .
	docker compose build
run:
	docker compose up
down:
	docker compose down
all: build run
