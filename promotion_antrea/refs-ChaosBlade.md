
# 官宣：恭喜ChaosBlade项目进入 CNCF Sandbox

阿里巴巴开源的混沌工程项目 ChaosBlade 通过 CNCF TOC 投票，顺利推进 CNCF Sandbox。CNCF 全称 Cloud Native Computing Foundation (云原生计算基金会) ，旨在为云原生软件构建可持续发展的生态系统，服务于厂商中立的快速增长的开源项目，如 Kubernetes、Prometheus、Envoy 等。

ChaosBlade github 地址：
https://github.com/chaosblade-io/chaosblade

# 项目介绍


ChaosBlade 是阿里巴巴 2019 年开源的混沌工程项目，包含混沌工程实验工具 chaosblade 和混沌工程平台 chaosblade-box，旨在通过混沌工程帮助企业解决云原生过程中高可用问题。实验工具 chaosblade 支持 3 大系统平台，4 种编程语言应用，共涉及 200 多个实验场景，3000 多个实验参数，可以精细化地控制实验范围。混沌工程平台 chaosblade-box 支持实验工具托管，除已托管 chaosblade 外，还支持 Litmuschaos 实验工具。已登记使用企业 40 多家，其中已在工商银行、中国移动、小米、京东等企业中落地使用。

# 核心能力


ChaosBlade 具备以下功能特点：

- 丰富的实验场景：包含基础资源（CPU、内存、网络、磁盘、进程、内核、文件等）、多语言应用服务（Java、C++、NodeJS、Golang 等）、Kubernetes 平台（覆盖 Container、Pod、Node 资源场景，包含上述实验场景）。

- 多样化的执行方式：除了使用平台白屏化操作，还可以通过工具自带的 blade 工具或者 kubectl、编码的方式执行。

- 便捷的场景扩展能力：所有的实验场景遵循混沌实验模型实现，并且不同层次场景对应不同的执行器，实现简单，易于扩展。

- 实验工具自动化部署：无需手动部署实验工具，实现实验工具在主机或集群上自动化部署。

- 支持开源实验工具托管：平台可托管业界主流的实验工具，如自身的 chaosblade 和外部的 litmuschaos 等。

- 统一混沌实验用户界面：用户无需关心不同工具的使用方式，在统一用户界面进行混沌实验。

- 多维度实验方式：支持从主机到 Kubernetes 资源，再到应用维度进行实验编排。

- 集成云原生生态：采用 Helm 部署管理，集成 Prometheus 监控，支持云原生实验工具托管等。

# 架构设计


Chaosblade-box 架构如下：


通过控制台页面可实现 chaosblade、litmuschaos 等已托管工具自动化部署，按照社区建立的混沌实验模型统一实验场景，根据主机、Kubernetes、应用来划分目标资源，通过目标管理器来控制，在实验创建页面，可以实现白屏化的目标资源选择。平台通过调用混沌实验执行来执行不同工具的实验场景，配合接入 prometheus 监控，可以观察实验 metric 指标，后续会提供丰富的实验报告。

Chaosblade-box 的部署也非常简单，具体可以查看：https://github.com/chaosblade-io/chaosblade-box/releases。

# 客户案例





# 未来规划


ChaosBlade 未来以云原生为基础，提供面向多集群、多环境、多语言的混沌工程平台和混沌工程实验工具。实验工具将继续聚焦在实验场景丰富度和稳定性方面，支持更多的 Kubernetes 资源场景和规范应用服务实验场景标准，提供多语言实验场景标准实现。混沌工程平台聚焦在简化混沌工程部署实施方面，后续会托管更多的混沌实验工具和兼容主流的平台，实现场景推荐，提供业务、系统监控集成，输出实验报告，在易用的基础上完成混沌工程操作闭环。

欢迎大家加入社区，共同推动混沌工程领域发展，切实在企业中落地，构建高可用的分布式系统。