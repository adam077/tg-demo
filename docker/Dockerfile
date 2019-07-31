#
# Build stage
#
FROM golang:1.11.1 AS build

WORKDIR /go/src/tg-demo

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags=jsoniter -o main src/main.go

EXPOSE 8080

#
# Production stage
#
FROM alpine:3.8

ENV TZ=Asia/Shanghai LANG=C.UTF-8 TERM=xterm

RUN apk --update add ca-certificates tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

WORKDIR /go/src/tg-demo

COPY --from=build /go/src/tg-demo/main .

ENTRYPOINT ["./main"]
