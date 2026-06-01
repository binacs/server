# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Local Development Environment

Local dev runs entirely via Docker Compose. **Always use docker compose for building, running, and debugging — do not run the server or its dependencies locally.**

```bash
# Start full environment (server + MySQL + Redis + Jaeger + cryptfunc services)
cd scripts/docker-compose && docker compose up --build

# Rebuild and restart only the server
cd scripts/docker-compose && docker compose up --build server_dc

# View logs
cd scripts/docker-compose && docker compose logs -f server_dc

# Stop everything
cd scripts/docker-compose && docker compose down
```

The compose stack (`scripts/docker-compose/docker-compose.yml`) includes:
- **server_dc** — the app itself (built from repo root Dockerfile), ports 80/443/9500/9999
- **mysql_dc** — MySQL (root/password, database `testdb`), port 3306
- **redis_dc** — Redis (password: `password`), port 6379
- **jaeger_dc** — Jaeger tracing UI at port 16686
- **cryptfunc services** — BASE64, AES, DES encryption microservices

Config for compose is at `scripts/docker-compose/config.toml`. Services reference each other by container name (e.g. `mysql_dc:3306`, `redis_dc:6379`). All containers share the `binacs_local` bridge network.

## Build Commands

```bash
make              # clean + build (outputs bin/server)
make build        # build only
make test         # run all tests with coverage
make test-coverage # generate HTML coverage report
make mock         # regenerate mocks (gateway + service)
make docker       # build Docker image

# Run a single test
go test ./service/ -run TestFunctionName -cover

# Regenerate protobuf code for a specific API
cd api && make <service>  # e.g., make crypto, make pastebin, or make all
```

## Architecture

This is a Go web server (binacs.space) using **Gin** for HTTP and **gRPC-Gateway** for gRPC+REST. It provides several services: crypto (encrypt/decrypt), tinyurl (URL shortening), pastebin (paste sharing with optional password protection), blog, COS (cloud object storage), and user management.

### Dependency Injection

The app uses `binacsgo/inject` for DI. All services are registered in `cmd/commands/start.go` with string-based inject names defined in `cmd/commands/inject_const.go`. Structs use `inject-name:"..."` struct tags to receive dependencies. Types implementing DI must have an `AfterInject()` method for post-injection initialization.

### Key Layers

- **cmd/commands/** - CLI entry point using cobra. `root.go` handles config loading and logger init; `start.go` wires all DI and boots the node.
- **gateway/** - Server layer. `node.go` orchestrates startup (pprof, trace, web, gRPC). `web.go` sets up Gin routes and handlers. `grpc.go` sets up gRPC server with TLS, auth interceptors, and grpc-gateway mux.
- **service/** - Business logic. `interface.go` defines all service interfaces. Each service (crypto, tinyurl, pastebin, blog, cos, user) has its own implementation file. DB services: `mysql.go` (xorm), `redis.go`.
- **api/** - Protobuf definitions and generated code for each gRPC service. Use `go generate` via `api/Makefile` to regenerate.
- **config/** - TOML-based configuration (`config.toml`). Supports hot reload. Configurable mode: `all`, `web`, or `grpc`.
- **types/** - Shared types, constants, and `table/` for DB table models.
- **middleware/** - Gin middleware for auth (token-based), TLS redirect, and Jaeger tracing.
- **mock/** - Generated mocks (golang/mock) for gateway and service interfaces.

### Service Registration Pattern

Each API service (crypto, tinyurl, pastebin, cos, user) implements a `Register()` method that registers both the gRPC server handler and the gRPC-Gateway reverse proxy handler simultaneously.

### Configuration

Config is loaded from a TOML file (default `./config.toml`). The `Config` struct in `config/config.go` contains sub-configs for web, gRPC, tracing (Jaeger), logging (lumberjack), Redis, MySQL, COS, and pprof.
