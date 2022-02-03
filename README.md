# Docker Proxy (HTTP Api)

使用Http请求进行常见Docker操作

使用时需要在同目录的conf文件夹下的config.json中添加配置

依赖：MongoDB， Redis

v0.1.1 还很简陋，支部更新中

示例配置文件

``` json
{
  "Server": {
    "port": 8080,
    "host": "0.0.0.0",
    "mode": "release"
  },
  "Mongo": {
    "Host": "localhost",
    "Port": 27017,
    "User": "root",
    "Password": "chino"
  },
  "Redis": {
    "Host": "localhost",
    "Port": 6379,
    "Password": ""
  }
}
```

