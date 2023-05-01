# stage 0 - build MySQL container
FROM mysql:8.0.23 AS mysql_build

ENV MYSQL_HOST=localhost \
    MYSQL_PORT=3306 \
    MYSQL_DBNAME=mydb \
    MYSQL_USER=myuser \
    MYSQL_PASSWORD=mypassword

COPY ./db/migrations/*.sql /docker-entrypoint-initdb.d/

CMD ["mysqld"]

# stage 0 - build Go application container
FROM golang:1.20.3-alpine AS build

# install some tools
RUN apk add --no-cache ca-certificates git make \
    && rm -rf /var/cache/apk/*

# set variables
ENV GOPATH=/go \
    PATH=$GOPATH/bin:$GOROOT/bin:$PATH \
    GOOS=linux \
    GOARCH=amd64

# set project folder location inside docker
WORKDIR /go/src/todo

# build executeable
COPY . .

ENV MYSQL_HOST=localhost \
    MYSQL_PORT=3306 \
    MYSQL_USER=myuser \
    MYSQL_PASSWORD=mypassword \
    MYSQL_DBNAME=mydb

RUN go mod tidy \
    && go mod vendor \
    && go build -o bin/todo app/main.go



# stage 1
FROM golang:1.20.3-alpine AS run

# set variables
RUN export GOPATH=/go \
    && export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

COPY --from=build /go/src/todo /go/src/todo

RUN go install -tags 'mysql,postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /go/src/todo

RUN chmod +x bin/todo

CMD ["sh", "-c", "migrate -database 'mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DBNAME}' -path db/migrations up && bin/todo"]