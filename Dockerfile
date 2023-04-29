# stage 0
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
RUN go mod tidy \
    && go mod vendor \
    && go build -o bin/todo app/main.go



# stage 1
FROM golang:1.20.3-alpine AS run

# set variables
RUN export GOPATH=/go \
    && export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

COPY --from=build /go/src/todo /go/src/todo

WORKDIR /go/src/todo

RUN chmod +x bin/todo

CMD ["bin/todo"]