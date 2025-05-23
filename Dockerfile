FROM golang:1.24.3-bookworm AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
COPY goradiru/ ./goradiru
COPY cmd/ ./cmd

RUN go build -o /app/main -ldflags '-s -w' main.go


FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    ffmpeg \
    && rm -rf /var/lib/apt/lists/*

RUN groupadd -g 1000 nonroot && useradd -u 1000 -g 1000 nonroot

USER nonroot:nonroot
WORKDIR /app

COPY --chown=nonroot:nonroot --from=build /app/main /app/goradiru

ENTRYPOINT ["/app/goradiru"]
