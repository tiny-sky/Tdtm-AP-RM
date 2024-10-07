## TDTM-AP/RM

此仓库为构建 TDTM 的 AP、RM 模块

提供三个 RM 服务模块
- order
- account
- stock

如果我们进行一个类似于购买商品的业务时，可能会涉及到用户的扣费、商品的出库、订单的生产，我们需要保证这些操作的原子性：要么全部成功，要求全部失败

## 运行

#### 启动 RM 模块
```sh
go run RM/main.go
```

#### 启动 AP 模块

AP 提供直连、服务发现两种方式

直连模式
```sh
go run AP/direct/direct.go
```
服务发现模式
```sh
go run AP/discovery/discovery.go
```
