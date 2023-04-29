dep:
	go mod tidy
	go mod vendor

# Use this only for development
dev:
	go build -o bin/todo app/main.go
	./bin/todo

build:
	set GOOS=linux && set GOARCH=amd64 && go build -o bin/todo app/main.go

docker-build:
	docker-compose up --build

docker-up:
	docker-compose up

migrate-up:
	migrate -database "mysql://root@tcp(localhost:3306)/todo" -path db/migrations up
	
migrate-down:
	migrate -database "mysql://root@tcp(localhost:3306)/todo" -path db/migrations down