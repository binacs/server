FROM golang:1.15-alpine AS binacsGoBuild

# ENV GO111MODULE=on

COPY . /src

RUN apk add --no-cache make git \
    && cd /src \
    && make

FROM alpine

COPY --from=binacsGoBuild /src/bin/server \
    /src/test \
    /work/

EXPOSE 9500 443 80

CMD ./work/server start --configFile /work/docker-compose/config.toml