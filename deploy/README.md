# Deploy scripts

## Usage

```
# Deploy sripts should be used from the git root of this repo
# Deploy StackRox locally on Kubernetes
$ ./deploy/k8s/deploy-local.sh

# Deploy StackRox locally on OpenShift
$ ./deploy/openshift/deploy-local.sh

# Deploy StackRox on a remote OpenShift cluster with an exposed route
$ LOAD_BALANCER=route ./deploy/openshift/deploy.sh
```

## Env variables

Most environment variables can be found in [common/env.sh](https://github.com/stackrox/rox/blob/e57c8fe3b98c3833f2b2ff0d634fc325b88e0372/deploy/common/env.sh).

| **Name**           | **Values**            | **Description**                                                                                                                                                            |
|--------------------|-----------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `COLLECTION_METHOD`  | `ebpf`  \| `kernel-module` | Set the collection method for collector.                                                                                                                                   |
| `HOTRELOAD`          | `true`  \| `false`         | `HOTRELOAD` mounts Sensor and Central local binaries into locally running pods. Only works with docker-desktop.  Alternatively you can use ./dev-tools/enabled-hotreload.sh. |
| `LOAD_BALANCER`      | `route` \| `lb`            | Configure how to expose Central, important if deployed on remote clusters. Use `route` for OpenShift, `lb` for Kubernetes.                                                 |
| `MAIN_IMAGE_TAG`     | `string`                   | Configure the image tag of the `stackrox/main` image to be deployed.                                                                                                       |
| `MONITORING_SUPPORT` | `true`  \| `false`         | Enable StackRox monitoring.                                                                                                                                                |
| `STORAGE`            | `none`  \| `pvc`           | Defines which storage to use for the Central database, to preserve data between Central restarts it is recommended to use `pvc`.                                                |