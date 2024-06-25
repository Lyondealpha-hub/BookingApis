FROM golang:1.22.2-alpine as base
RUN apk add openssl
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
FROM base as dev
RUN go install github.com/cosmtrek/air@v1.52.1;
COPY . .
EXPOSE 80
CMD ["sh", "-c", "air"]