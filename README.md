# provider-github

---

## Testing the Provider Locally

1. install minikube
2. `minikube start`
3. `kubectl create namespace crossplane-system`
4. `helm repo add crossplane-stable https://charts.crossplane.io/stable`
5. `helm repo update`
6. `helm install crossplane --namespace crossplane-system crossplane-stable/crossplane`
7. Using kubectl, apply a Secret containing credentials for the artifactory registry
    ```yaml
    apiVersion: v1
    type: kubernetes.io/dockerconfigjson
    data:
      .dockerconfigjson: {{ your secret here }}
    kind: Secret
    metadata:
      creationTimestamp: "2022-09-26T03:59:26Z"
      name: baasjfrogio
      namespace: crossplane-system
      resourceVersion: "7137"
      uid: df543815-f80d-4d7a-acd7-e019311bc1a3 
    ```
8. Using kubectl, apply a Secret containing credentials for GitHub, ensuring the credential has CRUD repo scopes
    ```yaml
    apiVersion: v1
    kind: Secret
    metadata:
      name: gh-creds
      namespace: crossplane-system
    type: Opaque
    stringData:
      credentials: {{ your github token }}
    ```
9. Using kubectl, apply the Provider to the cluster
    ```yaml
    apiVersion: pkg.crossplane.io/v1
    kind: Provider
    metadata:
      name: provider-gh
      namespace: crossplane-system
    spec:
      package: baas.jfrog.io/dkr-local-releases/platform/crossplane/provider-github:v0.0.0-79.g5af9b24
      packagePullPolicy: IfNotPresent
      revisionActivationPolicy: Automatic
      revisionHistoryLimit: 1
      packagePullSecrets:
        - name: baasjfrogio # {{ name of artifactory secret, in the same namespace}}
    ```
10. Using kubectl, apply the ProviderConfig to the cluster
    ```yaml
    apiVersion: github.crossplane.io/v1beta1
    kind: ProviderConfig
    metadata:
      name: provider-gh-config
      namespace: crossplane-system
    spec:
      credentials:
        source: Secret 
        secretRef:
          namespace: crossplane-system 
          name: gh-creds # {{ name of github creds secret, in the same namespace }} 
          key: credentials
    ```
11. Using kubectl, apply a test Claim to the cluster
    ```yaml
    apiVersion: repositories.github.crossplane.io/v1alpha1
    kind: Repository
    metadata:
     name: test-repo # {{ name of repo }} 
    spec:
      forProvider:
        owner: pgmitche # {{ name of repo owner }} 
        description: proving provider utility
      providerConfigRef:
        name: provider-gh-config # {{ name of ProviderConfig in same namespace }}
    ```

---

# Forks source README below

---

## Overview

`provider-github` is the Crossplane infrastructure provider for
[GitHub](https://github.com/). The provider that is built from the source code
in this repository can be installed into a Crossplane control plane and adds the
following new functionality:

* Custom Resource Definitions (CRDs) that model GitHub infrastructure and
  services
* Controllers to provision these resources in GitHub based on the users desired
  state captured in CRDs they create
* Implementations of Crossplane's [portable resource
  abstractions](https://crossplane.io/docs/master/concepts.html), enabling
  GitHub resources to fulfill a user's general need for cloud services

## Getting Started and Documentation

For getting started guides, installation, deployment, and administration, see
our [Documentation](https://crossplane.io/docs/latest).

## Contributing

provider-github is a community driven project and we welcome contributions. See
the Crossplane
[Contributing](https://github.com/crossplane/crossplane/blob/master/CONTRIBUTING.md)
guidelines to get started.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/12kmps/crossplane-provider-github/issues).

## Contact

Please use the following to reach members of the community:

* Slack: Join our [slack channel](https://slack.crossplane.io)
* Forums:
  [crossplane-dev](https://groups.google.com/forum/#!forum/crossplane-dev)
* Twitter: [@crossplane_io](https://twitter.com/crossplane_io)
* Email: [info@crossplane.io](mailto:info@crossplane.io)

## Roadmap

provider-github goals and milestones are currently tracked in the Crossplane
repository. More information can be found in
[ROADMAP.md](https://github.com/crossplane/crossplane/blob/master/ROADMAP.md).


## Governance and Owners

provider-github is run according to the same
[Governance](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md)
and [Ownership](https://github.com/crossplane/crossplane/blob/master/OWNERS.md)
structure as the core Crossplane project.

## Code of Conduct

provider-github adheres to the same [Code of
Conduct](https://github.com/crossplane/crossplane/blob/master/CODE_OF_CONDUCT.md)
as the core Crossplane project.

## Licensing

provider-github is under the Apache 2.0 license.

[![FOSSA
Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fcrossplane%2Fprovider-github.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fcrossplane%2Fprovider-github?ref=badge_large)