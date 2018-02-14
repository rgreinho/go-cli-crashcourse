# Golang crashcourse

## Goals

The goal is to learn the basic concepts of developing a project using the Golang ecosystem. We found a lot of
information on the web, but we had to dig left and right and glue them altogether.

To apply it to a simple, but real use case, we want to create a Go CLI client. For this tutorial we chose a simple REST
service: [ipify](https://www.ipify.org/).

Here are the steps that we will tackle about for this project:
1. Worskpace organization

  1.1 Scaffolding best practices (tools?)
2. Use cobra for the CLI
3. Lint the code
4. Dependency management (glide)
5. Debug the application (delve)
6. Write unit tests w/mocks (testify)
7. Use viper to read parameters from a configuration file
8. Write and publish documentation
9. Package the application

C'est parti!

## Prerequisites

### Install GO
```bash
brew install go
```

### Edit `~/.bash_profile`

Go expects you to use a unique workspace for all your projects. This workspace is usually located in `~/go`.
For convenience, we are going to update our  `~/.bash_profile` file with the following variables:
```bash
# Set Go variables.
export GOPATH=$(go env GOPATH)
export PATH=$PATH:$(go env GOPATH)/bin
```

## Workspace organization

We will start by creating a folder in our workspace for our application. According to the
[documentation](https://golang.org/doc/code.html#ImportPaths), it has to be located in
`${GOPATH}/src/github.com/rgreinho/go-cli-crashcourse`:
> If you keep your code in a source repository somewhere, then you should use the root of that source repository as your base path. For instance, if you have a GitHub account at github.com/user, that should be your base path.

```bash
GO_CRASH_COURSE="${GOPATH}/src/github.com/rgreinho/go-cli-crashcourse"
mkdir -p "${GO_CRASH_COURSE}"
cd "${GO_CRASH_COURSE}"
git init
```

### Project organization

```bash
.
├── LICENSE
├── cmd/
├── docs/
├── main.go
└── pkg/
```

### Resources

* [Filesystem Structure of a Go project](https://flaviocopes.com/go-filesystem-structure/)
* [Naming convention](https://talks.golang.org/2014/names.slide#1)
* [Organising Go Code](https://talks.golang.org/2014/organizeio.slide#11)
* [Project layout](https://github.com/golang-standards/project-layout)
* [Go Start](https://github.com/alco/gostart#motivation)
## CLI

Start by Configuring `cobra`. Putting the author name and the license is a good start:
```bash
cat << EOF > ~/.cobra.yaml
author: Rémy Greinhofer <remy.greinhofer@gmail.com>, Dashiel Lopez Mendez <hi@dashiel.me>
license: MIT
EOF
```
Refer to the "[configuring the cobra generator](https://github.com/spf13/cobra/blob/master/cobra/README.md)" section of
the official documentation for more details.

Now install `cobra`, initialize the project and create your first command:
```bash
go get -u github.com/spf13/cobra/cobra
cobra init
cobra add ip
```

### Resources

* [5 keys to create a killer CLI in Go](https://blog.alexellis.io/5-keys-to-a-killer-go-cli/)
* [Cobra](https://github.com/spf13/cobra)

## Linters

### goimports

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

## Dependency management

Install glide:
```bash
brew install glide
```

Initialize your repo and fetch the dependencies:
```bash
glide init
glide update
```

You should add the `vendor/` folder to your `.gitignore` file. It will help keep your repo small and all of your
dependencies are already well defined in `glide.yaml`.

And you can install your dependencies using `glide get` instead of `go get`:
```bash
glide get github.com/foo/bar
```

### Resources

* [Glide](https://glide.sh)

## Configuration file

### Resources

* [Viper](https://github.com/spf13/viper)

## Debugger - Delve

### Resources

* [Delve](https://github.com/derekparker/delve)

## Unit tests

Install testify:
```bash
glide get github.com/stretchr/testify
```

We're going to test the `api.go` file. In the `pkg` folder, create a file named `ipify_test.go`

### Resources

* [testify](https://github.com/stretchr/testify)
* <http://lucasfcosta.com/2017/01/11/Getting-Started-With-Testing-in-Go.html>

## Write and publish documentation

### Resources

* [GoDoc]

## Package the application

### Resources

## Resources

* [Learn](https://github.com/golang/go/wiki/Learn)
* [Go start](https://github.com/alco/gostart)
* [GolangBot tutorial series](https://golangbot.com/learn-golang-series/)
