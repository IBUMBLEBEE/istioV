# kubernetes-TrafficDash

**Programming language：**

* Golang: 1.10.1      解包工具和前端计算
* Golang: 1.10.1      流程控制中心
* rust: 1.24          基于前端的计算
* webassembly: 1.0.0  基于前端的计算

Thinking：

1. 来源于Dapper，现在与kubernetes POD为基本单位，在每个POD内共享网络内部署抓包工具，数据包带有特殊标签往外传输，数据包在解包端立即解包，通过前端计算，形成前端需要的流量数据，通过前端计算，形成网络流量转发图。
