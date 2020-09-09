FROM golang:1.14-alpine AS binacsGoBuild

# ENV GO111MODULE=on

COPY . /src

RUN apk add --no-cache make git && go env -w GOPROXY=https://goproxy.cn \
    && cd /src \
    && make

FROM alpine

COPY --from=binacsGoBuild /src/bin/server \
    /src/test \
    /work/

EXPOSE 9500 443 80

CMD ./work/server start --configFile /work/docker-compose/config.toml