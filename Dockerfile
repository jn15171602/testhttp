FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod download

# RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /app/myapp
RUN go build -o /app/myapp

CMD ["/myapp"]FROM ubuntu:latest