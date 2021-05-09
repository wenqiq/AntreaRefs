
# 恭喜Antrea项目进入 CNCF Sandbox

太平洋时间5月6日，经过云原生计算基金会（CNCF）技术监督委员会（TOC）投票通过，
VMware开源的项目Antrea正式成为沙箱级项目（Sandbox Level Project）。

Antrea项目是一个开源的Kubernetes CNI网络插件解决方案，基于Open vSwitch（OVS），
旨在提供Kubernetes原生的、高效、安全、跨平台的CNI和NetworkPolicy。

Antrea github 地址：
https://github.com/vmware-tanzu/antrea

# 项目介绍

![Antrea Overview](img/antrea_overview.svg.png)


Antrea最初于2019年11月发布，为Kubernetes集群实现了容器网络接口（CNI）和Kubernetes NetworkPolicy，
Antrea主要在Kubernetes集群3/4层提供网络服务和一些安全特性，它建立在Open vSwitch（OVS）的基础之上，
OVS是为分布式多层交换性能而优化的开源技术。
由于OVS中具有高度可扩展的体系结构，Antrea项目能够更加高效，例如，Open vSwitch使得Antrea以更高效的方式实现NetworkPolicy。

Antrea项目利用Open vSwitch作为网络数据平面，Open vSwitch是一种高性能可编程虚拟交换机，同时支持Linux和Windows。
Antrea基于Open vSwitch，能够以高性能和高效的方式实施Kubernetes网络策略，并支持Linux和Windows等多平台。
同时，借助Open vSwitch的“可编程”特性， Antrea能够在Open vSwitch的基础上实现广泛的网络、安全功能和服务。


# 核心能力


Antrea 具备以下功能特点：

- Kubernetes原生：     
  Antrea遵循最佳实践来扩展Kubernetes API，为用户提供熟悉的概念，同时在其自己的实现中尽可能复用用Kubernetes库。

- 基于Open vSwitch:     
  Antrea基于Open vSwitch来实现所有网络功能，包括Kubernetes Service负载平衡，
  并启用硬件卸载以支持最苛刻的工作负载。

- 兼容性:     
  支持在私有云，公共云和裸金属上运行Antrea，您可以根据自己的基础设施和实际情况选择适当的流量模式（有或没有overlay网络）。

- 支持Windows节点:    
  得益于Open vSwitch的可移植性，
  Antrea可以在Linux和Windows Kubernetes节点上使用相同的数据平面实现其功能。

- 支持复杂的网络策略模型:    
  Antrea提供了全面的网络策略模型， 它以Kubernetes网络策略为基础，
  同时添加新功能，例如策略分层，规则优先级和集群级策略。

- 故障排查和监控工具:    
  Antrea带有CLI和UI工具，可提供可视化和故障诊断能力（数据包跟踪，策略分析，流量检查）。
  可以展示Prometheus指标并支持导出网络流信息，该信息可以在Kibana仪表板中可视化。

- 数据加密:    
  使用覆盖Pod网络时，使用IPsec隧道对节点间Pod通信进行加密。

- 一键部署:    
  通过应用单个YAML清单文件来部署Antrea


# 架构设计


Antrea 架构如下：

![antrea architecture](img/arch.svg.png)

Antrea项目利用Open vSwitch作为网络数据平面，Open vSwitch是一种高性能可编程虚拟交换机，同时支持Linux和Windows。
Antrea基于Open vSwitch，能够以高性能和高效的方式实施Kubernetes网络策略，并支持Linux和Windows等多平台。
同时，借助Open vSwitch的“可编程”特性， Antrea能够在Open vSwitch的基础上实现广泛的网络、安全功能和服务。

以上架构是Linux平台的Antrea Agent运行情况，Windows平台的架构可以参考项目文档的Windows部分。

Antrea 的部署非常简单，部署Kubernetes集群之后，通过一键执行
`kubectl apply -f https://raw.githubusercontent.com/vmware-tanzu/antrea/main/build/yamls/antrea.yml`
即可安装Andrea的最新版本，也可以指定版本安装：
`kubectl apply -f https://github.com/vmware-tanzu/antrea/releases/download/<TAG>/antrea.yml`,
具体可以查看：https://github.com/vmware-tanzu/antrea/blob/main/docs/getting-started.md。

# 客户案例





# 未来规划

Antrea未来的规划包括：更好的支持Windows节点、网络策略增强、网络诊断和监测、更加灵活的Flexible IPAM、
网络流量出口策略、NFV和电信设备支持、Kubernetes节点安全、7层网络安全策略和可视化、NetworkPolicy规模和性能测试、
将OVS与DPDK或AF_XDP结合使用以实现更高性能等。

通过最新发布的 1.0 版本，你可以测试最新的特性，并设想下一个版本的新可能性。

欢迎大家加入社区，积极贡献代码，多多实践和测试，共同推动Kubernetes容器网络和网络策略发展，
讨论更多Antrea新特性的规划， 构建未来更高效的Kubernetes CNI和Network Policy。


# CNCF

CNCF (Cloud Native Computing Foundation)成立于2015年12月，隶属于Linux Foundation，是非营利性组织。

CNCF（云原生计算基金会）致力于培育和维护一个厂商中立的开源生态系统，来推广云原生技术。
通过将最前沿的模式民主化，让这些创新为大众所用。


# refs

https://www.cncf.io/sandbox-projects/

https://www.cncf.io/online-programs/securing-and-accelerating-the-kubernetes-cni-data-plane-with-project-antrea-and-nvidia-mellanox-connectx-smartnics/

https://github.com/vmware-tanzu/antrea