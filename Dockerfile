FROM golang:1.20-alpine AS builder
WORKDIR /src

# Cache deps
COPY go.mod .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /hello

FROM alpine:3.18
RUN addgroup -S appgroup && adduser -S appuser -G appgroup -u 1001
COPY --from=builder /hello /usr/local/bin/hello
COPY --from=builder /src/docs /docs
RUN chown appuser:appgroup /usr/local/bin/hello
RUN chown -R appuser:appgroup /docs || true
USER appuser
EXPOSE 8080
ENV PORT=8080
ENTRYPOINT ["/usr/local/bin/hello"]
