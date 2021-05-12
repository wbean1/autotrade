#first stage - builder
FROM golang:1.12.0-stretch as builder
COPY . /cf-garbage-truck
WORKDIR /cf-garbage-truck
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -v
#second stage
FROM alpine:latest
WORKDIR /root/
RUN apk add curl
COPY --from=builder /cf-garbage-truck .
CMD ["./cf-garbage-truck"]