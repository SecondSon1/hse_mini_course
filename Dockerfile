FROM golang:1.22-alpine3.20 as builder
RUN apk add --no-cache make

WORKDIR /app

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . ./
RUN CGO_ENABLED=0 make grpc_db_server

FROM alpine
WORKDIR /
COPY --from=builder /app/bin/grpc_db_server .

ENV SERVER_PORT=6969
EXPOSE 6969

CMD ["./grpc_db_server"]
