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

# Interval between API requests.
# Supports Go duration format (e.g. "15s")
# or plain integer in seconds (e.g. "15" = 15s).
ENV STF_API_REQUEST_INTERVAL=

# List of STF provider base URLs separated by ";"
# Example: https://demo.zebrunner.farm/;https://demo2.zebrunner.farm/
ENV STF_PROVIDERS=

# List of API tokens separated by ";" (must match order of STF_PROVIDERS)
# Example: token1;token2
ENV STF_TOKENS=

# List of device serials that will always be skipped, regardless of provider.
# Devices in this list will never be included.
# Example: serial1;serial2
ENV SKIP_SERIALS=

CMD ["./stf-monitor"]