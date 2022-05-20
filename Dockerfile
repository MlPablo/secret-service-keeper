FROM golang
 # Create working directory under /app
WORKDIR /app
 # Copy over all go config (go.mod, go.sum etc.)
COPY go.* .
 # Install any required modules
RUN go mod download
RUN go mod verify
 # Copy over Go source code
COPY . .
 # Run the Go build and output binary under hello_go_http
RUN go build -v -o /secretservice
 # Make sure to expose the port the HTTP server is using
EXPOSE 8080
 # Run the app binary when we run the container
CMD ["go","run","main.go"]