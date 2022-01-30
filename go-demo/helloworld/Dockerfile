FROM golang:1.17-alpine AS builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /build

COPY go.mod .
# COPY go.sum .

RUN go mod tidy

COPY . .

RUN go build -o ./helloworld .

FROM alpine

WORKDIR /

COPY --from=builder /build/helloworld /

EXPOSE 8080

CMD ["/helloworld"]