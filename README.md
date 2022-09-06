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
    ├─middlewares                中间件
    ├─internal                   项目服务核心组件
    │  ├─cache                   缓存组件
    │  ├─db                      数据库组件
    │  └─kafka                   kafka消息队列组件
    ├─static                     静态资源
    └─utils                      常用工具函数
   
  request -> apps -> router -> controller -> manager -> controller -> response
 ```
### run app:
```text
go run apps/main.go web-server --config ./config.yaml
```
<h6>-------------文中有许多借鉴于别人的blog进行集成的的地方，不足之处多指教-------------------</h6>
<h6>-------------文中有许多借鉴于别人的blog进行集成的的地方，不足之处多指教-------------------</h6>
<h6>-------------文中有许多借鉴于别人的blog进行集成的的地方，不足之处多指教-------------------</h6>
