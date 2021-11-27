# go-cookiecutter

[![build status](https://github.com/tech-thinker/go-cookiecutter/workflows/Go%20Tests/badge.svg)](https://github.com/tech-thinker/go-cookiecutter/blob/master/.github/workflows/go-tests.yaml)
[![CodeFactor](https://www.codefactor.io/repository/github/tech-thinker/go-cookiecutter/badge/main?s=752e2e187b014d8c50de1ad75867866bdc43aa73)](https://www.codefactor.io/repository/github/tech-thinker/go-cookiecutter/overview/main)
[![go version](https://img.shields.io/github/go-mod/go-version/tech-thinker/go-cookiecutter.svg)](https://github.com/tech-thinker/go-cookiecutter/blob/master/go.mod)
[![GoReportCard](https://goreportcard.com/badge/github.com/tech-thinker/go-cookiecutter)](https://goreportcard.com/badge/github.com/tech-thinker/go-cookiecutter)
[![Code style: standard](https://img.shields.io/badge/code%20style-standard-blue.svg)]()
[![License: MIT](https://img.shields.io/badge/License-MIT%20-blue.svg)](https://github.com/tech-thinker/go-cookiecutter/blob/master/LICENSE)
[![Github all releases](https://img.shields.io/github/downloads/tech-thinker/go-cookiecutter/total.svg)](https://github.com/tech-thinker/go-cookiecutter/releases/)

## Overview

Go-CookieCutter is an starter project for Golang micro service application

## Purpose

With the template, you can easily start a backend server development with a little tweaks in the code.

## Requirements

- [Golang](https://golang.org/dl/)
- [Docker](https://docs.docker.com/get-docker/) (Optional but recommended)
- [Postgresql](https://www.postgresql.org/download/) (**Only if you are not using `docker`**)

[Note: you can you use any sql database but required to change the sql driver in `config` as per the database server.]

## How to create project using this template?

_Coming Soon_

## Installation

Clone the repository,

```sh
git clone https://github.com/tech-thinker/go-cookiecutter.git
```

There is no need to install dependencies if you are using `docker`.

Install required dependencies,

```sh
cd go-cookiecutter
go mod tidy
```

## Migration

_Update soon_

## How to run?

Defining environment variables,

```sh
cp docker/.env{.sample,}

```

Edit `.env` with your favorite editor.

Export those variables,

```sh
export $(cat docker/.env | xargs)
```

Using Docker,

```sh
cd docker
docker-compose up --build
```

Test running,

```sh
curl http://localhost:3000/ping
```

# Contribute

- [Contributing](CONTRIBUTING.md)
- [Code of Conduct](CODE_OF_CONDUCT.md)

# License

- open sourced under [MIT](LICENSE) License
