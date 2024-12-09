FROM golang:1.23 AS build

WORKDIR /build

COPY go.mod go.sum ./     

RUN go mod download      

COPY . .       

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bookstorage ./cmd/main

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /build/bookstorage .

COPY --from=build /build/migrations ./migrations

CMD ["./bookstorage"]