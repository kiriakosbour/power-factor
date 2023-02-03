FROM golang:1.19-alpine AS builder
# Metadata for the build, compile only for linux target, smaller binary size

ENV GOOS linux
ENV GOARCH amd64
# Turn off CGO since that can result in dynamic links to libc/libmusl.
ENV CGO_ENABLED 0
WORKDIR /pf
#add git needed bellow
#RUN apk add git
RUN apk update && apk add ca-certificates && apk add tzdata
# Copy `go.mod` for definitions and `go.sum` to invalidate the next layer
# in case of a change in the dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# Copy and build the app
COPY . .
RUN go build -ldflags="-w -s" -o power-factor .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /pf/power-factor .
EXPOSE 8080
CMD ["./power-factor"]