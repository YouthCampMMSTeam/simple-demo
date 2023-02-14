

#### 微服务部分 文件结构

在microservice文件夹下，按照微服务分类。

单个微服务具体结构以relation为例子：

relation

- model: 数据库对象和对应操作

- api（可选）：如果微服务对应的接口需要对外暴露，则该文件夹用于提供对外http实现

- rpc：该微服务的rpc接口实现

  - 文件夹

    - client：rpc客户端代码，方便后续使用以及测试
    - config：配置文件设置，本次使用yaml文件作为配置文件
      - relation.yaml：本服务相关的参数设置

    - idl：接口idl文件，本次使用kitex实现rpc，需要使用idl文件定义接口和数据类型

    - impl：接口实际实现

    - kitex_gen：kitex代码生成的部分

    - svcctx：即serviceContext，用于存储本微服务相关的服务上下文信息，比如数据库连接、缓存连接、与其他微服务的rpc连接等等

  - 文件
    - handler.go ：提供rpc服务对象的封装
    - main.go : main函数，串联整个微服务