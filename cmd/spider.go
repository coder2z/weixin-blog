package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"strings"
	"wx-blog/config"
	Redis "wx-blog/redis"
	"wx-blog/utils"
)

func main() {

	i := config.Config{}
	conf := i.GetConf()
	redis := Redis.NewRedis(conf)

	//爬虫采集 收集到redis

	//获取需要爬取的地址
	webList := redis.HGetAll("BlogUrl").Val()
	for key, value := range webList {
		c := colly.NewCollector(
			colly.AllowedDomains(strings.Split(key, "//")[1]),
		)
		c.OnRequest(func(request *colly.Request) {
			//去重
			ok, _ := redis.HExists("BlogUrl", utils.GetMd5(request.URL.String())).Result()
			if ok {
				request.Abort()
				return
			}
			fmt.Println(request.URL.String())
		})
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			//获取所有a
			_ = e.Request.Visit(e.Attr("href"))

		})
		c.OnHTML("title", func(e *colly.HTMLElement) {

			//文章页面正则提取
			matched, _ := regexp.MatchString(value, e.Request.URL.String())
			if matched {

				//最新的连接存储到redis
				redis.HSet("BlogUrl_req", e.Text, e.Request.URL.String())

				//去记录重库
				redis.HSet("BlogUrl_db", utils.GetMd5(e.Request.URL.String()), e.Text)

			}
		})
		_ = c.Visit(key)
	}
}
