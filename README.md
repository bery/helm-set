# helm-env
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![GitHub release](https://img.shields.io/github/v/release/bery/helm-env.svg)](https://github.com/bery/helm-env/releases)
[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/helm-helm-env)](https://artifacthub.io/packages/search?repo=helm-helm-env)

This Helm plugin allows you to substitute the environment variables specified in your helm values file with their respective values in the environment from within a CICD pipeline.

## Install

The installation itself is simple as:

```bash
helm plugin install https://github.com/bery/helm-env.git
```
You can install a specific release version:
```bash
helm plugin install https://github.com/bery/helm-env.git --version <release version>
```

To use the plugin, you do not need any special dependencies. The installer will download the latest release with prebuilt binary from [GitHub releases](https://github.com/bery/helm-env/releases).

## Usage

### Single file usage
```bash
helm env -f <path to values file>
```

### Multiple files usage
```bash
helm env -f <path to values file> -f <path to values file> -f <path to values file>
```

### Directory usage
The plugin can also be used to recursively substitute environment variables in all the files in a specified directory.
```bash
helm env -f <path to directory>
```

### Mix files and directories
You can also decide to mix files and directories:
```bash
helm env -f <path to values file> -f <path to directory>
```

## Example
Sample helm values file:
```yaml
# values.yaml

image:
  repository: $REGISTRY/$IMAGE_NAME
  tag: $IMAGE_TAG
```
Environment variables configured in your environment (this should most likely be configured with your CI environment): 
```txt
REGISTRY => docker.com
IMAGE_NAME => helm-helm-env
IMAGE_TAG => test
```
Substitute Env:
```bash
helm env -f values.yaml
```
Result: 
```yaml
image:
  repository: docker.com/helm-helm-env
  tag: test
```
**Note:** If the value of the environment variable does not exist, it will be replaced with an empty string. For instance, from the above example, if `IMAGE_TAG` does not exist as an environment variable in the environment the result would have been: 

```yaml
image:
  repository: docker.com/helm-helm-env
  tag:
```

## Uninstall
```bash
helm plugin remove helm-env
```

## Testing locally
To test locally, run the command below to build and run the binary: 
> You need to have [Go](https://go.dev/) installed. Make sure to set `$GOPATH`
```bash
go build -o helm-env && ./helm-env -f </path/to/values/file>
```
## License

[MIT](LICENSE)
# helm-env
