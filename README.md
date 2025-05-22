# Mongo Ops Exporter

Mongo Ops Exporter is a small API written in Go, designed to expose slow operations in one or more MongoDB instances via HTTP endpoints. This tool can be useful for observability, SRE, or development teams that need to identify bottlenecks and performance issues in their Mongo databases.

## Features

- Support for multiple MongoDB instances.
- Configuration loaded from a YAML file.
- HTTP endpoints to list databases and query slow operations.
- Simple in-memory cache to avoid repeated queries.

## Requisites

- Go 1.18 or higher
- Access to the mongodb instances to be monitorized, and permissions to execute db.currentOp() command

## Instalation

Clone the repository:

```bash
git clone https://github.com/MGYOSBEL/mongo-ops-exporter.git
cd mongo-ops-exporter
```

Compile the binary

```bash
go build -o mongo-ops-exporter
```

Execute the server:

```bash
./mongo-ops-exporter -config.file=config.yaml
# o de forma abreviada
./mongo-ops-exporter -f config.yaml
```

By default, the server expose the endpoints in <http://localhost:8080>.

## Build with docker

The image can be built with the included Makefile just executing:

`make build`

This will generate an image tagged as **mongo-ops-exporter:0.0.1**. This can be changed editing the Makefile.

### Execute with docker

```
docker run -p 8080:8080 \
-v $(pwd)/config.yaml:/app/config.yaml \
mongo-ops-exporter:0.0.1 \
-config.file=config.yaml
```

Or just run a compose like this one:

```yaml
services:
  mongo-ops-exporter:
    image: mongo-ops-exporter:0.0.1
    container_name: mongo-ops-exporter
    ports:
      - "8088:8080"
    volumes:
      - ./mongo-ops-exporter.yaml:/etc/mongo-ops-exporter/config.yaml
    command:
      - "--config.file=/etc/mongo-ops-exporter/config.yaml"
    networks:
      - your-network
```
