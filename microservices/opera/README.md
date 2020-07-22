# Opera Service

This is the Opera service

Generated with

```
micro new --namespace=fuckopera --type=service opera
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: fuckopera.service.opera
- Type: service
- Alias: opera

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./opera-service
```

Build a docker image
```
make docker
```