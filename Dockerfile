From golang:1.14-alpine

RUN apk add --no-cache make git && go env -w GOPROXY=https://goproxy.cn

WORKDIR /src

COPY . .

EXPOSE 9500 443

CMD make && sleep 3s && ./bin/server start --configFile ./test/docker-compose/config.toml