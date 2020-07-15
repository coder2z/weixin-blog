package test

import (
	"fmt"
	"testing"
	"wx-blog/config"
	"wx-blog/weix"
)

func TestGetToken(t *testing.T) {
	i := config.Config{}
	conf := i.GetConf()

	wx := weix.WeiX{
		AppSecret: conf.WX.AppSecret,
		AppID:     conf.WX.AppID,
	}

	err := wx.GetToken()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(wx.AccessToken)

	_ = wx.Send(`{
   "filter":{
      "is_to_all":false,
      "tag_id":100
   },
   "text":{
      "content":"CONTENT"
   },
    "msgtype":"text"
}`)
}
