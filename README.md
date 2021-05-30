# Grumble server

Example configs are included in the `conf` directory, full documentation is in the `docs` directory.

## Simple startup

Copy example config

```shell
cp ./conf/example.toml ./config.toml
```

build and start server
```shell
make start
```

## Acknowledgements

Full list of dependencies is available in `go.mod`, though are worthy of special mention:

- [BoltDB](https://github.com/etcd-io/bbolt)
- [KSUID](https://github.com/segmentio/ksuid)

Some implementations inspired by:

- [Emitter](https://github.com/emitter-io/emitter)
- [Grafana](https://github.com/grafana/grafana)
