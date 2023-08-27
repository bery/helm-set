# helm-set
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![GitHub release](https://img.shields.io/github/v/release/bery/helm-set.svg)](https://github.com/bery/helm-set/releases)
[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/helm-helm-set)](https://artifacthub.io/packages/search?repo=helm-helm-set)

Helm-set plugin allows you to substitute the environment variables specified in your helm values file with their respective values in the environment from within a CICD pipeline. Heavily inspired by Terraform's [envsubst](https://developer.hashicorp.com/terraform/language/values/variables#environment-variables) feature.

## Install

The installation itself is simple as:

```bash
helm plugin install https://github.com/bery/helm-set.git
```
You can install a specific release version:
```bash
helm plugin install https://github.com/bery/helm-set.git --version <release version>
```

To use the plugin, you do not need any special dependencies. The installer will download the latest release with prebuilt binary from [GitHub releases](https://github.com/bery/helm-set/releases).

## Usage

### Simple usage
```bash
Use environment variables as helm values by taking all variables starting with HELM_VAR_ to --set values.

Usage:
  set [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  install     wrapper for helm install, decrypting secrets
  upgrade     wrapper for helm upgrade, decrypting secrets

Flags:
  -d, --debug     enable verbose output
      --dry-run   parameter parsing mode (rename, copy)
  -h, --help      help for set
  -v, --verbose   enable verbose output
```
### Examples
#### Basic usage
Set the `HELM_VAR_replicaCount` envrionment variable to `3` by:
```bash
export HELM_VAR_replicaCount=3
```
Replace `helm upgrade` with `helm set upgrade` and run:
```bash
helm set upgrade --install --dry-run=false <name> <chart>
```
This will replace the value of `replicaCount` in the `helm upgrade` command with the value of the environment variable `HELM_VAR_replicaCount`.
#### Result
```bash
helm upgrade upgrade --install --dry-run=false <name> <chart> --set replicaCount=3
```
#### Nested values
Goal: set value of image.tag to `latest` and value of image.pullPolicy to `Always`.
```bash
export HELM_VAR_image__tag="latest"
export HELM_VAR_image__pullPolicy="Always"
helm set upgrade --install --dry-run=false <name> <chart>
```
#### Result
```bash
helm upgrade upgrade --install <name> <chart> --set image.tag="latest" --set image.pullPolicy="Always"
```
#### Lists values
```bash
export HELM_VAR_list__item___0___path="your_path"
helm set upgrade --install --dry-run=false <name> <chart>
```
#### Result
```bash
helm upgrade upgrade --install <name> <chart> --set list.item[0].path="your_path"
```

## Uninstall
```bash
helm plugin remove helm-set
```

## License

[MIT](LICENSE)