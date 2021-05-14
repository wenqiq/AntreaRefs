# Enable pod networking and enforce network policies for Kubernetes clusters

Antrea is a Kubernetes-native project that implements the Container Network Interface (CNI) and Kubernetes NetworkPolicy thereby providing network connectivity and security for pod workloads. Antrea extends the benefit of programmable networks from Open vSwitch (OVS) to Kubernetes.

## Unified Network Stack

Antrea simplifies Kubernetes networking across differing clouds and operating systems with a unified network stack built atop Open vSwitch.

## Ease of Operations

Antrea is designed to ease deployment, operations and troubleshooting by adopting Kubernetes controller patterns and providing diagnostics consumable by tools network operators know and love.

## Flexibility & Extensibility

With the programmability of Open vSwitch under the hood, Antrea can be extended to support advanced network use cases like kernel bypass and network service mesh.

# Features

## Create overlay networks

Enable Kubernetes pod networking with IP overlay networks using VXLAN or Geneve for encapsulation. Optionally encrypt node-to-node communication using IPSec packet encryption.

## Enforce Network Polices

Antrea enforces Kubernetes Network Policy API which assigns network traffic filtering rules to pods. Extending firewall functionality to the pod edge enables nano-segmentation of workloads; service graphs are strictly enforced, policies follow workload rescheduling and scaling events, and if pods become compromised, egress restrictions limit further propagation.

## Improved Network Performance

OVS performs better than iptables, especially as the number of rules increases. There are numerous efforts in the OVS community to speed up packet IO and packet processing through technologies like Intel DPDK, AF_XDP sockets, hardware offloading, etc.


# refs

https://antrea.io