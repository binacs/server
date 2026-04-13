FROM golang:1.24-alpine AS binacsgobuild

RUN apk add --no-cache make git build-base

WORKDIR /src

# Cache dependency downloads in a separate layer
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make

FROM alpine

COPY --from=binacsgobuild /src/bin/server /src/test /work/

CMD ["./work/server", "version"]