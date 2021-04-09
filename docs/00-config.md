# Configuration

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
grumble --config-file="/path/to/config.json"
```

This example will load from the custom file first and then fallback to the default file.

``` shell
grumble --config-file="/path/to/overrides.json" --config-file="config.json"
```
