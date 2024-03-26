FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

# WORKDIR /build

# ADD go.mod .
# ADD go.sum .
# RUN go mod download
# COPY . .
# COPY ./etc /app/etc
# RUN go build -ldflags="-s -w" -o /app/order .

COPY . /src
WORKDIR /src

RUN --mount=type=cache,target=/go make build

RUN ls ./bin

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

EXPOSE 8000
EXPOSE 9000
VOLUME ./configs/ /data/conf/

WORKDIR /app
COPY --from=builder /src/bin/order /src/bin/order
COPY --from=builder /app/etc /app/etc

CMD ["./order", "-f", "etc/config.yaml"]
