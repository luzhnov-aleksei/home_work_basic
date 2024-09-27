FROM golang:1.20 AS builder
RUN git clone -b hw16_docker https://github.com/luzhnov-aleksei/home_work_basic.git /usr/local/go/src/go_sql/
WORKDIR /usr/local/go/src/go_sql
RUN go clean --modcache && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -mod=readonly -o go_sql /usr/local/go/src/go_sql/hw15_go_sql/server/main.go

FROM scratch

USER 1001

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/go/src/go_sql/go_sql /app/
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD curl --fail http://localhost:8080/health || exit 1
CMD ["/app/go_sql"]
