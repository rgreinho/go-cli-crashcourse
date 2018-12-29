# Golang crashcourse

## Goals

The goal is to learn the basic concepts of developing a project using the Golang ecosystem. We found a lot of
information on the web, but we had to dig left and right and glue them altogether.

To apply it to a simple, but real use case, we want to create a Go CLI client. For this tutorial we chose a simple REST
service: [ipify](https://www.ipify.org/).

Here are the steps that we will tackle about for this project:
1. Worskpace organization
    1. Scaffolding best practices (tools?)
2. Dependency management (glide)
3. Lint the code
4. Write unit tests w/mocks (testify)
5. Setup a CI (CircleCI)
6. Use cobra for the CLI
7. Use viper to read parameters from a configuration file
8. Debug the application (delve)
9. Write and publish documentation
10. Package the application

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

Refer to the link in the resource section to learn more about other standard directories.

The outputs of the build will go to `/dist`.


### Project organization

```bash
.
├── LICENSE
├── Makefile
├── README.md
├── .gitignore
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

## Dependency management

To manage the dependencies we are using `Go dep`:
```
brew install dep
```

Initialize your repo and fetch the dependencies:
```
dep init
```

You can run dep ensure to have `dep` import any new dependencies that you would need.

You should add the `vendor/` folder to your `.gitignore` file. It will help keep your repo small and all of your
dependencies are already well defined in `Gopkg.toml`.

And you can install your dependencies using `dep ensure -add`:
```bash
dep ensure -add github.com/foo/bar
```

### Resources

* [Dep](https://github.com/golang/dep)

## Linters

To lint our project we will use [GoMetaLinter](https://github.com/alecthomas/gometalinter).

```
brew tap alecthomas/homebrew-tap
brew install gometalinter
```

Create a `.gometalinter.json` configuration file
```json
{
  "Enable": [
    "errcheck",
    "goimports",
    "golint",
    "vet"
  ]
}
```

TODO: Check if we need to add more linters.

### Resources

* [Go report card](https://goreportcard.com/)
* [Go imports](https://godoc.org/golang.org/x/tools/cmd/goimports)
* [Go lint](https://github.com/golang/lint)
* [Go vet](https://golang.org/cmd/vet/)
* [Go Style Common Mistakes](https://golang.org/s/style)

## Unit tests

Install testify:
```bash
glide get github.com/stretchr/testify
```

We're going to test the `api.go` file. In the `pkg` folder, create a file named `ipify_test.go`

### Resources

* [testify](https://github.com/stretchr/testify)
* <http://lucasfcosta.com/2017/01/11/Getting-Started-With-Testing-in-Go.html>

## Setup a CI (CircleCI)

Create a configuration file in `.circleci/config.yml`.

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

## Configuration file

### Resources

* [Viper](https://github.com/spf13/viper)

## Debugger - Delve

### Resources

* [Delve](https://github.com/derekparker/delve)

## Write and publish documentation

### Resources

* [GoDoc]

## Package the application

### Resources

## Resources

* [Learn](https://github.com/golang/go/wiki/Learn)
* [Go start](https://github.com/alco/gostart)
* [GolangBot tutorial series](https://golangbot.com/learn-golang-series/)
