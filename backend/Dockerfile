FROM golang:1.23-alpine

RUN apk add --no-cache git

RUN go install github.com/air-verse/air@latest

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["air", "-c", ".air.toml"]