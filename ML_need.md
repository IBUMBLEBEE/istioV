# iptables日志

## iptables日志了解

1. iptables 拥有强大的log（数据包记录）功能，可以把按特定规则匹配的ip数据包以log的形式传递到用户层供用户分析，这样我们就可以非常方便的了解内核中都有哪些ip数据包在传递以及这些数据包的内容。

2. iptables有三种log记录形式，分别是log、ulog、nflog。
    1. log用于将匹配的数据包记录到系统的syslog中去，用户也可以直接通过dmesg命令查看。log命令只记录包头的一些。
    2. ulog通过netlink套接字将数据包多播到指定netlink多播组，这样任何感兴趣的进程都可以通过建立netlink套接字来接受内核中的数据包信息。ulog可以将整个数据包拷贝并发送给应用程序，当然也可以指定发送方数据包的字节数。
    3. nflog不仅可以接受来自iptables的数据，还可以向iptables通过netlink套接字发送控制数据。