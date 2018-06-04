![Civil Logo](docs/civil_logo_white.png?raw=true)

---
[Civil](https://joincivil.com/) is a decentralized and censorship resistant ecosystem for online Journalism. Read more in our whitepaper.

This repository contains open-source code to capture and handle Civil-specific smart contract event log data. It is written in `golang`. It currently captures Civil TCR and Civil Newsroom related events, but can be expanded to capture additional events.

For Civil's main open-source tools and packages, check out [http://github.com/joincivil/Civil](http://github.com/joincivil/Civil).

## Contributing

Civil's ecosystem is free and open-source, we're all part of it and you're encouraged to be a part of it with us.  We are looking to evolve this into something the community will find helpful and effortless to use.

If you're itching to dwelve deeper inside, [**help wanted**](https://github.com/joincivil/civil-events-crawler/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22)
and [**good first issue**](https://github.com/joincivil/civil-events-crawler/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22) labels are good places to get started and learn the architecture.

## Install Requirements

This project is using `make` to run setup, builds, tests, etc.  

Ensure that your `$GOPATH` and `$GOROOT` are setup properly in your shell configuration and that this repo is cloned into the appropriate place in the `$GOPATH`. i.e. `$GOPATH/src/github.com/joincivil/civil-events-crawler/`

To setup the necessary requirements:

```
make setup
```

### Dependencies

Uses `vgo` for dependency management, although we may keep a `/vendor/` and use `go get` for backwards compatibility with `go` and some tooling may still rely on `$GOPATH`.  

**NOTE: Some decisions and detail needed here**

## Code Generation

There are a few places where code/artifacts need to be moved or generated before the project can be built, tested, and/or linted.  This is likely a place that can be streamlined and improved as time goes on.

### Contract .abi & .bin

This project relies on artifacts from the main Civil repository [http://github.com/joincivil/Civil](http://github.com/joincivil/Civil).  Please clone the Civil repository into a directory accessible by this repository.

To build and copy the Civil contract .abi/.bin files, run the `scripts/abi_retrieval.sh` script:

```
scripts/abi_retrieval.sh /full/path/to/main/civil/repo /full/path/to/civil-events-crawler/abi
```

The destination directory is generally the `civil-events-crawler/abi` directory in this repository.  This will produce the `.abi` & `.bin` files from the artifacts in the `Civil` repository.

### Solidity Wrappers

There are Solidity wrappers that are created by `abigen` from the `go-ethereum` package.  These will go into `pkg/generated/<contract>` directories. These are generated by running:

```
make generate-contracts
```

This is run along with the `make lint`, `make build`, `make test` commands below.

### Contract Watchers

There is a number of `Watch*` methods for each Civil Solidity contract wrapper that allow us to listen and stream contract events.  The wrappers around these `Watch*` methods are generated using the `cmd/watchgen` command with `go generate`.  These will be placed into the `pkg/listener` directory.

```
make generate-watchers
```
This is run along with the `make lint`, `make build`, `make test` commands below.


## Lint

Check all the packages for linting errors using a variety of linters via `gometalinter`.  Check the `Makefile` for the up to date list of linters.

```
make lint
```

## Build


```
make build
```

## Testing

Runs the tests and checks code coverage across the project.  Produces a `coverage.txt` file for use later.

```
make test
```

## Code Coverage Tool

Run `make test` and launches the HTML code coverage tool.

```
make cover
```


