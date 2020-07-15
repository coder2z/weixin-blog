# 博客微信公众号推送

主要是把博客的发布的新文章推送到微信公众号。

# 项目启动

## 配置config

在config下的config.yml文件。

## 启动爬虫
```
cd cmd
go build spider.go
./spider
```
## 发送公众号
```
cd cmd
go build wx.go
./wx
```

# 定时发送

可以使用liunx的crontab操作

```vim /etc/crontab```

在末尾添加

//每天凌晨2点执行

```* 02 * * * user-name cd /{项目目录}/cmd && ./spider```

//每天凌晨2点10分执行

```10 02 * * * user-name cd /{项目目录}/cmd && ./wx```



