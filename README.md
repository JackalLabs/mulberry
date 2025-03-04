# Jackal EVM Bridge

## Init
To init your system, you can get a new address like this.
```shell
mulberry wallet address
```
This will also create the `~/.mulberry` directory and all the config files. You can adjust where this goes with the `home` flag.

## Config
For sample configuration files, see [DEPLOY.md](DEPLOY.md). Other EVM networks can be added as `networks_config` entries.

## Testing

Run `./scripts/test.sh` to start a test environment.