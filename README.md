# kli

This repo contains a CLI to access, query and manage KRE and KDL.


## Frameworks and libraries

- [spf13/cobra](https://github.com/spf13/cobra) used as CLI framework. 
- [golangci-lint](https://golangci-lint.run/) as linters runner.


## Development

You can compile the binary with this command: 

```bash
go build -o kli cmd/main.go
```

And then test run any command: 
```bash
./kli help

# Output: 
Use Konstellation API from the command line.

Usage:
  kli [command]

Available Commands:
  kre         Manage KRE
  server      Manage servers for kli

Flags:
  -h, --help   help for kli

Use "kli [command] --help" for more information about a command.

```

Example: 

```bash
./kli server ls

# Output
SERVER URL                                  
local* http://api.kre.local                 
int    https://api.your-domain.com 
```


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

## Creating Releases

The build system automatically compiles cross-platform binaries to any git tag named vX.Y.Z. 

To test out the build system, publish a prerelease tag with a name such as vX.Y.Z-pre.0 or vX.Y.Z-rc.1. 
Note that such a release will still be public, but it will be marked as a "prerelease", meaning that it won't show up as
a "latest" release.


### Tagging a new release

`git tag v1.2.3 && git push origin v1.2.3`

Wait several minutes for builds to run: https://github.com/konstellation-io/kli/actions
Verify release is displayed and has correct assets: https://github.com/konstellation-io/kli/releases

(Optional) Delete any pre-releases related to this release


### Release locally for debugging

A local release can be created for testing without creating anything official on the release page.

- Make sure [GoReleaser](https://goreleaser.com/install/) is installed
- Run: 
    `goreleaser --skip-validate --skip-publish --rm-dist`
- Find the built binaries under `dist/` folder.
