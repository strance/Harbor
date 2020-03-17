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
- -I: 在指定位置插入规则，默认在头部，如果指定编号，如 -I INPUT 3，则新增规则的编号为3.
- -s: 指定源地址,多个地址之间用“逗号”隔开，不能包含空格。还可以指定网段，如10.6.0.0/46。!取反，如 ! -s 192.168.1.146。
- -d: 指定目标地址,多个地址之间用“逗号”隔开，不能包含空格。还可以指定网段，如10.6.0.0/46。!取反，如 ! -d 192.168.1.146。
- -j: 指定匹配条件满足时执行的动作即target
- -A: 在尾部追加规则
- -R：修改规则，除了指定规则编号，必须指定规则对应的原有匹配条件
- -P: 指定要修改的链
- -p: 指定协议（protocol）,如tcp,udp,icmp,udplite,icmpv6,esp,ah,sctp,mh,默认匹配所有协议
- -i: 指定报文流入网卡，只适用于PREROUTING,INPUT,FORWARD链
- -o: 指定报文流出网卡，只适用于FORWARD,OUTPUT,POSTROUTING链

### 扩展匹配条件
- tcp模块
  - --dport: 指定报文的目标端口。需要使用对应的扩展模块：-m tcp.当使用了-p选项指定了报文的协议时，如果没有使用-m指定对应的扩展模块，使用了扩展匹配条件，iptables默认会调用与-p选项对应的协议名称相同的模块。如果扩展匹配田间所依赖的扩展模块名与-p对应的协议名称不同，则不能省略-m选项。可以指定端口范围，如--dport 22:25, --dport :22, --dport 80:.
    ```shell
    # 拒绝来自192.168.1.146 22端口的报文，-m tcp可以省略，因为都是扩展模块名和协议名相同
    $ iptables -t filter -I INPUT -s 192.168.1.146 -p tcp -m tcp --dport 22 -j REJECT
    ```
   - --sport: 指定报文的源端口。使用同--dport.
- multiport模块
  - --sport或者--dport只能指定连续的端口范围，要指定多个离散的端口，需要使用multiport模块。--dports，--sprots，多个端口之间用逗号隔开，不能有空格.不能省略-m选项，只适用于tcp和udp协议。
- iprange模块
  - --src-range,--dst-range: 指定一段连续的IP地址范围，如 `iptables -t filter -I INPUT -p tcp -m iprange --src-range 192.168.1.127-192.168.1.146 -j DROP`
- string
  - --algo: 用于指定匹配算法，bm与Kmp二选一,如
  ```shell
  # 拒绝包含"2345"字符串的报文
  $ iptables -t filter -I INPUT -m string --algo bm --string "2345" -j REJECT
  ```
  - --string: 用于指定需要匹配的字符串。
- time模块
  - --timestart: 指定开始时间
  - --timestop: 指定结束时间
  ```shell
  # 拒绝每天早上9点到下午6点80端口的报文
  $ iptables -t filter -I INPUT -p tcp --dport 80 -m time --timestart 09:00:00 --timestop 18:00:00 -j REJECT
  ```
  - --weekdays: 指定星期几。可选值：1，2，3，4，5，6，7或者Mon,Tue,Wed,Thu,Fri,Sat,Sun.多个值之间用逗号隔开，不能有空格.可以!.
  - --monthdays: 指定每月的哪天。可以!.
  - --datestart: 指定具体哪天开始，如2020-03-02
  - --datestop: 指定具体哪天结束，如2020-03-02
- connlimit模块：限制每个IP地址同时链接到Server端的链接数量。
  - --connlimit-above: 指定每个IP链接的数量上限。
  ```shell
  # 限制每个客户端IP的ssh链接数量最多2个
  $ iptables -t filter -I INPUT -p tcp --dport 22 -m -connlimit --connlimit-above 2 -j REJECT 
  ```
  - --connlimit-mask: 限制"某类网段"的链接数量
  
### 示例
- 查看INPUT表的规则
  ```shell
  $ iptables --line -nvL INPUT
  ```
- 在第二条规则前增加一条规则，拒绝192.168.1.146上的所有报文访问
  ```shell
  $ iptables -t filter -I INPUT 2 -s 192.168.1.146 -j DROP
  ```
### 保存规则
- 规则默认保存在/etc/sysconfig/iptables文件中。
  ```shell
  # 查看规则
  $ cat /etc/sysconfig/iptables
  # 保存规则
  $ service iptables save
  # 如果误操作之后没有保存规则，可以重启iptables到上次保存/etc/sysconfig/iptables文件时的模样
  $ service iptables restart
  ```

