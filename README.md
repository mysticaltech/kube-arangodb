# ArangoDB Kubernetes Operator

[![Docker Pulls](https://img.shields.io/docker/pulls/arangodb/kube-arangodb.svg)](https://hub.docker.com/r/arangodb/kube-arangodb/)

ArangoDB Kubernetes Operator helps to run ArangoDB deployments
on Kubernetes clusters.

To get started, follow the Installation instructions below and/or
read the [tutorial](https://www.arangodb.com/docs/stable/tutorials-kubernetes.html).

## State

The ArangoDB Kubernetes Operator is still in **development**.

Running ArangoDB deployments (single, active-failover or cluster)
is reasonably stable, and we're in the process of validating
production readiness of various Kubernetes platforms.

The feature set of the ArangoDB Kubernetes Operator is close to what
it is intended to be.

[Documentation](./docs/README.md)

### Production readiness state

Beginning with Version 0.3.11 we maintain a production readiness
state for individual new features, since we expect that new
features will first be released with an "alpha" or "beta" readiness
state and over time move to full "production readiness".

Operator will supports versions supported on providers and maintained by Kubernetes.
Once version is not supported anymore it will go into "Deprecating" state and will be marked as deprecated on Minor release.

Kubernetes versions starting from 1.16 are supported and tested, charts and manifests can use API Versions which are not present in older versions.

The following table has the general readiness state, the table below
covers individual newer features separately.

| Platform            | Kubernetes Version | ArangoDB Version | State      | Remarks               | Provider Remarks                   |
|---------------------|--------------------|------------------|------------|-----------------------|------------------------------------|
| Google GKE          | 1.17               | >= 3.6.0         | Production | Don't use micro nodes |                                    |
| Google GKE          | 1.18               | >= 3.6.0         | Production | Don't use micro nodes |                                    |
| Google GKE          | 1.19               | >= 3.6.0         | Production | Don't use micro nodes |                                    |
| Google GKE          | 1.20               | >= 3.6.0         | Production | Don't use micro nodes |                                    |
| Azure AKS           | 1.18               | >= 3.6.0         | Production |                       |                                    |
| Azure AKS           | 1.19               | >= 3.6.0         | Production |                       |                                    |
| Azure AKS           | 1.20               | >= 3.6.0         | Production |                       |                                    |
| Amazon EKS          | 1.16               | >= 3.6.0         | Production |                       | [Amazon EKS](./docs/providers/eks) |
| Amazon EKS          | 1.17               | >= 3.6.0         | Production |                       | [Amazon EKS](./docs/providers/eks) |
| Amazon EKS          | 1.18               | >= 3.6.0         | Production |                       | [Amazon EKS](./docs/providers/eks) |
| Amazon EKS          | 1.19               | >= 3.6.0         | Production |                       | [Amazon EKS](./docs/providers/eks) |
| Amazon EKS          | 1.20               | >= 3.6.0         | Production |                       | [Amazon EKS](./docs/providers/eks) |
| IBM Cloud           | 1.17               | >= 3.6.0         | Deprecated |                       |                                    |
| IBM Cloud           | 1.18               | >= 3.6.0         | Production |                       |                                    |
| IBM Cloud           | 1.19               | >= 3.6.0         | Production |                       |                                    |
| IBM Cloud           | 1.20               | >= 3.6.0         | Production |                       |                                    |
| OpenShift           | 3.11               | >= 3.6.0         | Production |                       |                                    |
| OpenShift           | 4.2                | >= 3.6.0         | Production |                       |                                    |
| BareMetal (kubeadm) | 1.16               | >= 3.6.0         | Production |                       |                                    |
| BareMetal (kubeadm) | 1.17               | >= 3.6.0         | Production |                       |                                    |
| BareMetal (kubeadm) | 1.18               | >= 3.6.0         | Production |                       |                                    |
| BareMetal (kubeadm) | 1.19               | >= 3.6.0         | Production |                       |                                    |
| BareMetal (kubeadm) | 1.20               | >= 3.6.0         | Production |                       |                                    |
| BareMetal (kubeadm) | 1.21               | >= 3.6.0         | Production |                       |                                    |
| Minikube            | 1.14+              | >= 3.6.0         | Devel Only |                       |                                    |
| Other               | 1.14+              | >= 3.6.0         | Devel Only |                       |                                    |

Feature-wise production readiness table:

| Feature                                 | Operator Version | ArangoDB Version | ArangoDB Edition      | State        | Enabled | Flag                                       | Remarks                                                                  |
|-----------------------------------------|------------------|------------------|-----------------------|--------------|---------|--------------------------------------------|--------------------------------------------------------------------------|
| Pod Disruption Budgets                  | 0.3.10           | Any              | Community, Enterprise | Alpha        | True    | N/A                                        | N/A                                                                      |
| Pod Disruption Budgets                  | 0.3.11           | Any              | Community, Enterprise | Production   | True    | N/A                                        | N/A                                                                      |
| Volume Resizing                         | 0.3.10           | Any              | Community, Enterprise | Alpha        | True    | N/A                                        | N/A                                                                      |
| Volume Resizing                         | 0.3.11           | Any              | Community, Enterprise | Production   | True    | N/A                                        | N/A                                                                      |
| Disabling of liveness probes            | 0.3.10           | Any              | Community, Enterprise | Alpha        | True    | N/A                                        | N/A                                                                      |
| Disabling of liveness probes            | 0.3.11           | Any              | Community, Enterprise | Production   | True    | N/A                                        | N/A                                                                      |
| Volume Claim Templates                  | 0.3.11           | Any              | Community, Enterprise | Alpha        | True    | N/A                                        | N/A                                                                      |
| Volume Claim Templates                  | 1.0.0            | Any              | Community, Enterprise | Production   | True    | N/A                                        | N/A                                                                      |
| Prometheus Metrics Exporter             | 0.3.11           | Any              | Community, Enterprise | Alpha        | True    | N/A                                        | Prometheus required                                                      |
| Prometheus Metrics Exporter             | 1.0.0            | Any              | Community, Enterprise | Production   | True    | N/A                                        | Prometheus required                                                      |
| Sidecar Containers                      | 0.3.11           | Any              | Community, Enterprise | Alpha        | True    | N/A                                        | N/A                                                                      |
| Sidecar Containers                      | 1.0.0            | Any              | Community, Enterprise | Production   | True    | N/A                                        | N/A                                                                      |
| Operator Single Mode                    | 1.0.4            | Any              | Community, Enterprise | Production   | False   | --mode.single                              | Only 1 instance of Operator allowed in namespace when feature is enabled |
| TLS SNI Support                         | 1.0.3            | >= 3.7.0         | Enterprise            | Production   | True    | --deployment.feature.tls-sni               | N/A                                                                      |
| TLS Runtime Rotation Support            | 1.0.4            | > 3.7.0          | Enterprise            | Alpha        | False   | --deployment.feature.tls-rotation          | N/A                                                                      |
| TLS Runtime Rotation Support            | 1.1.0            | > 3.7.0          | Enterprise            | Production   | True    | --deployment.feature.tls-rotation          | N/A                                                                      |
| JWT Rotation Support                    | 1.0.4            | > 3.7.0          | Enterprise            | Alpha        | False   | --deployment.feature.jwt-rotation          | N/A                                                                      |
| JWT Rotation Support                    | 1.1.0            | > 3.7.0          | Enterprise            | Production   | True    | --deployment.feature.jwt-rotation          | N/A                                                                      |
| Encryption Key Rotation Support         | 1.0.4            | > 3.7.0          | Enterprise            | Alpha        | False   | --deployment.feature.encryption-rotation   | N/A                                                                      |
| Encryption Key Rotation Support         | 1.1.0            | > 3.7.0          | Enterprise            | Production   | True    | --deployment.feature.encryption-rotation   | N/A                                                                      |
| Encryption Key Rotation Support         | 1.2.0            | > 3.7.0          | Enterprise            | NotSupported | False   | --deployment.feature.encryption-rotation   | N/A                                                                      |
| Version Check                           | 1.1.4            | >= 3.7.0         | Community, Enterprise | Alpha        | False   | --deployment.feature.upgrade-version-check | N/A                                                                      |
| Operator Maintenance Management Support | 1.0.7            | >= 3.7.0         | Community, Enterprise | Alpha        | False   | --deployment.feature.maintenance           | N/A                                                                      |
| Operator Maintenance Management Support | 1.2.0            | >= 3.7.0         | Community, Enterprise | Production   | True    | --deployment.feature.maintenance           | N/A                                                                      |
| Operator Internal Metrics Exporter      | 1.1.9            | >= 3.7.0         | Community, Enterprise | Alpha        | False   | --deployment.feature.metrics-exporter      | N/A                                                                      |
| Operator Internal Metrics Exporter      | 1.2.0            | >= 3.7.0         | Community, Enterprise | Production   | True    | --deployment.feature.metrics-exporter      | N/A                                                                      |
| Operator Internal Metrics Exporter      | 1.2.3            | >= 3.7.0         | Community, Enterprise | Production   | True    | --deployment.feature.metrics-exporter      | It is always enabled                                                     |
| Operator Ephemeral Volumes              | 1.2.2            | >= 3.7.0         | Community, Enterprise | Alpha        | False   | --deployment.feature.ephemeral-volumes     | N/A                                                                      |

## Release notes for 0.3.16

In this release we have reworked the Helm charts. One notable change is
that we now create a new service account specifically for the operator.
The actual deployment still runs by default under the `default` service
account unless one changes that. Note that the service account under
which the ArangoDB runs needs a small set of extra permissions. For
the `default` service account we grant them when the operator is
deployed. If you use another service account you have to grant these
permissions yourself. See
[here](docs/Manual/Deployment/Kubernetes/DeploymentResource.md#specgroupserviceaccountname-string)
for details.

## Installation of latest release using Kubectl

```bash
kubectl apply -f https://raw.githubusercontent.com/arangodb/kube-arangodb/1.2.6/manifests/arango-crd.yaml
kubectl apply -f https://raw.githubusercontent.com/arangodb/kube-arangodb/1.2.6/manifests/arango-deployment.yaml
# To use `ArangoLocalStorage`, also run
kubectl apply -f https://raw.githubusercontent.com/arangodb/kube-arangodb/1.2.6/manifests/arango-storage.yaml
# To use `ArangoDeploymentReplication`, also run
kubectl apply -f https://raw.githubusercontent.com/arangodb/kube-arangodb/1.2.6/manifests/arango-deployment-replication.yaml
```

This procedure can also be used for upgrades and will not harm any
running ArangoDB deployments.

## Installation of latest release using kustomize

Installation using [kustomize](https://kubernetes.io/docs/tasks/manage-kubernetes-objects/kustomization/) looks like installation from yaml files,
but user is allowed to modify namespace or resource names without yaml modifications.

IT is recommended to use kustomization instead of handcrafting namespace in yaml files - kustomization will replace not only resource namespaces,
but also namespace references in resources like ClusterRoleBinding.

Example kustomization file:
```
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

namespace: my-custom-namespace

bases:
  - https://github.com/arangodb/kube-arangodb/manifests/kustomize/deployment/?ref=1.0.3
```

## Installation of latest release using Helm

Only use this procedure for a new install of the operator. See below for
upgrades.

```bash
# The following will add the helm repository to your local cache
helm repo add arangodb https://arangodb.github.io/kube-arangodb
# The following will install the custom resources required by the operators.
helm install kube-arangodb-crd arangodb/kube-arangodb-crd
# The following will install the operator for `ArangoDeployment` &
# `ArangoDeploymentReplication` resources.
helm install kube-arangodb arangodb/kube-arangodb
# To use `ArangoLocalStorage`, set field `operator.features.storage` to true
helm install kube-arangodb arangodb/kube-arangodb --set "operator.features.storage=true"
```

## Upgrading the operator using Helm

To upgrade the operator to the latest version with Helm, just run:

```bash
helm upgrade kube-arangodb arangodb/kube-arangodb --atomic --reuse-values

## Building

```bash
DOCKERNAMESPACE=<your dockerhub account> make
kubectl apply -f manifests/arango-deployment-dev.yaml
# To use `ArangoLocalStorage`, also run
kubectl apply -f manifests/arango-storage-dev.yaml
# To use `ArangoDeploymentReplication`, also run
kubectl apply -f manifests/arango-deployment-replication-dev.yaml
```

## ArangoExporter

[ArangoExporter](https://github.com/arangodb-helper/arangodb-exporter) project has been merged with ArangoOperator.
Starting from ArangoDB 3.6 Servers expose metrics endpoint with prometheus compatible format. From this point Exporter
is used only for TLS and/or Authentication termination to be compatible with all Prometheus installations.

ArangoExporter documentation can be found [here](./docs/design/exporter.md)
