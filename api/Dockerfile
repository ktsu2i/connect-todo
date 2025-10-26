FROM golang:1.25-bookworm AS dev

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8081

CMD ["air", "-c", ".air.toml"]
