

Requirements:
1. Go v1.20.3
2. redis

Run without docker:
1. Run `go mod tidy`
2. Run `go mod vendor`
3. If you have install Make then run `make dev`
4. If you haven't installed Make then run `go build -o bin/todo app/main.go` then `./bin/todo`


Requirements:
1. Docker & Docker Compose

Run with docker
1. Run `docker-compose up --build` or `make docker-build`
2. If you already run `docker-compose up --build` then for after that you can just run `docker-compose up` or `make docker-up`