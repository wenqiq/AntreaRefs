# 使用Andrea实现Kubernetes集群的pod间网络通信并增强网络策略(NetworkPolicy)

Antrea是一个基于Kubernetes原生的项目，它实现了容器网络接口（CNI）和Kubernetes NetworkPolicy，
从而为Pod工作负载提供网络连接和安全性。 
Antrea将可编程网络的优势从Open vSwitch（OVS）扩展到Kubernetes。

## 统一网络堆栈

Antrea通过在Open vSwitch之上构建的统一网络堆栈，简化了跨不同云和操作系统的Kubernetes网络。

## 易于操作

Antrea的设计采用Kubernetes控制器模式，并为网络开发者或运维提供相关一些常用、高效的工具，从而简化部署、日常操作和故障定位。

## 灵活性和可扩展性

借助Open vSwitch的可编程性，Antrea可以扩展为支持高级网络特性，例如内核旁路和网络服务网格。

# 特征

## 创建覆盖网络

使用VXLAN或Geneve进行IP覆盖网络来实现Kubernetes pod间网络通信。 
（可选）使用IPSec数据包加密来加密节点到节点的通信。

## 增强Kubernetes网络策略

Antrea enforces Kubernetes Network Policy API which assigns network traffic filtering rules to pods. Extending firewall functionality to the pod edge enables nano-segmentation of workloads; service graphs are strictly enforced, policies follow workload rescheduling and scaling events, and if pods become compromised, egress restrictions limit further propagation.

Antrea实现了Kubernetes网络策略API，该API将网络流量过滤规则作用到Pod。
将防火墙功能扩展到Pod边缘，可以实现工作负载的纳米细分；
严格执行服务图，策略遵循工作负载重新安排和扩展事件，并且如果pod受到威胁，则出口流量限制会限制进一步的传播。

## 改善网络性能

OVS的性能优于iptables，尤其是随着规则数量的增加。 
OVS社区中进行了许多努力，比如通过诸如Intel DPDK，AF_XDP套接字，硬件卸载等技术来加速数据包IO和数据包处理。

