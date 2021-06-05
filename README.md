# golang-basic-project

A tiny live demo of basic Golang project.

## Badges

[![Test](https://github.com/rnazmolab/golang-basic-project/actions/workflows/test.yml/badge.svg)](https://github.com/rnazmolab/golang-basic-project/actions/workflows/test.yml)

## Purpose of this repo

Create a live demo of basic Golang project **for my own reference**.

## Contents

- Go1.16
- Go Modules
- Tests
- GitHub Actions
- VS Code

## TODO

- Use golangci-lint
  - [x] Add  new workflow `lint` with golangci-lint
  - [ ] Add `.golangci.yml`
  - [ ] Add a new script`/devel-tools/install-golangci-lint.sh`. (ref: https://golangci-lint.run/usage/install/#linux-and-windows)
  - [ ] Add a new command `lint` to Makefile (ref: https://golangci-lint.run/usage/quick-start/)
- Use Dependabot
  - Ref: https://dependabot.com/go/
  - https://songmu.jp/riji/entry/2020-03-23-dependabot.html
  - https://sue445.hatenablog.com/entry/2021/03/20/075551
  - https://shogo82148.github.io/blog/2021/03/17/actions-check-permissions/
- Add commands to Makefile
  - lint
  - bench
  - (build)
  - cross-build
  - install
  - coverage
  - bump
  - Use Codecov
- Add benchmark
- Use goreleaser
- Add `/tools/tools.go`
- Add `.gitattributes`
- Add badges
  - `go report`: https://goreportcard.com/
  - `refarence`: https://go.dev/ (via https://pkg.go.dev/badge/)
  - `coverage`: (ref: https://shields.io/category/coverage)

## Memo

### 'Golang Standard Project Layout'

Read: https://github.com/golang-standards/project-layout/issues/117#issuecomment-828503689

## Env

```console
$ property         
property v0.1.1 - A tiny Bash script to get OS and other
software version info. https://github.com/rnazmo/property
============================================================
OS NAME      : Kali GNU/Linux Rolling
OS VERSION   : 2021.2
Default Shell: Zsh
Bash VERSION : 5.1.4(1)-release (x86_64-pc-linux-gnu)
Zsh VERSION  : 5.8 (x86_64-debian-linux-gnu)
CPU ARCH     : x86_64
============================================================

$ go version                      
go version go1.16 linux/amd64
```

## Resources

TODO:

## Ref

TODO:
