# Deploy antrea


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



