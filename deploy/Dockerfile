FROM golang:1.25.1-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o stf-monitor ./cmd/main.go

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/stf-monitor .

ENV STF_PROVIDERS=
ENV STF_TOKENS=
ENV SKIP_SERIALS=

CMD ["./stf-monitor"]