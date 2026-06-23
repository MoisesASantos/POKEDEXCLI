# Pokedex CLI

A command-line Pokédex built in **Go (Golang)** that interacts with the PokéAPI. This project was developed as a practical introduction to Go, focusing on API consumption, caching, concurrency, data modeling, and CLI application design.

## Features

* Browse Pokémon location areas
* Navigate forward and backward through paginated results
* Explore a location area and discover which Pokémon can be found there
* Catch Pokémon using a probability-based capture system
* Store captured Pokémon in a personal Pokédex
* Inspect detailed information about captured Pokémon
* In-memory caching system to reduce unnecessary API requests
* Automatic cache cleanup using goroutines and mutexes

---

## Commands

### Help

Display all available commands.

```bash
help
```

### Map

Display the next page of location areas.

```bash
map
```

### Map Back

Display the previous page of location areas.

```bash
mapb
```

### Explore

Explore a location area and list all Pokémon encounters.

```bash
explore canalave-city-area
```

Example output:

```text
Exploring canalave-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
 - magikarp
```

### Catch

Attempt to catch a Pokémon.

```bash
catch pikachu
```

Example output:

```text
Throwing a Pokeball at pikachu...
pikachu was caught!
```

or

```text
Throwing a Pokeball at pikachu...
pikachu escaped!
```

### Inspect

Display detailed information about a previously caught Pokémon.

```bash
inspect pikachu
```

Example output:

```text
Name: pikachu
Height: 4
Weight: 60
Stats:
 - hp: 35
 - attack: 55
 - defense: 40
 - special-attack: 50
 - special-defense: 50
 - speed: 90
Types:
 - electric
```

### Pokedex

List all captured Pokémon.

```bash
pokedex
```

Example output:

```text
Your Pokedex:
 - pikachu
 - charmander
 - squirtle
```

---

## Project Structure

```text
.
├── main.go
├── repl/
│   └── clean_input.go
├── pokedex/
│   ├── client.go
│   ├── commands.go
│   ├── registry.go
│   └── types.go
├── internal/
│   └── pokecache/
│       └── cache.go
└── go.mod
```

### Components

#### CLI Layer

Responsible for:

* Reading user input
* Parsing commands
* Executing command callbacks

#### Pokedex Package

Contains:

* API communication
* Command implementations
* Data structures
* Pokémon management logic

#### Cache Package

Custom in-memory cache implementation featuring:

* Thread-safe access using `sync.Mutex`
* Configurable expiration interval
* Automatic cleanup through goroutines
* Raw response storage as `[]byte`

---

## Technologies Used

* Go
* PokéAPI
* net/http
* encoding/json
* goroutines
* sync.Mutex
* math/rand

---

## Learning Objectives

This project was designed to practice:

* Consuming REST APIs
* JSON serialization and deserialization
* Struct modeling
* Package organization
* Interfaces
* Error handling
* Concurrency primitives
* State management
* Building interactive CLI applications

---

## Running the Project

Clone the repository:

```bash
git clone git@github.com:MoisesASantos/POKEDEXCLI.git
cd POKEDEXCLI
```

Run the application:

```bash
go run main.go
```

Or build an executable:

```bash
go build
./POKEDEXCLI
```

---

## Example Session

```text
Pokedex > map

canalave-city-area
eterna-city-area
pastoria-city-area

Pokedex > explore canalave-city-area

Exploring canalave-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
 - magikarp

Pokedex > catch tentacool

Throwing a Pokeball at tentacool...
tentacool was caught!

Pokedex > inspect tentacool

Name: tentacool
Height: 9
Weight: 455
Stats:
 - hp: 40
 - attack: 40
 - defense: 35
Types:
 - water
 - poison
```

---

## What I Learned

Building this project helped me gain hands-on experience with:

* Go's type system and structs
* Interfaces and abstraction
* Working with external APIs
* Concurrency with goroutines
* Synchronization with mutexes
* Designing maintainable CLI applications
* Organizing medium-sized Go projects into packages

As my first Go project, it provided a strong foundation for future backend, systems programming, and distributed systems projects.

---

## API Reference

Data provided by:

[PokéAPI](https://pokeapi.co?utm_source=chatgpt.com)

A free and open Pokémon RESTful API used for educational and personal projects.
