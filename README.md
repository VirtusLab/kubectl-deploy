# kubectl-deploy

[kubectl](https://kubernetes.io/docs/reference/kubectl/kubectl/) + [render](https://github.com/VirtusLab/render) + [crypt](https://github.com/VirtusLab/crypt) family

**kubectl-deploy** it's a really simple `kubectl` plugin which renders Kubernetes manifests (go-template) and applies them.

The way it works is similar to tiller-less [Helm 3](https://github.com/helm/community/blob/master/helm-v3/000-helm-v3.md), but:
- follows the standard [go-template](https://golang.org/pkg/text/template/) API 
- adds [custom render functions](https://github.com/VirtusLab/render/blob/master/README.md#notable-standard-and-sprig-functions)
- provides Secret Management via [crypt](https://github.com/VirtusLab/crypt)
- uses native Kubernetes authorization mechanism (everything is done via `kubectl`)

## Installation

Place **kubectl-deploy** in your `PATH`:

```bash
curl -#L \
    --url "https://raw.githubusercontent.com/VirtusLab/kubectl-deploy/master/kubectl-deploy" \
    --output "/usr/local/bin/kubectl-deploy"
chmod +x "/usr/local/bin/kubectl-deploy"
```

More info at [Extend kubectl with plugins](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/).

## Usage
  
Ensure **kubectl-deploy** plugin is recognized by `kubectl`:

```console
kubectl plugin list
```

Example usage:
 
```console
Usage:
  kubectl deploy [flags]

Flags:
      --config string      config.yaml (required)
  -c, --context string     k8s context (optional)
  -d, --debug              enable debug logging level output (optional)
      --dry-run            If true, only print the object that would be sent, without sending it
  -f, --file string        k8s manifest or directory
  -h, --help               help for kubectl
  -n, --namespace string   k8s namespace (optional)
  -t, --trace              enable trace logging level output (optional)

```

For more advanced templates and rendering please take a look at [render#usage](https://github.com/VirtusLab/render/blob/master/README.md#notable-standard-and-sprig-functions).

## Examples

### Kubernetes manifests

```console
kubectl deploy --file examples/manifests --config examples/config.yaml
```

### Helm Charts

Currently rendering and applying the Helm chart is not supported. 
Mostly because Helm does not follow the standard [go-template](https://golang.org/pkg/text/template/) API.
 
As a workaround please use `helm template` command first. 
 
## Similar projects

- [Helm Plugin "No Magic"](https://github.com/giantswarm/helm-nomagic)
  
## Contribution

Feel free to file issues or pull requests.