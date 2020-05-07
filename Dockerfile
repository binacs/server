FROM golang:1.14-alpine AS binacsGoBuild

# ENV GO111MODULE=on

COPY . /src

RUN apk add --no-cache make git && go env -w GOPROXY=https://goproxy.cn \
    && cd /src \
    && make

FROM alpine

COPY --from=binacsGoBuild /src/bin/server \
    /src/test \
    /src/test/

EXPOSE 9500 443

CMD sleep 30s && ./src/test/server start --configFile /src/test/docker-compose/config.toml