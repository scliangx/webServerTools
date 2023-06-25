# webServerTools

### 基于go语言gin框架常用的组件搭建而成的web脚手架,开箱即用，专注于业务开发


#### 目录详细信息
 ```text
    ├─apps                       程序入口
    ├─common                     常用的项目公共代码
    │  ├─error                   项目全局错误信息
    │  ├─global                  项目全局参数
    │  ├─logger                  项目日志
    │  └─response                项目统一响应
    ├─config                     项目配置
    ├─deploy                     项目部署文件
    ├─http_server                业务处理
    │  ├─controllers             控制层
    │  ├─manager                 实际处理业务层
    │  ├─models                  数据库模型
    │  └─routes                  服务路由
    ├─internal                   项目服务核心组件
    │  ├─db                      数据库组件
    │  ├─elasticsearch           es组件相关实现
    │  ├─grpc                    grpc相关案例
    │  ├─kafka                   kafka消息队列组件
    │  ├─mongodb                 mongodb组件
    │  ├─redis                   redis组件
    │  └─sessions                session实现
    ├─middlewares                中间件
    ├─proto                      proto文件存放目录
    └─utils                      常用工具函数

  request -> apps -> router -> controller -> manager -> controller -> response
 ```
### run app:
```shell
# 编译proto
protoc -I . --go_out=plugins=grpc:. ./*.proto

# 运行程序
go run apps/main.go web-server --config ./config.yaml
```

### 有任何问题欢迎联系
![coder小新]([https://github.com/coderitx/webServerTools/assets/54300717/a6308135-9f01-4b77-8213-06e2f4979c6f](https://wx-1317712916.cos.ap-beijing.myqcloud.com/%E6%89%AB%E7%A0%81_%E6%90%9C%E7%B4%A2%E8%81%94%E5%90%88%E4%BC%A0%E6%92%AD%E6%A0%B7%E5%BC%8F-%E6%A0%87%E5%87%86%E8%89%B2%E7%89%88.png))


<h6>-------------文中有许多借鉴于别人的blog进行集成的的地方，不足之处多指教-------------------</h6>
<h6>-------------文中有许多借鉴于别人的blog进行集成的的地方，不足之处多指教-------------------</h6>
<h6>-------------文中有许多借鉴于别人的blog进行集成的的地方，不足之处多指教-------------------</h6>
