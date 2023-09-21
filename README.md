# vyos-tunnel-ipip

1. 业务需求，隧道对端是动态IP，做了域名解析

2. sourceip 和 remote 都可用域名

3. 使用方法

```
Usage of ./vyos:
  -remote string
        remote
  -sourceip string
        sourceip (default "10.88.88.53")
  -tun string
        tun (default "tun0")
  -tunip string
        tunip (default "198.0.0.1/30")
```

3. remote  域名
4. 1 分钟检测一次，域名IP是否有变化
