

# antrea-controller logs

```bash
 tail -f deploy_antrea.log 

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
customresourcedefinition.apiextensions.k8s.io/networkpolicies.security.antrea.tanzu.vmware.com created
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
configmap/antrea-config-5ct9ktdt77 created
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
```


# delete

`kubectl delete -f antrea.yaml`

```bash
qiuwenqi@wenqiq-a01 Kubernetes % kubectl delete -f antrea.yaml 
customresourcedefinition.apiextensions.k8s.io "antreaagentinfos.clusterinformation.antrea.tanzu.vmware.com" deleted
customresourcedefinition.apiextensions.k8s.io "antreaagentinfos.crd.antrea.io" deleted
customresourcedefinition.apiextensions.k8s.io "antreacontrollerinfos.clusterinformation.antrea.tanzu.vmware.com" deleted
customresourcedefinition.apiextensions.k8s.io "antreacontrollerinfos.crd.antrea.io" deleted
customresourcedefinition.apiextensions.k8s.io "clustergroups.core.antrea.tanzu.vmware.com" deleted
customresourcedefinition.apiextensions.k8s.io "clustergroups.crd.antrea.io" deleted
customresourcedefinition.apiextensions.k8s.io "clusternetworkpolicies.crd.antrea.io" deleted
customresourcedefinition.apiextensions.k8s.io "clusternetworkpolicies.security.antrea.tanzu.vmware.com" deleted
customresourcedefinition.apiextensions.k8s.io "egresses.crd.antrea.io" deleted
customresourcedefinition.apiextensions.k8s.io "externalentities.core.antrea.tanzu.vmware.com" deleted
customresourcedefinition.apiextensions.k8s.io "externalentities.crd.antrea.io" deleted
customresourcedefinition.apiextensions.k8s.io "networkpolicies.crd.antrea.io" deleted
customresourcedefinition.apiextensions.k8s.io "networkpolicies.security.antrea.tanzu.vmware.com" deleted
customresourcedefinition.apiextensions.k8s.io "tiers.crd.antrea.io" deleted
customresourcedefinition.apiextensions.k8s.io "tiers.security.antrea.tanzu.vmware.com" deleted
customresourcedefinition.apiextensions.k8s.io "traceflows.crd.antrea.io" deleted
customresourcedefinition.apiextensions.k8s.io "traceflows.ops.antrea.tanzu.vmware.com" deleted
serviceaccount "antctl" deleted
serviceaccount "antrea-agent" deleted
serviceaccount "antrea-controller" deleted
clusterrole.rbac.authorization.k8s.io "aggregate-antrea-clustergroups-edit" deleted
clusterrole.rbac.authorization.k8s.io "aggregate-antrea-clustergroups-view" deleted
clusterrole.rbac.authorization.k8s.io "aggregate-antrea-policies-edit" deleted
clusterrole.rbac.authorization.k8s.io "aggregate-antrea-policies-view" deleted
clusterrole.rbac.authorization.k8s.io "aggregate-traceflows-edit" deleted
clusterrole.rbac.authorization.k8s.io "aggregate-traceflows-view" deleted
clusterrole.rbac.authorization.k8s.io "antctl" deleted
clusterrole.rbac.authorization.k8s.io "antrea-agent" deleted
clusterrole.rbac.authorization.k8s.io "antrea-cluster-identity-reader" deleted
clusterrole.rbac.authorization.k8s.io "antrea-controller" deleted
clusterrolebinding.rbac.authorization.k8s.io "antctl" deleted
clusterrolebinding.rbac.authorization.k8s.io "antrea-agent" deleted
clusterrolebinding.rbac.authorization.k8s.io "antrea-controller" deleted
configmap "antrea-config-5ct9ktdt77" deleted
service "antrea" deleted
deployment.apps "antrea-controller" deleted
apiservice.apiregistration.k8s.io "v1alpha1.stats.antrea.io" deleted
apiservice.apiregistration.k8s.io "v1alpha1.stats.antrea.tanzu.vmware.com" deleted
apiservice.apiregistration.k8s.io "v1beta1.controlplane.antrea.tanzu.vmware.com" deleted
apiservice.apiregistration.k8s.io "v1beta1.networking.antrea.tanzu.vmware.com" deleted
apiservice.apiregistration.k8s.io "v1beta1.system.antrea.io" deleted
apiservice.apiregistration.k8s.io "v1beta1.system.antrea.tanzu.vmware.com" deleted
apiservice.apiregistration.k8s.io "v1beta2.controlplane.antrea.io" deleted
apiservice.apiregistration.k8s.io "v1beta2.controlplane.antrea.tanzu.vmware.com" deleted
daemonset.apps "antrea-agent" deleted
mutatingwebhookconfiguration.admissionregistration.k8s.io "crdmutator.antrea.io" deleted
mutatingwebhookconfiguration.admissionregistration.k8s.io "crdmutator.antrea.tanzu.vmware.com" deleted
validatingwebhookconfiguration.admissionregistration.k8s.io "crdvalidator.antrea.io" deleted
validatingwebhookconfiguration.admissionregistration.k8s.io "crdvalidator.antrea.tanzu.vmware.com" deleted
```


delete log

```bash
E0424 19:42:09.280454       1 reflector.go:382] pkg/mod/github.com/tnqn/client-go@v0.18.4-1/tools/cache/reflector.go:125: Failed to watch *v1alpha2.ClusterGroup: unknown (get clustergroups.crd.antrea.io)
E0424 19:42:09.280468       1 reflector.go:382] pkg/mod/github.com/tnqn/client-go@v0.18.4-1/tools/cache/reflector.go:125: Failed to watch *v1alpha2.ClusterGroup: unknown (get clustergroups.core.antrea.tanzu.vmware.com)
E0424 19:42:09.292603       1 reflector.go:382] pkg/mod/github.com/tnqn/client-go@v0.18.4-1/tools/cache/reflector.go:125: Failed to watch *v1alpha1.ClusterNetworkPolicy: unknown (get clusternetworkpolicies.crd.antrea.io)
E0424 19:42:09.422146       1 reflector.go:382] pkg/mod/github.com/tnqn/client-go@v0.18.4-1/tools/cache/reflector.go:125: Failed to watch *v1alpha2.ExternalEntity: unknown (get externalentities.core.antrea.tanzu.vmware.com)
E0424 19:42:09.427140       1 reflector.go:382] pkg/mod/github.com/tnqn/client-go@v0.18.4-1/tools/cache/reflector.go:125: Failed to watch *v1alpha2.ExternalEntity: unknown (get externalentities.crd.antrea.io)
E0424 19:42:09.528201       1 reflector.go:382] pkg/mod/github.com/tnqn/client-go@v0.18.4-1/tools/cache/reflector.go:125: Failed to watch *v1alpha1.ClusterNetworkPolicy: unknown (get clusternetworkpolicies.security.antrea.tanzu.vmware.com)
E0424 19:42:09.556120       1 reflector.go:382] pkg/mod/github.com/tnqn/client-go@v0.18.4-1/tools/cache/reflector.go:125: Failed to watch *v1alpha1.Tier: unknown (get tiers.security.antrea.tanzu.vmware.com)
E0424 19:42:09.588288       1 reflector.go:382] pkg/mod/github.com/tnqn/client-go@v0.18.4-1/tools/cache/reflector.go:125: Failed to watch *v1alpha1.NetworkPolicy: unknown (get networkpolicies.crd.antrea.io)
E0424 19:42:09.597344       1 reflector.go:382] pkg/mod/github.com/tnqn/client-go@v0.18.4-1/tools/cache/reflector.go:125: Failed to watch *v1alpha1.NetworkPolicy: unknown (get networkpolicies.security.antrea.tanzu.vmware.com)
E0424 19:42:09.639272       1 reflector.go:382] pkg/mod/github.com/tnqn/client-go@v0.18.4-1/tools/cache/reflector.go:125: Failed to watch *v1alpha1.Tier: unknown (get tiers.crd.antrea.io)
E0424 19:42:09.685260       1 reflector.go:382] pkg/mod/github.com/tnqn/client-go@v0.18.4-1/tools/cache/reflector.go:125: Failed to watch *v1alpha1.Traceflow: unknown (get traceflows.crd.antrea.io)
E0424 19:42:09.717972       1 reflector.go:382] pkg/mod/github.com/tnqn/client-go@v0.18.4-1/tools/cache/reflector.go:125: Failed to watch *v1alpha1.Traceflow: unknown (get traceflows.ops.antrea.tanzu.vmware.com)
E0424 19:42:10.145848       1 reflector.go:178] pkg/mod/github.com/tnqn/client-go@v0.18.4-1/tools/cache/reflector.go:125: Failed to list *v1alpha1.ClusterNetworkPolicy: clusternetworkpolicies.crd.antrea.io is forbidden: User "system:serviceaccount:kube-system:antrea-controller" cannot list resource "clusternetworkpolicies" in API group "crd.antrea.io" at the cluster scope
E0424 19:42:10.472704       1 reflector.go:178] pkg/mod/github.com/tnqn/client-go@v0.18.4-1/tools/cache/reflector.go:125: Failed to list *v1alpha2.ClusterGroup: clustergroups.crd.antrea.io is forbidden: User "system:serviceaccount:kube-system:antrea-controller" cannot list resource "clustergroups" in API group "crd.antrea.io" at the cluster scope
E0424 19:42:10.490236       1 reflector.go:178] pkg/mod/github.com/tnqn/client-go@v0.18.4-1/tools/cache/reflector.go:125: Failed to list *v1alpha2.ExternalEntity: externalentities.core.antrea.tanzu.vmware.com is forbidden: User "system:serviceaccount:kube-system:antrea-controller" cannot list resource "externalentities" in API group "core.antrea.tanzu.vmware.com" at the cluster scope
E0424 19:42:10.490332       1 reflector.go:178] pkg/mod/github.com/tnqn/client-go@v0.18.4-1/tools/cache/reflector.go:125: Failed to list *v1alpha1.ClusterNetworkPolicy: clusternetworkpolicies.security.antrea.tanzu.vmware.com is forbidden: User "system:serviceaccount:kube-system:antrea-controller" cannot list resource "clusternetworkpolicies" in API group "security.antrea.tanzu.vmware.com" at the cluster scope
I0424 19:42:10.600078       1 controller.go:342] Stopping Antrea controller
rpc error: code = NotFound desc = an error occurred when try to find container "7d93e3ffbe22a732cb92c45ced2f45c570992320194b860feff58e885ca9d6f4": not found% 
```
