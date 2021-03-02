FROM golang:alpine as builder

ARG PKG=${PKG}
ARG GITCOMMIT=${GITCOMMIT}
ARG BUILDVERSION=${BUILDVERSION}
ARG BUILDDATE=${BUILDDATE}
ARG TARGET=${TARGET}

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn,direct"

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache build-base

WORKDIR /go/src/${PKG}

COPY . .

RUN cd /go/src/${PKG}/cmd/${TARGET} && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v \
    -o /go/bin/${TARGET} \
    -ldflags "-X ${PKG}/pkg/info.GitCommit=${GITCOMMIT} \
    -X ${PKG}/pkg/info.BuildDate=${BUILDDATE} \
    -X ${PKG}/pkg/info.Version=${BUILDVERSION}"

FROM alpine:3.12

ARG TARGET=${TARGET}

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache ca-certificates

COPY --from=builder /go/bin/${TARGET} /usr/local/bin

EXPOSE 8001

CMD ["preupgradechecks"]
