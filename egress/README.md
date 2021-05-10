# Egress Policy


## Egress

`Egress` enables a CRD API for Antrea that supports specifying which egress
(SNAT) IP the traffic from the selected Pods to the external network should use.
When a selected Pod accesses the external network, the egress traffic will be
tunneled to the Node that hosts the egress IP if it's different from the Node
that the Pod runs on and will be SNATed to the egress IP when leaving that Node.
Usage example:

```yaml
apiVersion: crd.antrea.io/v1alpha2
kind: Egress
metadata:
  name: egress-web
spec:
  appliedTo:
    podSelector:
      matchLabels:
        role: web
    namespaceSelector:
      matchLabels:
        env: prod
  egressIP: 10.0.10.8
```

The `appliedTo` field specifies the grouping criteria of Pods to which the
Egress applies to. Pods can be selected cluster-wide using `podSelector`. If set
with a `namespaceSelector`, Pods from Namespaces selected by the
namespaceSelector will be selected. Empty `appliedTo` selects nothing. The field
is mandatory.

The `egressIP` field specifies the egress (SNAT) IP the traffic from the
selected Pods to the external network should use. **The IP must be assigned to
an arbitrary interface of one Node, and one Node only. It must be reachable from
all Nodes.** For IPv4 cluster, it must be an IPv4 address; for IPv6 cluster, it
must be an IPv6 address. The field is mandatory.

**Note**: If more than one Egress applies to a Pod and they specify different
`egressIP`, the effective egress IP will be selected randomly.

In the above example, the Egress applies to Pods which match the labels
"role=web" from Namespaces which match the labels "env=prod". The source IPs of
their egress traffic to external network will be translated to 10.0.10.8.

### Requirements for this Feature

This feature is currently only supported for Nodes running Linux and "encap"
mode. The support for Windows and other traffic modes will be added in the
future.

# issues

https://github.com/vmware-tanzu/antrea/issues/667

## Add Egress CRD types

https://github.com/vmware-tanzu/antrea/pull/1433

## Add iptables interface for implementing Egress

https://github.com/vmware-tanzu/antrea/pull/1998

## Support automatic failover for Egress

https://github.com/vmware-tanzu/antrea/issues/2128

