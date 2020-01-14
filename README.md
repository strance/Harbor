"# Harbor" 

"# iptables"
- 报文流向
  - 到本机某进程的报文： PREROUTING --> INPUT
  - 由本机发出的报文： PREROUTING --> FORWARD --> POSTROUTING
  - 由本机的某进程发出的报文： OUTPUT --> POSTROUTING
- 规则表
  - filter表：负责过滤功能，防火墙；内核模块：iptables_filter
  - nat表：network address translation, 网络地址转换功能；内核模块：iptable_nat
  - mangle表： 拆解报文，做出修改，并重新封装的功能；内核模块：iptable_mangle
  - raw表：关闭nat表上启用的连接追踪机制：内核模块：iptable_raw
- netfilters链包含的表及优先级（从左到右）
  - PREROUTING: raw, mangle, nat表
  - INPUT: mangle, （centOS7中还有nat表），filter表
  - FORWARD: mangle, filter表
  - OUTPUT: raw, mangle, nat, filter表
  - POSTROUTING: mangle, nat表
