# douyin-project



## 抖音项目服务端简单实现

实现接口参考：https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof#

### 使用说明

分为微服务和上层服务两块，项目执行方式（下面步骤按顺序执行）

1. 创建数据库，sql文件在/sql文件夹中（之后记得修改每一个微服务中的mysql连接地址）
2. 启动etcd（如果端口有修改，修改每一个微服务/上层服务中的配置文件）
3. 微服务启动，微服务在/microservice，启动方法是，在每一个微服务的./rpc文件夹中执行：`go run main.go`
4. 上层服务启动，每一个服务在/service中，启动方法是，进入每一个上层服务的中间夹后，执行`go run main.go`