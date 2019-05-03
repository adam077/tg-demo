FROM golang:1.11.1 AS build
ENV TZ=Asia/Shanghai LANG=C.UTF-8 TERM=xterm
WORKDIR /go/src/go-go-go
COPY . .
RUN go build -tags=jsoniter -o main src/main.go
EXPOSE 8080
ENTRYPOINT ["./main"]
