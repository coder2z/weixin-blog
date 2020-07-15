package test

import (
	"fmt"
	"testing"
	"wx-blog/weix"
)

func TestGetToken(t *testing.T) {
	wx := weix.WeiX{
		AppID:     "wx4f7b7433b130cd99",
		AppSecret: "defede7ff66c746a6ca09b8ab48ce211",
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
