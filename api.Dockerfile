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
COPY ./entrypoint.sh /entrypoint.sh

# wait-for-it requires bash, which alpine doesn't ship with by default. Use wait-for instead
ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh

ENTRYPOINT [ "sh", "/entrypoint.sh" ]