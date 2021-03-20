FROM golang:1.16-alpine3.13  as builder

ARG PKG=${PKG}
ARG GITCOMMIT=${GITCOMMIT}
ARG BUILDVERSION=${BUILDVERSION}
ARG BUILDDATE=${BUILDDATE}
ARG TARGET=${TARGET}

WORKDIR /go/src/${PKG}

COPY . .

ENV CGO_ENABLED=0

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add --no-cache make git && \
    env TARGETS=${TARGET} make build && cp bin/${TARGET} /go/bin/${TARGET}

FROM alpine:3.13

ARG TARGET=${TARGET}

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/bin/${TARGET} /usr/local/bin

EXPOSE 8001

CMD ["preupgradechecks"]
