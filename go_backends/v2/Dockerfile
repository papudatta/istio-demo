FROM golang:1.13.12-alpine AS build
RUN set -ex; \
    apk update; \
    apk add --no-cache git
ADD . /src
WORKDIR /src
RUN go get -d -v -t
RUN go build -v -o goserv2
FROM alpine:3.4
# RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
# ENV DB db
COPY --from=build /src/goserv2 /usr/local/bin/goserv2
RUN chmod +x /usr/local/bin/goserv2
EXPOSE 9091
CMD ["goserv2"]