# LightRiders

LightRiders is a classic multiplayer action game. You can either play with your
opponent on the same computer (`hotseat`), or play through the network.

## Installation

Need to have Go installed. After that clone the repository

```
git clone https://github.com/alexander-lazarov/lightriders
go mod download
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
