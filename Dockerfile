FROM golang:bullseye

# Set the working directory
WORKDIR /app

RUN go mod init minitwit 

#RUN go install github.com/cosmtrek/air@latest

# Copy the server code into the container
COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY src /app/src
COPY src/web/templates /app/web/templates
COPY src/web/static /app/web/static
RUN go build /app/src/server.go

# Make port 8080 available to the host
EXPOSE 8080

# Build and run the server when the container is started
#RUN go build /app/server.go
#ENTRYPOINT ./server

CMD ["/app/server"]