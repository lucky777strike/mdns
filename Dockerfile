FROM golang:latest AS build-env
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod /src/
RUN go mod download
COPY . .
RUN  go build  -o ./mdns  -ldflags="-s -w" -gcflags="all=-trimpath=/src" -asmflags="all=-trimpath=/src"


FROM alpine:latest
RUN apk add --no-cache ca-certificates \
    && rm -rf /var/cache/*
WORKDIR /app
COPY --from=build-env /src/mdns .
EXPOSE 8080 1234
ENTRYPOINT [ "./mdns" ]
