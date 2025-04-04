FROM golang:latest

# Get Go tools
RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

CMD ["/docker-entrypoint.sh"]
