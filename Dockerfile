FROM golang:1.17-alpine AS binacsGoBuild

COPY . /src

RUN apk add --no-cache make git && \
    \
    cd /src && \
    \
    make

FROM alpine

COPY --from=binacsGoBuild /src/bin/server /src/test /work/

CMD ./work/server version