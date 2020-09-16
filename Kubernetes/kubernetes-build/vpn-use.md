# vpn client on ubuntu

- 安装pptp客户端及配置
```
apt install pptp-linux
# 创建连接(永久), 新增后的配置写入配置文件/etc/ppp/chap-secrets中
pptpsetup --create vpn --server **(your vpn ip) --username **(name) --password **(password) --encrypt --start
# 将默认路由指向pptp连接, 即可开始使用vpn
route add default ppp0
```
- 使用
```bash
# 连线, vpn是创建的vpn名字
pon vpn
route add default ppp0
# 断开连接
poff vpn
```
