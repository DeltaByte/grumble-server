#!/bin/sh -e

PERMISSIONS_OK=0

if [ ! -r "$GRUMBLE_PATHS_CONFIG" ]; then
    echo "GRUMBLE_PATHS_CONFIG='$GRUMBLE_PATHS_CONFIG' is not readable."
    PERMISSIONS_OK=1
fi

if [ ! -w "$GRUMBLE_PATHS_DATABASE" ]; then
    echo "GRUMBLE_PATHS_DATABASE='$GRUMBLE_PATHS_DATABASE' is not writable."
    PERMISSIONS_OK=1
fi

if [ ! -w "$GRUMBLE_PATHS_BACKUP" ]; then
    echo "GRUMBLE_PATHS_BACKUP='$GRUMBLE_PATHS_BACKUP' is not writable."
    PERMISSIONS_OK=1
fi

if [ ! -w "$GRUMBLE_PATHS_MEDIA" ]; then
    echo "GRUMBLE_PATHS_MEDIA='$GRUMBLE_PATHS_MEDIA' is not writable."
    PERMISSIONS_OK=1
fi

if [ ! -r "$GRUMBLE_PATHS_HOME" ]; then
    echo "GRUMBLE_PATHS_HOME='$GRUMBLE_PATHS_HOME' is not readable."
    PERMISSIONS_OK=1
fi

if [ $PERMISSIONS_OK -eq 1 ]; then
    echo "You may have issues with file permissions."
fi

export HOME="$GRUMBLE_PATHS_HOME"

exec grumble-server --config-file="$GRUMBLE_PATHS_CONFIG/server.toml"