# Prepare
FROM golang:1.17.0-buster AS prepare
WORKDIR /source
COPY go.mod .
COPY go.sum .
RUN go mod download

# Build
FROM prepare AS build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/miner cmd/main.go

# Run
FROM scratch as run
COPY --from=build /source/bin/miner /miner
EXPOSE 8080
ENTRYPOINT ["/miner"]