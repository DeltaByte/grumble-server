# Build
FROM golang:1.16-alpine AS build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o bin/grumble-server server.go


# Run
FROM alpine:latest

COPY --from=build /go/src/app/bin/ ./
RUN mkdir /data /data/database /data/media /data/logs

ENV GRUMBLE_PORT=80
ENV GRUMBLE_STORAGE_DATABASE=/data/database
ENV GRUMBLE_STORAGE_MEDIA=/data/media
ENV GRUMBLE_STORAGE_LOGS=/data/logs

EXPOSE 80/tcp

ENTRYPOINT [ "./grumble-server" ]