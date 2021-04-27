

# Create Kind Kubernetes Cluster

`kind create cluster --name "antrea-cluster" --config kind_cluster.yaml`


```bash
qiuwenqi@wenqiq-a01 Kubernetes % kind create cluster --name "antrea-cluster" --config kind_cluster.yaml
Creating cluster "antrea-cluster" ...
 ✓ Ensuring node image (kindest/node:v1.20.2) 🖼 
 ✓ Preparing nodes 📦 📦 📦  
 ✓ Writing configuration 📜 
 ✓ Starting control-plane 🕹️ 
 ✓ Installing StorageClass 💾 
 ✓ Joining worker nodes 🚜 
Set kubectl context to "kind-antrea-cluster"
You can now use your cluster with:

kubectl cluster-info --context kind-antrea-cluster

Thanks for using kind! 😊
```

`cat kind_cluster.yaml`

```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  disableDefaultCNI: true
  podSubnet: 10.10.0.0/16
nodes:
- role: control-plane
- role: worker
- role: worker
```


get nodes

`kubectl get nodes`

```bash
qiuwenqi@wenqiq-a01 Kubernetes % kubectl get nodes
NAME                           STATUS     ROLES                  AGE   VERSION
antrea-cluster-control-plane   NotReady   control-plane,master   97s   v1.20.2
antrea-cluster-worker          NotReady   <none>                 65s   v1.20.2
antrea-cluster-worker2         NotReady   <none>                 64s   v1.20.2
```

get pods

`kubectl get pod -A`

```bash
qiuwenqi@wenqiq-a01 Kubernetes % kubectl get pod -A
NAMESPACE            NAME                                                   READY   STATUS    RESTARTS   AGE
kube-system          coredns-74ff55c5b-92g77                                0/1     Pending   0          3m8s
kube-system          coredns-74ff55c5b-w46zp                                0/1     Pending   0          3m8s
kube-system          etcd-antrea-cluster-control-plane                      1/1     Running   0          3m17s
kube-system          kube-apiserver-antrea-cluster-control-plane            1/1     Running   0          3m17s
kube-system          kube-controller-manager-antrea-cluster-control-plane   1/1     Running   0          3m17s
kube-system          kube-proxy-4gwfw                                       1/1     Running   0          2m52s
kube-system          kube-proxy-4s9w7                                       1/1     Running   0          3m9s
kube-system          kube-proxy-9wxmh                                       1/1     Running   0          2m51s
kube-system          kube-scheduler-antrea-cluster-control-plane            1/1     Running   0          3m17s
local-path-storage   local-path-provisioner-78776bfc44-kqxpq                0/1     Pending   0          3m8s
```


# Deploy antrea

These instructions assume that you have built the Antrea Docker image locally
(e.g. by running `make` from the root of the repository).

```bash
# "fix" the host's veth interfaces (for the different Kind Nodes)
kind get nodes | xargs ./hack/kind-fix-networking.sh

# load the Antrea Docker image in the Nodes
kind load docker-image projects.registry.vmware.com/antrea/antrea-ubuntu:latest

# deploy Antrea
./hack/generate-manifest.sh --kind | kubectl apply -f -
```

## kind-fix-networking.sh

kind get nodes | xargs ./hack/kind-fix-networking.sh

```bash
qiuwenqi@wenqiq-a01 antrea % kind get nodes --name antrea-cluster | xargs ./hack/kind-fix-networking.sh
Unable to find image 'antrea/ethtool:latest' locally
latest: Pulling from antrea/ethtool
da7391352a9b: Pulling fs layer
14428a6d4bcd: Pulling fs layer
2c2d948710f2: Pulling fs layer
8585700a019e: Pulling fs layer
2c2d948710f2: Download complete
14428a6d4bcd: Verifying Checksum
14428a6d4bcd: Download complete
8585700a019e: Verifying Checksum
8585700a019e: Download complete
da7391352a9b: Download complete
da7391352a9b: Pull complete
14428a6d4bcd: Pull complete
2c2d948710f2: Pull complete
8585700a019e: Pull complete
Digest: sha256:6e50086cbe45bea21c343c760f9ec97797fed3f6b8b6d43ff0890284e2a11c54
Status: Downloaded newer image for antrea/ethtool:latest
Disabling TX checksum offload for node antrea-cluster-control-plane (veth56b0e1f)
Actual changes:
tx-checksumming: off
	tx-checksum-ip-generic: off
	tx-checksum-sctp: off
tcp-segmentation-offload: off
	tx-tcp-segmentation: off [requested on]
	tx-tcp-ecn-segmentation: off [requested on]
	tx-tcp-mangleid-segmentation: off [requested on]
	tx-tcp6-segmentation: off [requested on]
net.ipv4.tcp_retries2 = 4
Disabling TX checksum offload for node antrea-cluster-worker (veth05d1c2f)
Actual changes:
tx-checksumming: off
	tx-checksum-ip-generic: off
	tx-checksum-sctp: off
tcp-segmentation-offload: off
	tx-tcp-segmentation: off [requested on]
	tx-tcp-ecn-segmentation: off [requested on]
	tx-tcp-mangleid-segmentation: off [requested on]
	tx-tcp6-segmentation: off [requested on]
net.ipv4.tcp_retries2 = 4
Disabling TX checksum offload for node antrea-cluster-worker2 (veth919ac00)
Actual changes:
tx-checksumming: off
	tx-checksum-ip-generic: off
	tx-checksum-sctp: off
tcp-segmentation-offload: off
	tx-tcp-segmentation: off [requested on]
	tx-tcp-ecn-segmentation: off [requested on]
	tx-tcp-mangleid-segmentation: off [requested on]
	tx-tcp6-segmentation: off [requested on]
net.ipv4.tcp_retries2 = 4
```

## load images

```bash
qiuwenqi@wenqiq-a01 antrea % kind load docker-image projects.registry.vmware.com/antrea/antrea-ubuntu:latest --name antrea-cluster
Image: "projects.registry.vmware.com/antrea/antrea-ubuntu:latest" with ID "sha256:9052bc068f38bd7a116765c858646b5352892917e64c63bcf5db55ad56ba9f31" not yet present on node "antrea-cluster-control-plane", loading...
Image: "projects.registry.vmware.com/antrea/antrea-ubuntu:latest" with ID "sha256:9052bc068f38bd7a116765c858646b5352892917e64c63bcf5db55ad56ba9f31" not yet present on node "antrea-cluster-worker", loading...
Image: "projects.registry.vmware.com/antrea/antrea-ubuntu:latest" with ID "sha256:9052bc068f38bd7a116765c858646b5352892917e64c63bcf5db55ad56ba9f31" not yet present on node "antrea-cluster-worker2", loading...
```

## create Antrea

```bash
qiuwenqi@wenqiq-a01 antrea % ./hack/generate-manifest.sh --kind | kubectl apply -f -
customresourcedefinition.apiextensions.k8s.io/antreaagentinfos.clusterinformation.antrea.tanzu.vmware.com created
customresourcedefinition.apiextensions.k8s.io/antreaagentinfos.crd.antrea.io created
customresourcedefinition.apiextensions.k8s.io/antreacontrollerinfos.clusterinformation.antrea.tanzu.vmware.com created
customresourcedefinition.apiextensions.k8s.io/antreacontrollerinfos.crd.antrea.io created
customresourcedefinition.apiextensions.k8s.io/clustergroups.core.antrea.tanzu.vmware.com created
customresourcedefinition.apiextensions.k8s.io/clustergroups.crd.antrea.io created
customresourcedefinition.apiextensions.k8s.io/clusternetworkpolicies.crd.antrea.io created
customresourcedefinition.apiextensions.k8s.io/clusternetworkpolicies.security.antrea.tanzu.vmware.com created
customresourcedefinition.apiextensions.k8s.io/egresses.crd.antrea.io created
customresourcedefinition.apiextensions.k8s.io/externalentities.core.antrea.tanzu.vmware.com created
customresourcedefinition.apiextensions.k8s.io/externalentities.crd.antrea.io created
customresourcedefinition.apiextensions.k8s.io/networkpolicies.crd.antrea.io created
customresourcedefinition.apiextensions.k8s.io/tiers.crd.antrea.io created
customresourcedefinition.apiextensions.k8s.io/tiers.security.antrea.tanzu.vmware.com created
customresourcedefinition.apiextensions.k8s.io/traceflows.crd.antrea.io created
customresourcedefinition.apiextensions.k8s.io/traceflows.ops.antrea.tanzu.vmware.com created
serviceaccount/antctl created
serviceaccount/antrea-agent created
serviceaccount/antrea-controller created
clusterrole.rbac.authorization.k8s.io/aggregate-antrea-clustergroups-edit created
clusterrole.rbac.authorization.k8s.io/aggregate-antrea-clustergroups-view created
clusterrole.rbac.authorization.k8s.io/aggregate-antrea-policies-edit created
clusterrole.rbac.authorization.k8s.io/aggregate-antrea-policies-view created
clusterrole.rbac.authorization.k8s.io/aggregate-traceflows-edit created
clusterrole.rbac.authorization.k8s.io/aggregate-traceflows-view created
clusterrole.rbac.authorization.k8s.io/antctl created
clusterrole.rbac.authorization.k8s.io/antrea-agent created
clusterrole.rbac.authorization.k8s.io/antrea-cluster-identity-reader created
clusterrole.rbac.authorization.k8s.io/antrea-controller created
clusterrolebinding.rbac.authorization.k8s.io/antctl created
clusterrolebinding.rbac.authorization.k8s.io/antrea-agent created
clusterrolebinding.rbac.authorization.k8s.io/antrea-controller created
configmap/antrea-config-bm47htbkt5 created
service/antrea created
deployment.apps/antrea-controller created
apiservice.apiregistration.k8s.io/v1alpha1.stats.antrea.io created
apiservice.apiregistration.k8s.io/v1alpha1.stats.antrea.tanzu.vmware.com created
apiservice.apiregistration.k8s.io/v1beta1.controlplane.antrea.tanzu.vmware.com created
apiservice.apiregistration.k8s.io/v1beta1.networking.antrea.tanzu.vmware.com created
apiservice.apiregistration.k8s.io/v1beta1.system.antrea.io created
apiservice.apiregistration.k8s.io/v1beta1.system.antrea.tanzu.vmware.com created
apiservice.apiregistration.k8s.io/v1beta2.controlplane.antrea.io created
apiservice.apiregistration.k8s.io/v1beta2.controlplane.antrea.tanzu.vmware.com created
daemonset.apps/antrea-agent created
mutatingwebhookconfiguration.admissionregistration.k8s.io/crdmutator.antrea.io created
mutatingwebhookconfiguration.admissionregistration.k8s.io/crdmutator.antrea.tanzu.vmware.com created
validatingwebhookconfiguration.admissionregistration.k8s.io/crdvalidator.antrea.io created
validatingwebhookconfiguration.admissionregistration.k8s.io/crdvalidator.antrea.tanzu.vmware.com created
Error from server (BadRequest): error when creating "STDIN": CustomResourceDefinition in version "v1" cannot be handled as a CustomResourceDefinition: v1.CustomResourceDefinition.Spec: v1.CustomResourceDefinitionSpec.Versions: []v1.CustomResourceDefinitionVersion: v1.CustomResourceDefinitionVersion.Schema: v1.CustomResourceValidation.OpenAPIV3Schema: v1.JSONSchemaProps.Properties: v1.JSONSchemaProps.Properties: v1.JSONSchemaProps.v1.JSONSchemaProps.Type: Items: unmarshalerDecoder: json: cannot unmarshal string into Go struct field JSONSchemaProps.properties.items of type v1.JSONSchemaProps, error found in #10 byte of ...|:"object"},"type":"a|..., bigger context ...|":"array"}},"required":["action"],"type":"object"},"type":"array"},"priority":{"format":"float","max|...
```


## get pods

```bash
qiuwenqi@wenqiq-a01 antrea % kubectl get pods -A
NAMESPACE            NAME                                                   READY   STATUS    RESTARTS   AGE
kube-system          antrea-agent-2q7xr                                     2/2     Running   0          51s
kube-system          antrea-agent-q8vvx                                     2/2     Running   0          51s
kube-system          antrea-agent-wcr9j                                     2/2     Running   0          51s
kube-system          antrea-controller-86594745f5-6trgt                     1/1     Running   0          51s
kube-system          coredns-74ff55c5b-92g77                                1/1     Running   0          22m
kube-system          coredns-74ff55c5b-w46zp                                1/1     Running   0          22m
kube-system          etcd-antrea-cluster-control-plane                      1/1     Running   0          22m
kube-system          kube-apiserver-antrea-cluster-control-plane            1/1     Running   0          22m
kube-system          kube-controller-manager-antrea-cluster-control-plane   1/1     Running   0          22m
kube-system          kube-proxy-4gwfw                                       1/1     Running   0          22m
kube-system          kube-proxy-4s9w7                                       1/1     Running   0          22m
kube-system          kube-proxy-9wxmh                                       1/1     Running   0          22m
kube-system          kube-scheduler-antrea-cluster-control-plane            1/1     Running   0          22m
local-path-storage   local-path-provisioner-78776bfc44-kqxpq                1/1     Running   0          22m
```


# kubectl create -f cnp_demo.yaml

`qiuwenqi@wenqiq-a01 Kubernetes % cat cnp_demo.yaml`


```yaml
apiVersion: crd.antrea.io/v1alpha1
kind: ClusterNetworkPolicy
metadata:
  name: blacklist
spec:
    priority: 5
    tier: securityops
    appliedTo:
      - namespaceSelector:
          matchLables:
            env: dummy
    ingress:
      - action: Drop
```

# kubectl get clusternetworkpolicies.crd.antrea.io

`qiuwenqi@wenqiq-a01 antrea % kubectl get clusternetworkpolicies.crd.antrea.io`

```bash
NAME        TIER          PRIORITY   DESIRED NODES   CURRENT NODES   AGE
blacklist   securityops   5          1               1               16s
```

# kubectl get clusternetworkpolicies

```bash
qiuwenqi@wenqiq-a01 antrea % kubectl get clusternetworkpolicies              
NAME        TIER          PRIORITY   DESIRED NODES   CURRENT NODES   AGE
blacklist   securityops   5          1               1               34s
```

# kubectl describe clusternetworkpolicies

`qiuwenqi@wenqiq-a01 antrea % kubectl describe clusternetworkpolicies blacklist
`

```bash
Name:         blacklist
Namespace:    
Labels:       <none>
Annotations:  <none>
API Version:  crd.antrea.io/v1alpha1
Kind:         ClusterNetworkPolicy
Metadata:
  Creation Timestamp:  2021-04-23T17:57:37Z
  Generation:          1
  Managed Fields:
    API Version:  crd.antrea.io/v1alpha1
    Fields Type:  FieldsV1
    fieldsV1:
      f:spec:
        f:appliedTo:
        f:ingress:
      f:status:
        .:
        f:currentNodesRealized:
        f:desiredNodesRealized:
        f:observedGeneration:
        f:phase:
    Manager:      antrea-controller
    Operation:    Update
    Time:         2021-04-23T17:57:37Z
    API Version:  crd.antrea.io/v1alpha1
    Fields Type:  FieldsV1
    fieldsV1:
      f:spec:
        .:
        f:priority:
        f:tier:
    Manager:         kubectl-create
    Operation:       Update
    Time:            2021-04-23T17:57:37Z
  Resource Version:  168293
  UID:               60a9022d-01b7-4a0a-aee0-22956c7faa4a
Spec:
  Applied To:
    Namespace Selector:
      Match Lables:
        Env:  dummy
  Ingress:
    Action:  Drop
    Name:    ingress-drop-946c95a
  Priority:  5
  Tier:      securityops
Status:
  Current Nodes Realized:  1
  Desired Nodes Realized:  1
  Observed Generation:     1
  Phase:                   Realized
Events:                    <none>
```

# kubectl get clusternetworkpolicies json

`kubectl get clusternetworkpolicies blacklist -o json`

```json
{
    "apiVersion": "crd.antrea.io/v1alpha1",
    "kind": "ClusterNetworkPolicy",
    "metadata": {
        "creationTimestamp": "2021-04-23T17:57:37Z",
        "generation": 1,
        "managedFields": [
            {
                "apiVersion": "crd.antrea.io/v1alpha1",
                "fieldsType": "FieldsV1",
                "fieldsV1": {
                    "f:spec": {
                        "f:appliedTo": {},
                        "f:ingress": {}
                    },
                    "f:status": {
                        ".": {},
                        "f:currentNodesRealized": {},
                        "f:desiredNodesRealized": {},
                        "f:observedGeneration": {},
                        "f:phase": {}
                    }
                },
                "manager": "antrea-controller",
                "operation": "Update",
                "time": "2021-04-23T17:57:37Z"
            },
            {
                "apiVersion": "crd.antrea.io/v1alpha1",
                "fieldsType": "FieldsV1",
                "fieldsV1": {
                    "f:spec": {
                        ".": {},
                        "f:priority": {},
                        "f:tier": {}
                    }
                },
                "manager": "kubectl-create",
                "operation": "Update",
                "time": "2021-04-23T17:57:37Z"
            }
        ],
        "name": "blacklist",
        "resourceVersion": "168293",
        "uid": "60a9022d-01b7-4a0a-aee0-22956c7faa4a"
    },
    "spec": {
        "appliedTo": [
            {
                "namespaceSelector": {
                    "matchLables": {
                        "env": "dummy"
                    }
                }
            }
        ],
        "ingress": [
            {
                "action": "Drop",
                "name": "ingress-drop-946c95a"
            }
        ],
        "priority": 5,
        "tier": "securityops"
    },
    "status": {
        "currentNodesRealized": 1,
        "desiredNodesRealized": 1,
        "observedGeneration": 1,
        "phase": "Realized"
    }
}
```


# kubectl get clusternetworkpolicies

`qiuwenqi@wenqiq-a01 antrea % kubectl get clusternetworkpolicies blacklist -o yaml`

```yaml
apiVersion: crd.antrea.io/v1alpha1
kind: ClusterNetworkPolicy
metadata:
  creationTimestamp: "2021-04-23T17:57:37Z"
  generation: 1
  managedFields:
  - apiVersion: crd.antrea.io/v1alpha1
    fieldsType: FieldsV1
    fieldsV1:
      f:spec:
        f:appliedTo: {}
        f:ingress: {}
      f:status:
        .: {}
        f:currentNodesRealized: {}
        f:desiredNodesRealized: {}
        f:observedGeneration: {}
        f:phase: {}
    manager: antrea-controller
    operation: Update
    time: "2021-04-23T17:57:37Z"
  - apiVersion: crd.antrea.io/v1alpha1
    fieldsType: FieldsV1
    fieldsV1:
      f:spec:
        .: {}
        f:priority: {}
        f:tier: {}
    manager: kubectl-create
    operation: Update
    time: "2021-04-23T17:57:37Z"
  name: blacklist
  resourceVersion: "168293"
  uid: 60a9022d-01b7-4a0a-aee0-22956c7faa4a
spec:
  appliedTo:
  - namespaceSelector:
      matchLables:
        env: dummy
  ingress:
  - action: Drop
    name: ingress-drop-946c95a
  priority: 5
  tier: securityops
status:
  currentNodesRealized: 1
  desiredNodesRealized: 1
  observedGeneration: 1
  phase: Realized
```


# Development

projects.registry.vmware.com/antrea/antrea-ubuntu:latest



