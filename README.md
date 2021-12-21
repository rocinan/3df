# go-clean-frame

## Changelog
- **v1**: test

## Description
This project has  4 Domain layer :
 * Domain Layer
 * Delivery Layer
 * Service Layer
 * Usecase Layer  
 * Repository Layer

 - domain   领域层 modal 和 interface
 - delivery 运输层 http/ws/rpc... 以及接口数据预处理
 - service  服务层 需要一直运行的服务，比如定时任务，代理服务等等
 - usecase  用例层 纯业务逻辑 不引入约束
 - repo     存储层 cache db ...
 
#### The diagram:

![golang clean architecture](https://github.com/bxcodec/go-clean-arch/raw/master/clean-arch.png)


### How To Run This Project
Since the project already use Go Module, I recommend to put the source code in any folder but GOPATH.