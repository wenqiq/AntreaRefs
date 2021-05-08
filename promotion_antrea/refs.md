
# 为什么选Antrea

- 开源CNI插件

- 都采用成熟稳定的技术, 功能&性能俱佳

- 都有可选的商业支持

- 重点, 都支持不做Overlay封装的路由模式. 这样就可以免去多层Overlay带来的系统开销以及管理运维复杂性的问题了.

- 重点, 都支持No-SNAT


Antrea 不做Overlay封装和No-SANT在其文档中说明确说明使用方法和设计缘由, 
有兴趣的小伙伴可以访问下方链接:
https://github.com/vmware-tanzu/antrea/blob/main/docs/noencap-hybrid-modes.md

具体方法就是在部署Antrea的YMAL之前, 先编辑好配置参数:

```yaml
antrea-agent.conf: |
    ... ...
    trafficEncapMode: noEncap

    noSNAT: true # Set to true to disable Antrea SNAT for external traffic
    ... ...
```


# release


on 13 Mar
v0.13.1
1a26a8c  
on 27 Feb
v0.12.2
908adfd  
on 26 Feb
v0.11.3
6efba31  
on 12 Feb
v0.13.0
90b1cc7  
on 11 Feb
v0.12.1
2645b10  
on 11 Feb
v0.11.2
b654333  
on 23 Dec 2020
v0.12.0
2bb88bb  
on 21 Nov 2020
v0.11.1
67d9317  
on 14 Nov 2020
v0.11.0
50d0270  
on 12 Nov 2020
v0.10.2
0c2b182  

on 1 Oct 2020
v0.10.1
c0235ad  
on 25 Sep 2020
v0.10.0
1e2076d  
on 4 Sep 2020
v0.9.3
29e999a  
on 28 Aug 2020
v0.9.2
e482e01  
on 22 Aug 2020
v0.9.1
68fa221  
on 14 Aug 2020
v0.9.0
0622c0a  
on 14 Jul 2020
v0.8.2
e9ff6c8  
on 10 Jul 2020
v0.8.1
1c7bea9  
on 3 Jul 2020
v0.8.0
826894e  
on 16 Jun 2020
v0.7.2
46f84e3  

## v0.7.1

on 6 Jun 2020

c912180  

## v0.7.0

on 30 May 2020

2a46474  

## v0.7.0-rc.1

on 30 May 2020

2a46474 

## v0.6.0

on 30 Apr 2020

a90aa24  

## v0.5.1

on 2 Apr 2020
 
a5cf823

## v0.5.0

on 26 Mar 2020
 
02dc8b3  

## v0.5.0-rc.1

on 26 Mar 2020
 
02dc8b3

## v0.4.1

on 28 Feb 2020
 
80d86d4

## v0.4.0

on 21 Feb 2020
 
bdb0890

## v0.3.0

on 24 Jan 2020
 
5e71f87  

## v0.2.0

on 20 Dec 2019

1256f13 

## v0.1.1

on 28 Nov 2019

2f2f26b  

## v0.1.0

on 19 Nov 2019

02c648b

## v0.1.0-rc.2

on 19 Nov 2019

02c648b 

## v0.1.0-rc.1

on 19 Nov 2019

7410a1b  