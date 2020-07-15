package weix

import (
	"encoding/json"
	"errors"
	"fmt"
	"wx-blog/request"
)

const (
	WxGetAccessToken = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v"
	WxSendText       = "https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=%v"
)

type WeiX struct {
	AppID       string
	AppSecret   string
	AccessToken string
}

func (w *WeiX) GetToken() (err error) {
	url := fmt.Sprintf(WxGetAccessToken, w.AppID, w.AppSecret)
	i := &request.Request{}
	i.Call("GET", url, nil)
	if i.Err == nil {
		var data map[string]interface{}
		err = json.Unmarshal(i.Body, &data)
		if err == nil {
			a := data["access_token"]
			if a != nil {
				w.AccessToken = data["access_token"].(string)
				return
			}
			err = errors.New("请检查AppID，AppSecret")
			return
		}
	}
	return
}

func (w *WeiX) Send(message string) (err error) {
	url := fmt.Sprintf(WxSendText, w.AccessToken)
	i := &request.Request{}
	i.Call("POST", url, []byte(message))
	if i.Err == nil {
		var data map[string]interface{}
		err = json.Unmarshal(i.Body, &data)
		if err == nil {
			fmt.Println(data)
		}
	}
	return
}
