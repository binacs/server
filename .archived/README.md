## 数据库配置相关

在启动 `server` 之前, 需要在 mysql 中先行建立DB, 使用 `collation : utf8mb4` 的配置并新建 testdb。

> `-e MYSQL_DATABASE=testdb` 与下列mysql指令等效：
>
> ```mysql
> CREATE SCHEMA `testdb` DEFAULT CHARACTER SET utf8mb4 ;
> ```

1. linux 可视化管理工具: Workbench

```shell
sudo apt-get install mysql-workbench
```

2. 对于 windows macOS: Navicat

此时在 host 项需配置 `docker-machine ip` 而非 `localhost` nor `127.0.0.1` 。