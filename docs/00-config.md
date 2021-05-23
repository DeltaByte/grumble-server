# Configuration

The config object consists of several keys in the `core` section, and several sub-objects which are attached to the `core` config.
For example the `paths` config when represented in a JSON file would look like the example below, for `ENV` vars the hierarchy is replaced with an underscore,
e.g. `GRUMBLE_PATHS_DATABASE`.

```json
{
  "host": "127.0.0.1",
  "storage": {
    "database": "/data/database/"
  }
}
```

## Loading from environment variables

Grumble will try to load config from the OS environment using keys prefixed with `GRUMBLE_`. The precedence for loading config is listed below from highest to lowest.

1. Environment
2. Config file
3. Internal defaults

## Loading config from a different location

> **WARNING:** Adding the CLI flag will prevent Grumble from loading the default `config.json`

By default Grumble will try to load settings from `settings.json`, however, if you want to  load the config from a different location then you can do so by
starting the server whilst using the `config-file` flag. Grumble supports `YAML`, `JSON`, and `TOML` file formats. If you wish to separate your configs then
you can simply add the flag multiple times (earlier files have priority).

This example will only load config from the specified file.

```shell
grumble-server --config-file="/path/to/config.json"
```

This example will load from the custom file first and then fallback to the default file.

``` shell
grumble-server --config-file="/path/to/overrides.json" --config-file="config.json"
```

## Available Settings

_core:_

| Name   | Type   | Description               | Default Value | Validation |
| ------ | ------ | ------------------------- | ------------- | ---------- |
| port   | number | HTTP port                 | `8080`        |            |
| host   | string | Host Address              | `0.0.0.0`     |            |
| banner | bool   | Display banner on startup | `true`        |            |

_paths:_

| Name     | Type   | Description                                                              | Default Value        | Validation |
| -------- | ------ | ------------------------------------------------------------------------ | -------------------- | ---------- |
| database | string | Directory for storing the raw database files                             | `./storage/database` |            |
| media    | string | Directory for storing media assets, such as images embedded in a message | `./storage/media`    |            |
| logs     | string | Directory for log files                                                  | `./storage/logs`     |            |
| backup   | string | Directory for storing backups                                            | `./storage/backup`   |            |

_sentry:_

| Name   | Type   | Description                                                                                                                                                  | Default Value | Validation |
| ------ | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------- | ---------- |
| enable | bool   | Enable to automatic reporting of errors to [Sentry](https://sentry.io)                                                                                       | `true`        |            |
| dsn    | string | [Sentry DSN](https://docs.sentry.io/product/sentry-basics/dsn-explainer/), this is useful for if you do want error reporting, but only to a private instance | `0.0.0.0`     |            |

_backup:_

| Name     | Type   | Description                                                                                                       | Default Value | Validation |
| -------- | ------ | ----------------------------------------------------------------------------------------------------------------- | ------------- | ---------- |
| schedule | string | How often to run the backups, setting this ot zero with prevent backups from running.                             | `6h0m`        |            |
| amount   | number | How many backups to retain, e.g. keeping 28 backups if the schedule is `6h` will keep backups for 7 days.         | `28`          |            |
| group    | bool   | Enables grouping backups by day, butting each group into a subdirectory, e.g. `backups/2020-01-01/db-2001.tar.gz` | `false`       |            |
