# 博客微信公众号推送

主要是把博客的发布的新文章推送到微信公众号。

# 项目启动

## 配置config

在config下的config.yml文件。

## 在redis中添加需要爬虫的网站

```
//key固定为key
//field为需要爬虫的网站
//field为需要爬取的数据的路由正则（比如：https://myxy99.cn/posts/? 就是爬取所有路由为https://myxy99.cn/posts开始的）
hset key field value

// 示例
// hset BlogUrl https://myxy99.cn https://myxy99.cn/posts/?

```





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

每天凌晨2点执行

```* 02 * * * user-name cd /{项目目录}/cmd && ./spider```

每天凌晨2点10分执行

```10 02 * * * user-name cd /{项目目录}/cmd && ./wx```



