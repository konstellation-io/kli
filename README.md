[![Tests][tests-badge]][tests-link]
[![GitHub Release][release-badge]][release-link]
[![Go Report Card][report-badge]][report-link]
[![License][license-badge]][license-link]
[![Coverage][coverage-badge]][coverage-link]

# kli

This repo contains a CLI to access, query and manage KRE and KDL.


## Frameworks and libraries

- [gomock](https://github.com/golang/mock) a mock library.
- [spf13/cobra](https://github.com/spf13/cobra) used as CLI framework.
- [joho/godotenv](https://github.com/joho/godotenv) used to parse env files.
- [golangci-lint](https://golangci-lint.run/) as linters runner.


## Development

You can compile the binary with this command: 

```bash
go build -o kli cmd/main.go
```

Then run any command: 
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


### Setting Version variables

1. You can set a Version variable as a key/value pair directly: 

```bash
./kli kre version config version-id-123456 --set SOME_VAR="any value"
# Output:
# [✔] Config updated for version 'version-id-123456'.
```

2. Add a value from an environment variable:

```bash
export SOME_VAR="any value"
./kli kre version config version-id-123456 --set-from-env SOME_VAR
# Output:
# [✔] Config updated for version 'version-id-123456'.
```

3. Add multiple variables from a file:

```text
# variables.env file
SOME_VAR=12345
ANOTHER_VAR="some long string... "
```

```bash
./kli kre version config version-id-123456 --set-from-file variables.env
# Output:
# [✔] Config updated for version 'version-id-123456'.
```

NOTE: `godotenv` library currently doesn't support multiline variables, as stated in
[PR #118 @godotenv](https://github.com/joho/godotenv/pull/118). Use next example as a workaround. 


4. Add a file as value:

```bash
export SOME_VAR=$(cat any_file.txt) 
./kli kre version config version-id-123456 --set-from-env SOME_VAR
# Output:
# [✔] Config updated for version 'version-id-123456'.
```



## Testing

To create new tests install [GoMock](https://github.com/golang/mock). Mocks used on tests are generated with 
**mockgen**, when you need a new mock, add the following:

```go
//go:generate mockgen -source=${GOFILE} -destination=$PWD/mocks/${GOFILE} -package=mocks
```

To generate the mocks execute:
```sh
$ go generate ./...
```

### Run tests

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




<!-- JUST BADGES & LINKS -->
[tests-badge]: https://img.shields.io/github/workflow/status/konstellation-io/kli/Test
[tests-link]: https://github.com/konstellation-io/kli/actions?query=workflow%3ATest

[release-badge]: https://img.shields.io/github/release/konstellation-io/kli.svg?logo=github&labelColor=262b30
[release-link]: https://github.com/konstellation-io/kli/releases

[report-badge]: https://goreportcard.com/badge/github.com/konstellation-io/kli
[report-link]: https://goreportcard.com/report/github.com/konstellation-io/kli

[license-badge]: https://img.shields.io/github/license/konstellation-io/kli
[license-link]: https://github.com/konstellation-io/kli/blob/master/LICENSE

[coverage-badge]: https://sonarcloud.io/api/project_badges/measure?project=kli&metric=coverage
[coverage-link]: https://sonarcloud.io/dashboard?id=kli
