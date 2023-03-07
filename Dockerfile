FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

COPY src /app/src
COPY src/web/templates /app/web/templates
COPY src/web/static /app/web/static

RUN go build /app/src/server.go

EXPOSE 8080

CMD ["/app/server"]
