# Releasing steps

This document contains all the steps I used to create the releases of **gophoner** under [releases tab](https://github.com/M4elstr0m/gophoner/releases/latest).

> [!CAUTION]
> You may not be allowed to publish your own version of the **gophoner**, please refer to the project's license.
> In this case, the following information is provided for reference only.

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

### System-level requirements

> [!IMPORTANT]
> The following steps are **one-time requirements** but still needed for [Project-specific requirements](#project-specific-requirements) (see below). 
> 
> You can skip these steps if you already setup everything for a previous project for example.

- Fork `microsoft/winget-pkgs` to `M4elstr0m/winget-pkgs`, then add a token with RW access to that fork and permission to open pull requests against the upstream repo.
- Create an empty `M4elstr0m/homebrew-tap` repository, then add a token with RW access to it.

### Project-specific requirements

> [!IMPORTANT]
> The following requirements are specific to this project. Follow these steps carefully.

Install `godotenv`
```sh
go install github.com/joho/godotenv/cmd/godotenv@latest
```

Install `goreleaser`
```sh
go install github.com/goreleaser/goreleaser/v2@latest
```

In `.env`, with a fine-grained access token scoped to the repository ([M4elstr0m/gophoner](https://github.com/M4elstr0m/gophoner)) with RW permission on "Contents"
```sh
GITHUB_TOKEN=github_pat_xxxxxxxx
```

Also in the `.env`, add the tokens you created for the other installation methods (see [System-level requirements](#system-level-requirements) above)
```sh
WINGET_GITHUB_TOKEN=github_pat_xxxxxxxx
HOMEBREW_TAP_GITHUB_TOKEN=github_pat_xxxxxxxx
```

### On-release requirements

> [!IMPORTANT]
> The following steps are needed **on each new release**.
>
> The repository must be public.

Create a tag on the repository (obviously bump the version each release)
```sh
git tag -a v1.0.0 -m "v1.0.0"
git push origin v1.0.0
```

Publish a new release
```sh
godotenv goreleaser release --clean
```