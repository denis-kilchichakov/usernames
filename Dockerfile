# syntax=docker/dockerfile:1
# Build stage
FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

ADD cmd/uzrnames .

RUN go build -o main .

# Production stage
FROM alpine:latest 
COPY --from=0 /app/main ./

CMD [ "./main" ]