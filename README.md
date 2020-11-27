# kli

This repo contains a CLI to access, query and manage KRE and KDL.


## Frameworks and libraries

- [golangci-lint](https://golangci-lint.run/) as linters runner.


## Run tests

```sh
go test ./...
```


## Linters

`golangci-lint` is a fast Go linters runner. It runs linters in parallel, uses caching, supports yaml config, has
integrations with all major IDE and has dozens of linters included.

As you can see in the `.golangci.yml` config file of this repo, we enable more linters than the default and
have more strict settings.

To run `golangci-lint` execute:
```
golangci-lint run
```
