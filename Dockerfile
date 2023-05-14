FROM golang:alpine3.18 as builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY cmd /app/cmd/
RUN go build -o app-binary cmd/server/server.go

FROM alpine:3.18.0
WORKDIR /
COPY --from=builder /app/app-binary /app-binary
ENTRYPOINT [ "/app-binary" ] 
