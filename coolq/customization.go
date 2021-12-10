package coolq

import (
	"bytes"
	"github.com/Mrs4s/MiraiGo/message"
	"io/ioutil"
	"net/http"
	"strings"
)

var SendQQ = func(a int64, b interface{}, groupID int64) {

}
var SendQQGroup = func(a int64, b int64, c interface{}) {

}

func (bot *CQBot) InitSend() {
	SendQQ = func(uid int64, msg interface{}, groupID int64) {
		if bot == nil {
			return
		}
		if uid == 0 {
			return
		}

		switch msg.(type) {
		case string:
			if bot != nil {
				if strings.Contains(msg.(string), "data:image") {
					bot.SendPrivateMessage(uid, groupID, &message.SendingMessage{Elements: []message.IMessageElement{&LocalImageElement{File: "./output.jpg"}}})
				} else {
					bot.SendPrivateMessage(uid, groupID, &message.SendingMessage{Elements: []message.IMessageElement{&message.TextElement{Content: msg.(string)}}})
				}
			}
		case *http.Response:
			data, _ := ioutil.ReadAll(msg.(*http.Response).Body)
			bot.SendPrivateMessage(uid, groupID, &message.SendingMessage{Elements: []message.IMessageElement{&LocalImageElement{Stream: bytes.NewReader(data)}}})
		}

	}
	SendQQGroup = func(gid int64, uid int64, msg interface{}) {
		if bot == nil {
			return
		}
		switch msg.(type) {
		case string:
			if bot != nil {
				bot.SendGroupMessage(gid, &message.SendingMessage{Elements: []message.IMessageElement{&message.AtElement{Target: uid}, &message.TextElement{Content: msg.(string)}}})
			}
		case *http.Response:
			data, _ := ioutil.ReadAll(msg.(*http.Response).Body)
			bot.SendGroupMessage(gid, &message.SendingMessage{Elements: []message.IMessageElement{&message.AtElement{Target: uid}, &message.TextElement{Content: "\n"}, &LocalImageElement{Stream: bytes.NewReader(data)}}})
		}
	}
}
