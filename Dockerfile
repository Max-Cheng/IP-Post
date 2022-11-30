FROM --platform=$BUILDPLATFORM golang:1.18 as builder
ARG TARGETARCH
ARG TARGETOS
WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLE=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build .

FROM --platform=$BUILDPLATFORM alpine:latest
WORKDIR /app
RUN apk add --no-cache libc6-compat
COPY --from=builder /app/IPPost .
ENV PORT 9768
ENTRYPOINT ["./IPPost"]