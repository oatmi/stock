package wechat

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type TextMessage struct {
	Msgtype string `json:"msgtype"`
	Text    Text   `json:"text"`
}

type Text struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

// SendTextMessage 发送一条文本消息
//
//	curl 'https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=3abf3cf5-ba21-44e8-b5d8-60abd83f9c7a' \
//	   -H 'Content-Type: application/json' \
//	   -d '
//	   {
//	    "msgtype": "text",
//	    "text": {
//	        "content": "广州今日天气：29度，大\n部分多云，降雨概率：60%",
//	        "mentioned_list":["yangtao","@all"],
//	    }
//	}'
func SendTextMessage(ctx context.Context, webHook string, message Text) error {
	msg := TextMessage{
		Msgtype: "text",
		Text:    message,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(msg)
	if err != nil {
		return err
	}

	resp, err := http.Post(webHook, "application/json", &buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

type MarkdownMessgage struct {
	Msgtype  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
}

type Markdown struct {
	Content string `json:"content"`
}

const MarkdownTemplate01 = `### 新的入库申请
发起人：<font color="info">杨涛</font>

库存申请列表:
> 30g无纺布 x 1
> 30g无纺布 x 1
> 30g无纺布 x 1
> 30g无纺布 x 1
> 30g无纺布 x 1

前往 [审核地址]() 进行审核。
`

// SendMarkdownMessage 发送一条`markdown`消息
func SendMarkdownMessage(ctx context.Context, webHook string, message Markdown) error {
	msg := MarkdownMessgage{
		Msgtype:  "markdown",
		Markdown: message,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(msg)
	if err != nil {
		return err
	}

	resp, err := http.Post(webHook, "application/json", &buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
