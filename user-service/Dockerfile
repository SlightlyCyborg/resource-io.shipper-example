#Dockerfile taken mostly from https://github.com/EwanValentine/shippy/blob/master/user-service/Dockerfile

FROM golang:1.9.0 as builder

# This dir must have the same structure as the go directory
WORKDIR /go/src/resource-io/shipper/user-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/resource-io/shipper/user-service/user-service .

# I built a run script, because go-micro uses several options that need to be included
CMD ["./user-service"]
