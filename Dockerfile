FROM golang:1.24-alpine AS binacsgobuild

COPY . /src

RUN apk add --no-cache make git build-base && \
    \
    cd /src && \
    \
    make

FROM alpine

COPY --from=binacsgobuild /src/bin/server /src/test /work/

CMD ["./work/server", "version"]