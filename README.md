### Harbor 
- 架构
- 安装
  - 解压harbor安装包，然后执行./prepare进行初始化，/data/sercet文件夹下会生成认证需要的所需文件
  - /data/sercet文件夹下的文件拷贝到其他节点，确保多实例公用同一套认证文件
  - 对于https认证，需要将生成的证书拷贝到docker,即/etc/docker/certs.d/，拷贝完成后重启docker
  - 数据库初始化时，需要执行初始化脚本，脚本在harbor-core的容器内/harbor/migrations/postgresql
  
  
### iptables
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
  - PREROUTING: raw --> mangle --> nat表
  - INPUT: mangle --> （centOS7中还有nat表）--> filter表
  - FORWARD: mangle --> filter表
  - OUTPUT: raw --> mangle --> nat --> filter表
  - POSTROUTING: mangle --> nat表

### Target 处理方式
- ACCEPT: 允许数据包通过
- DROP：直接丢弃数据包，不给任何回应消息，客户端过了超时时间才会有反应。
- REJETCT: 拒绝数据包通过，必要时会给数据发送端一个响应的消息，客户端刚请求就会收到拒绝的消息。
- SNAT: 源地址转换，解决内网用户用同一个公网地址上网的问题。
- MASQUERADE: SNAT的一种特殊形式，适用于动态的、临时会变的IP上。
- DNAT: 目标地址转换。
- REDIRECT: 在本机做端口映射。
- LOG: 在/var/log/messages文件中记录日志信息，然后数据包传递给下一条规则，也就是说除了记录以外不对数据包做任何其他操作，仍然让下一条规则去匹配。

### iptables 选项
- -t: 指定要操作的表，缺省项 -t filter
- -L: 列出规则
- -v: 显示更详细的信息
- -n: 不对IP地址进行名称解析（比如不加-n选项 0.0.0.0/0 就显示未 anywhere）
- --line: 显示规则的编号
- -F: 清空规则

- 增加一条规则，拒绝192.168.1.146上的所有报文访问
  ```shell
  $ iptables -t filter -I INPUT -s 192.168.1.146 -j DROP
  ```

