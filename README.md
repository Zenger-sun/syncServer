# syncServer

一个基于tcp的C/S游戏同步服务器，用于测试游戏的同步需求。

## Version  
v0.1  
> 基本能用。目前能想到的后续计划：1、集群化；2、加入LockStep同步；3、多种传输协议尝试；4、优化Context功能，如日志模块和配置模块  

## Quick start  

### Requirments
Go version>=1.16

### Installation
```go

git clone https://github.com/Zenger-sun/syncServer.git  

// 根据需要修改消息 
pb/proto
sh gen.sh
```

### Build & Run

```go
// main()请参考 example/syncServer.go
go build -o server.exe
```

## Documentation

sync包含的模块如下:  
![sync](https://user-images.githubusercontent.com/22719311/158414263-fd1ddd8e-f0b2-434e-b5de-7d4f1ffe2e07.png)
listen采用tcp连接，自定义head用于解包，未加密，后续会考虑多种传输协议    
考虑到玩家操作都是顺序的，采用actor封装transport，由transport顺序校验、分发  
抽象出serviceM用于管理syncServer下的所有service，serviceM还可以管理路由，为后续的cluster提供接口，除此之外，serviceM还充当service的监控器，如果有service panic，还会重启service
