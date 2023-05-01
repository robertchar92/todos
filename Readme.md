

Requirements:
1. Go v1.20.3
2. mysql

Run without docker:
1. Run `go mod tidy`
2. Run `go mod vendor`
3. Add .env file with example from .env.example
4. Run go install -tags 'mysql,postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
5. Run migrate -database "mysql://{mysql_user}:{mysql_passwor}@tcp({mysql_host}:{mysql_port})/{database_name}" -path db/migrations up (Change all variables inside {} to your mysql variables)
6. If you have install Make then run `make dev`
7. If you haven't installed Make then run `go build -o bin/todo app/main.go` then `./bin/todo`


Requirements:
1. Docker & Docker Compose

Run with docker
1. Add .env file with example from .env.example
2. Run `docker-compose up --build` or `make docker-build`
3. If you already run `docker-compose up --build` then for after that you can just run `docker-compose up` or `make docker-up`