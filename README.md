<!--<p align="center">
  <a href="https://github.com/federicotorres233">
    <picture>
      <source height="125" media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/gofiber/docs/master/static/img/logo-dark.svg">
      <img height="125" alt="Fiber" src="https://raw.githubusercontent.com/gofiber/docs/master/static/img/logo.svg">
    </picture>
  </a>
  <br>-->
<p align="center">
  <em><b>GoKeys</b> is a password database manager for the terminal written in Go and built using <a href="https://github.com/spf13/cobra">Cobra</a>, a library for creating CLI applications.</em>
</p>

---

## ⭐️ To do

- ~~Autocopy to clipboard~~
- ~~Add -—install~~
- ~~Remove logs && log everything to a logfile~~
- ~~Use SQLite~~
- ~~Don't echo passwords~~
- Add dockerfile
- Fancier terminal
- Organize sensitive files and change permissions
- Encrypt database
- Concurrency for faster performance
- Add more options
- Get by other options
- CRUD
- Tests
- Reorganise folders and files
- Select database location
- More output

## ⚙️ Installation

The easiest way to start using GoKeys is to get the latest binary from the github repository.

```bash
wget -qO- https://github.com/FedericoTorres233/gokeys/releases/download/v0.1.0/gokeys.tar.gz | gzip -d | tar xvfz -
```

Done! You can start using it:

```bash
gokeys --install # Install password database
```

### Installing from source

GoKeys requires **Go version `1.20` or higher** to run. If you need to install or upgrade Go, visit the [official Go download page](https://go.dev/dl/). Clone the repository and cd into it:

```bash
git clone https://github.com/FedericoTorres233/gokeys && cd gokeys
```

Install project dependencies:

```bash
go mod tidy
```

Next, build the project from source. Be sure you have the `make` command installed:

```
make
```

Done! You can start using it:

```bash
./bin/gokeys --install # Install password database
```

## ⚡️ Quickstart

Here are basic examples of usage: 

```go
package main

import (
    "log"

    "github.com/gofiber/fiber/v3"
)

func main() {
    // Initialize a new Fiber app
    app := fiber.New()

    // Define a route for the GET method on the root path '/'
    app.Get("/", func(c fiber.Ctx) error {
        // Send a string response to the client
        return c.SendString("Hello, World 👋!")
    })

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}
```

## 🎯 Features

-   Something [Something](example.com)

## 👍 Contribute

If you want to say **Thank You** and/or support the project:

1. Add a [GitHub Star](https://github.com/federicotorres233/gokeys/stargazers) to the project.
2. Submit a [pull request](https://github.com/FedericoTorres233/gokeys/pulls) to add a new feature
3. Create an [issue](https://github.com/FedericoTorres233/gokeys/issues) if you find any bug
