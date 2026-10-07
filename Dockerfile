FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o incident-agent ./cmd/agent


FROM alpine:3.22

WORKDIR /app

RUN addgroup -S agent && \
    adduser -S agent -G agent

COPY --from=builder /app/incident-agent .

USER agent

ENTRYPOINT ["./incident-agent"]
