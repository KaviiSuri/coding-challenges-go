# Reddish

Reddish is an attempt to make a redis-server clone. The purpose is not to replace or challenge redis, rather to understand it's inner working and learn socket programming, concurrency and parsing from the project.

## Dependencies
This project expects the `redis-cli` and `redis-benchmark` scripts to already exist on your system for certain use cases, like running the client, or running a benchmark.

## Usage
**To Build**
```sh
make build
```

**To Run**
```sh
make run

**Connect with Redis CLI**
```sh
make redis-cli

**To Run the Benchmark**
```sh
./scripts/run-benchmark.sh
```

## Features
This project only supports very basic SET and GET commands as of now.

