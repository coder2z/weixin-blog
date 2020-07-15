package main

import (
	"log"
	"wx-blog/config"
	Redis "wx-blog/redis"
	"wx-blog/weix"
)

func main() {
	i := config.Config{}
	conf := i.GetConf()
	redis := Redis.NewRedis(conf)

	//获取最新的文章：
	data, err := redis.HGetAll("BlogUrl_req").Result()

	if err != nil {
		log.Panic(err.Error())
	}

	if data == nil {
		log.Panic("没有新文章")
		return
	}

	//封装发送的数据
	var message string
	for key, value := range data {
		message += key + ":" + value + "<br>"
	}

	//初始化微信
	wx := weix.WeiX{
		AppSecret: conf.WX.AppSecret,
		AppID:     conf.WX.AppID,
	}
	//获取Token
	err = wx.GetToken()
	if err != nil {
		log.Panic(err.Error())
	}

	//封装消息
	sendMessage := `{
   "filter":{
      "is_to_all":false,
      "tag_id":100
   },
   "text":{
      "content":` + message + `
   },
    "msgtype":"text"
}`

	//发送消息
	err = wx.Send(sendMessage)
	if err != nil {
		log.Panic(err.Error())
	}

	//待发数据
	redis.HDel("BlogUrl_req")

	log.Println("文章发送成功!")
}
