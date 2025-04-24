FROM golang:alpine AS build-env
COPY . /src
RUN cd /src && go build -o mongo-ops-exporter main.go

FROM alpine:3.21
RUN mkdir -p /app
COPY --from=build-env src/mongo-ops-exporter /app
WORKDIR /app
ENTRYPOINT [ "./mongo-ops-exporter" ]
