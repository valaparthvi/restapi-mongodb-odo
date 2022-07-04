# Percona Server for MongoDB

This chart deploys a minimal Percona Server for MongoDB Cluster on Kubernetes controlled by Percona Operator for MongoDB.

Useful links:
- [Operator Github repository](https://github.com/percona/percona-server-mongodb-operator)
- [Operator Documentation](https://www.percona.com/doc/kubernetes-operator-for-psmongodb/index.html)

## Pre-requisites
* Percona Operator for MongoDB running in your Kubernetes cluster. See installation details [here](https://github.com/percona/percona-helm-charts/blob/main/charts/psmdb-operator) or in the [Operator Documentation](https://www.percona.com/doc/kubernetes-operator-for-psmongodb/helm.html).
* Kubernetes 1.19+
* Helm v3

# Chart Details
This chart will deploy a minimal Percona Server for MongoDB Cluster in Kubernetes. It will create a Custom Resource, and the Operator will trigger the creation of corresponding Kubernetes primitives: StatefulSets, Pods, Secrets, etc.

## Installing the Chart
To install the chart with the `minimal-cluster` release name using a dedicated namespace (recommended):

```sh
helm install my-db ./psmdb-db --namespace my-namespace
```

The templates included in this chart are created from the [cr-minimal.yaml](https://raw.githubusercontent.com/percona/percona-server-mongodb-operator/v1.12.0/deploy/cr-minimal.yaml) provided by Percona for mongodb service, and its percona/psmdb-db chart.