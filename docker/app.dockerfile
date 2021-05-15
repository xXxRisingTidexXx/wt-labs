FROM golang:1.16.3-alpine3.13
WORKDIR /go/src/app
ENV GOOS "linux"
ENV GOARCH "amd64"
ENV CGO_ENABLED 0
ADD https://raw.githubusercontent.com/eficode/wait-for/master/wait-for .
COPY go.mod .
COPY go.sum .
COPY cmd cmd
COPY pkg pkg
COPY internal internal
RUN chmod +x wait-for && apk add --no-cache git && go get ./...
