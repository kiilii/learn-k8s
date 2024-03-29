FROM golang:alpine AS builder

ARG APP_RELATIVE_PATH

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata


COPY . /src
WORKDIR /src/${APP_RELATIVE_PATH}
RUN --mount=type=cache,target=/go go mod download
RUN --mount=type=cache,target=/go go build -ldflags="-s -w" -o /app/server .

FROM scratch

ARG APP_RELATIVE_PATH

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/server /app/server
COPY --from=builder /src/output/gateway /app/gateway
COPY --from=builder /src/${APP_RELATIVE_PATH}/etc /app/etc

CMD ["./server", "-f", "etc/config.yaml"]
