FROM golang:alpine as builder
WORKDIR /go/src/app
COPY . .
RUN go get -d ./...
RUN go build -o main
RUN go build cmd/migrate.go

FROM alpine:latest
WORKDIR /root
COPY --from=builder /go/src/app/main .
COPY --from=builder /go/src/app/migrate .
COPY public public
COPY templates templates
COPY conf conf

RUN apk add tzdata

EXPOSE 8843

CMD ["./main"]
