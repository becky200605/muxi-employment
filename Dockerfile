FROM golang:1.25.3-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct

# 替换阿里源，加必要软件
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update --no-cache \
    && apk add --no-cache tzdata

WORKDIR /build

# 下载依赖
ADD go.mod .
ADD go.sum .
RUN go mod download

# 复制源码
COPY . .

# 编译
RUN go build -ldflags="-s -w" -o /app/muxiemployment muxiemployment.go

# 构建最小镜像
FROM scratch

# 时区、证书
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app

# 拷贝可执行文件
COPY --from=builder /app/muxiemployment /app/muxiemployment

# 拷贝 etc 配置文件
COPY ./etc /app/etc

# ⚡ 关键：把 internal/data 图片目录也拷贝进去
COPY ./internal/data /app/internal/data

CMD ["./muxiemployment", "-f", "etc/muxi_employment.yaml"]
