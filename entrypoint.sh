wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

# CompileDaemon --build=""

go mod tidy

go mod vendor

go build -o bin/todo app/main.go

./bin/todo