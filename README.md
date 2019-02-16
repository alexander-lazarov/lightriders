# LightRiders

LightRiders is a classic multiplayer action game. You can either play with your
opponent on the same computer (`hotseat`), or play through the network.

## Installation

Need to have Go supproting modules installed (>=1.11). After making sure that
you have the right `go` version installed (`go version`), run in a terminal:

```
git clone https://github.com/alexander-lazarov/lightriders
go mod download

# build
go build

# or just run
go run main.Go --help
```

## Running tests
```
go test ./...
```

## Build

```
go build
```

## Run

```
lightriders [hotseat]       - 2 players on one computer
lightriders join            - host a multiplayer game
lightriders host <hostname> - join a multiplayer game
```

## Dependencies

Run `go mod download` to download dependencies.

## TODOs

- [x] Basic data types
- [ ] Repo stuff - pre-commit hooks
- [x] Hook up ebitien game engine
- [x] Sort out networking
