FROM golang:alpine AS build-env
COPY . /src
RUN cd /src && go build -o mongo-ops-exporter main.go

FROM alpine:3.21
RUN mkdir -p /app
COPY --from=build-env src/mongo-ops-exporter /app
COPY --from=build-env src/config.yaml /app
WORKDIR /app
EXPOSE 8080
ENTRYPOINT [ "./mongo-ops-exporter" ]
