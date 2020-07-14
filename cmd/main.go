package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"strings"
	Redis "wx-blog/redis"
)

func getMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {

	redis := Redis.NewRedis()

	//爬虫采集 收集到redis

	webList:=redis.HGetAll("BlogUrl").Val()

	for key, value := range webList {
		c := colly.NewCollector(
			colly.AllowedDomains(strings.Split(key, "//")[1]),
		)
		// Find and visit all links
		c.OnRequest(func(request *colly.Request) {
			//分析最新的连接
			ok, _ := redis.HExists("BlogUrl", getMd5(request.URL.String())).Result()
			if ok {
				request.Abort()
				return
			}
			fmt.Println(request.URL.String())
		})

		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			_ = e.Request.Visit(e.Attr("href"))

		})
		c.OnHTML("title", func(e *colly.HTMLElement) {
			matched, _ := regexp.MatchString(value, e.Request.URL.String())
			if matched {
				//缓存到发送
				//最新的连接存储到redis

				redis.HSet("BlogUrl_req", e.Text, e.Request.URL.String())

				//去记录重库
				redis.HSet("BlogUrl_db", getMd5(e.Request.URL.String()), e.Text)

			}
		})
		_ = c.Visit(key)
	}

	//公众号推送，根据最新的redis
}
