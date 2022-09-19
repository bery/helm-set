# helm-set
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![GitHub release](https://img.shields.io/github/v/release/bery/helm-set.svg)](https://github.com/bery/helm-set/releases)
[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/helm-helm-set)](https://artifacthub.io/packages/search?repo=helm-helm-set)

This Helm plugin allows you to substitute the environment variables specified in your helm values file with their respective values in the environment from within a CICD pipeline.

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
export HELM_VAR_replicaCount=3
helm set upgrade --install --dry-run=false xxx ealenn/echo-server
```
### Result
```
```bash
helm upgrade upgrade --install --dry-run=false xxx ealenn/echo-server --set replicaCount=3
```
```
## Uninstall
```bash
helm plugin remove helm-set
```

## License

[MIT](LICENSE)
# helm-set
