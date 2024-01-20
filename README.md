## Installation

#### Homebrew

```
brew tap yana1205-test/sample-app https://github.com/yana1205-test/sample-app
brew install yana1205-test-sample-app
```

When you upgrade
```
brew update
brew upgrade yana1205-test-sample-app
```

## Quickstart

## Build at local
```
goreleaser release --snapshot --clean
```

## Release
```
GITHUB_TOKEN=<PAT> goreleaser release
```