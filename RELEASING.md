# Releasing steps

This document contains all the steps I used to create the releases of **gophoner** under [releases tab](https://github.com/M4elstr0m/gophoner/releases/latest).

## Pre-release testing

Use these commands to create dev builds and test it locally.

```sh
make run ARGS="help"
make run ARGS="interactive"
# ...

# or

make build
```

## Release

Install `godotenv`
```sh
go install github.com/joho/godotenv/cmd/godotenv@latest
```

Install `goreleaser`
```sh
go install github.com/goreleaser/goreleaser/v2@latest
```

In `.env`, with a fine-grained access token scoped to the repository with RW permission on "Contents"
```sh
GITHUB_TOKEN=github_pat_xxxxxxxx
```

Create a tag on the repository (obviously bump the version each release)
```sh
git tag -a v1.0.0 -m "v1.0.0"
git push origin v1.0.0
```

Publish a new release
```sh
godotenv goreleaser release --clean
```