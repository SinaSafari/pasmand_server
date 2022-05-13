FROM golang:1.18-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o main ./cmd/api/main.go

WORKDIR /dist

RUN cp /build/main .

EXPOSE 8080

CMD ["/dist/main"]