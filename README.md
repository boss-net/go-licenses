# chklicenses

![app-pipeline](https://github.com/khulnasoft/chk-licenses/workflows/app-pipeline/badge.svg)
![Go Report Card](https://goreportcard.com/badge/github.com/khulnasoft/chk-licenses)

A go dependency license checker.

*This is thin a wrapper around [google's license classifier](https://www.github.com/google/licenseclassifier) forked from [go-license](https://www.github.com/google/go-licenses) with a few extra options.*

## Installation

```bash
# install the latest version to ./bin
curl -sSfL https://raw.githubusercontent.com/khulnasoft/chk-licenses/master/chklicenses.sh | sh 

# install a specific version to another directory
curl -sSfL https://raw.githubusercontent.com/khulnasoft/chk-licenses/master/chklicenses.sh | sh -s -- -b ./path/to/bin v1.26.0
```

## Usage

```bash
# list the licenses of all of your dependencies...
chklicenses list                        # ... from ./go.mod
chklicenses list ~/some/path            # ... from ~/some/path/go.mod
chklicenses list github.com/some/repo   # ... from a remote repo

# pass/fail of user-specified license restrictions (by .chklicenses.yaml)
chklicenses check
chklicenses check ~/some/path
chklicenses check github.com/some/repo
```

The `.chklicenses.yaml` can specify a simple allow-list or deny-list license name regex patterns (by SPDX name):

```bash
permit:
  - BSD.*
  - MIT.*
  - Apache.*
  - MPL.*
```

```bash
forbid:
  - GPL.*
```

```bash
ignore-packages:
  - github.com/some/repo
forbid:
  - GPL.*
```

Note: either allow or deny lists can be specified, not both.
