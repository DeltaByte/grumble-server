# Build
FROM golang:1.16-alpine3.13 as builder

RUN apk add --no-cache gcc g++

WORKDIR $GOPATH/src/github.com/grumblechat/server

COPY go.mod go.sum ./

RUN go mod verify

COPY . .

RUN go run scripts/build.go build

# Run
FROM alpine:3.13

ARG GRUMBLE_UID="472"
ARG GRUMBLE_GID="0"

# set PATH
ENV PATH="/usr/share/grumble/bin:$PATH" \
    GRUMBLE_PATHS_HOME="/usr/share/grumble" \
    GRUMBLE_PATHS_CONFIG="/etc/grumble" \
    GRUMBLE_PATHS_DATABASE="/var/lib/grumble/database" \
    GRUMBLE_PATHS_BACKUP="/var/lib/grumble/backups" \
    GRUMBLE_PATHS_MEDIA="/var/lib/grumble/media" \
    GRUMBLE_PATHS_LOGS="/var/log/grumble"

WORKDIR $GRUMBLE_PATHS_HOME

# create group if needed
RUN if [ ! $(getent group "$GRUMBLE_GID") ]; then \
      addgroup -S -g $GRUMBLE_GID grumble; \
    fi

# create user
RUN export GRUMBLE_GID_NAME=$(getent group $GRUMBLE_GID | cut -d':' -f1) && \
    adduser -S -u $GRUMBLE_UID -G "$GRUMBLE_GID_NAME" grumble

# create directories
RUN mkdir -p "$GRUMBLE_PATHS_CONFIG" \
             "$GRUMBLE_PATHS_DATABASE" \
             "$GRUMBLE_PATHS_BACKUP" \
             "$GRUMBLE_PATHS_MEDIA" \
             "$GRUMBLE_PATHS_LOGS"

# set directory ownership
RUN chown -R "$GRUMBLE_UID:$GRUMBLE_GID_NAME" \
             "$GRUMBLE_PATHS_CONFIG" \
             "$GRUMBLE_PATHS_DATABASE" \
             "$GRUMBLE_PATHS_BACKUP" \
             "$GRUMBLE_PATHS_MEDIA" \
             "$GRUMBLE_PATHS_LOGS"

# set directory permissions
RUN chmod -R 777 "$GRUMBLE_PATHS_CONFIG" \
                 "$GRUMBLE_PATHS_DATABASE" \
                 "$GRUMBLE_PATHS_BACKUP" \
                 "$GRUMBLE_PATHS_MEDIA" \
                 "$GRUMBLE_PATHS_LOGS"

# copy example config
COPY conf/example.toml "$GF_PATHS_CONFIG/server.toml"

# copy startup script
COPY ./build/docker/run.sh /run.sh
RUN chmod +x /run.sh

# copy binaries from build stage
COPY --from=builder /go/src/github.com/grumblechat/server/bin/*/grumble-server ./bin/
#COPY --from=builder /go/src/github.com/grumblechat/server/bin/*/grumble-cli ./bin/

# expose ports
EXPOSE 80

ENTRYPOINT [ "/run.sh" ]