FROM golang:bullseye



RUN apt-get update
RUN apt-get install htop -y

# Set the working directory
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

# Copy the server code into the container
#COPY . .

RUN go get ./...

# Make port 8080 available to the host
EXPOSE 8080

# Build and run the server when the container is started
#RUN go build /app/server.go
#ENTRYPOINT ./server

CMD ["air"]
