package jpush

import (
	"fmt"
	"github.com/ylywyn/jpush-api-go-client"
)

const (
	appKey = "8483f050717eea0d2bb5d672"
	secret = "a07ce32cb4106635ac94d7de"
)

func Push(ad jpushclient.Audience, msg jpushclient.Message) {
	//Platform
	var pf jpushclient.Platform
	pf.Add(jpushclient.ANDROID)
	pf.Add(jpushclient.IOS)
	pf.Add(jpushclient.WINPHONE)
	//pf.All()

	//Audience
	//var ad jpushclient.Audience
	//s := []string{"140fe1da9e1214a0dae"}
	//ad.SetTag(s)
	//ad.SetAlias(s)
	//ad.SetID(s)
	//ad.All()

	//Notice
	var notice jpushclient.Notice
	notice.SetAlert("alert_test")
	notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: "AndroidNotice"})
	notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: "IOSNotice"})
	notice.SetWinPhoneNotice(&jpushclient.WinPhoneNotice{Alert: "WinPhoneNotice"})

	//var ext map[string]interface{}
	//ext = make(map[string]interface{})
	//ext["type"] = 44
	//ext["orderId"] = 342
	//var msg jpushclient.Message
	//msg.Title = "Hello"
	//msg.Content = "点的约会邀请，请按时邀约"
	//msg.Extras = ext
	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&ad)
	payload.SetMessage(&msg)
	payload.SetNotice(&notice)

	bytes, _ := payload.ToBytes()
	fmt.Printf("%s\r\n", string(bytes))

	//push
	c := jpushclient.NewPushClient(secret, appKey)
	str, err := c.Send(bytes)
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	} else {
		fmt.Printf("ok:%s", str)
	}
}
