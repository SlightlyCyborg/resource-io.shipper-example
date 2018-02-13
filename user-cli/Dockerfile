FROM golang:1.9.0 as builder

WORKDIR /go/src/resource-io/shipper/user-cli

COPY . .

RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/resource-io/shipper/user-cli/user-cli  .

CMD ["./user-cli"]