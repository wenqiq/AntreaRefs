# 


The incorrect input of fields that have known schema should be rejected:

`cat acnp_demo.yaml`
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

`kubectl create -f acnp_demo.yaml`
error: error validating "acnp_demo.yaml": error validating data: ValidationError(ClusterNetworkPolicy.spec.appliedTo[0].namespaceSelector): unknown field "matchLables" in io.antrea.crd.v1alpha1.ClusterNetworkPolicy.spec.appliedTo.namespaceSelector; if you choose to ignore these errors, turn validation off with --validate=false

all of namespaceSelector/podSelector parameters verification in CRD schema will be improved.