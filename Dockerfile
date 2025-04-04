FROM golang:1.23

# Get Go tools
RUN go install github.com/air-verse/air@latest

WORKDIR /app
COPY . .
RUN go mod download

COPY docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

CMD ["/docker-entrypoint.sh"]
