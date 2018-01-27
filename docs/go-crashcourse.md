# Golang crashcourse

## Goals

The goal is to learn the basic concepts of developing a project using the Golang ecosystem. We found a lot of
information on the web, but we had to dig left and right and glue them altogether.

To apply it to a simple, but real use case, we want to create a Go CLI client. For this tutorial we chose a simple REST
service: [ipify](https://www.ipify.org/)

Here are the steps that we will tackle about for this project:
1. Worskpace organization

  1.1 Scaffolding best practices (tools?)

2. Use cobra for the CLI
3. Lint the code

2. Dependency management (glide)
5. Debug the application (delve)
6. Write unit tests w/mocks (testify)

4. Use viper to read parameters from a configuration file
8. Write and publish documentation
9. Package the application

C'est parti!

## Prerequisites

1. Install GO


```bash
brew install go
```

2. Edit `~/.bash_profile`

```bash
# Set Go variables.
export GOPATH=$(go env GOPATH)
export PATH=$PATH:$(go env GOPATH)/bin
```

## Workspace organization

### Project organization

```bash
GO_CRASH_COURSE="${GOPATH}/src/github.com/rgreinho/go-cli-crashcourse"
mkdir -p "${GO_CRASH_COURSE}"
cd "${GO_CRASH_COURSE}"
git init
go get -u github.com/spf13/cobra/cobra
```

### Configure Cobra

See <https://github.com/spf13/cobra/blob/master/cobra/README.md>

```bash
cat << EOF > ~/.cobra.yaml
author: Rémy Greinhofer <remy.greinhofer@gmail.com>, Dashiel Lopez Mendez <hi@dashiel.me>
license: MIT
EOF
```

```bash
cobra init
cobre add ip
```

### Resources

* [Filesystem Structure of a Go project](https://flaviocopes.com/go-filesystem-structure/)
* [Naming convention](https://talks.golang.org/2014/names.slide#1)
* [Organising Go Code](https://talks.golang.org/2014/organizeio.slide#11)
* [Project layout](https://github.com/golang-standards/project-layout)

## Dependency management

## CLI

### Resources

* [5 keys to create a killer CLI in Go](https://blog.alexellis.io/5-keys-to-a-killer-go-cli/)
* [Cobra](https://github.com/spf13/cobra)

## Linters

### goimports command

```bash
go get golang.org/x/tools/cmd/goimports
goimports -w .
```

### go vet

### golint

```
go get -u github.com/golang/lint/golint
golint
```

### Go cyclo

```
go get github.com/fzipp/gocyclo
gocyclo .
```

Note: Try to get a percentage score

### Resources

* [Go report card](https://goreportcard.com/)
* [Go imports](https://godoc.org/golang.org/x/tools/cmd/goimports)

## Configuration file

### Resources

* [Viper](https://github.com/spf13/viper)

## Debugger - Delve

### Resources

* [Delve](https://github.com/derekparker/delve)

## Unit tests

### Resources

* [testify](https://github.com/stretchr/testify)


## Write and publish documentation

### Resources

* [GoDoc]

## Package the application

### Resources

## Resources

* [Learn](https://github.com/golang/go/wiki/Learn)
* [Go start](https://github.com/alco/gostart)
* [GolangBot tutorial series](https://golangbot.com/learn-golang-series/)
