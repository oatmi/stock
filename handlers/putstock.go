package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oatmi/stock/wechat"
)

const wechatBotWebHook = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=3abf3cf5-ba21-44e8-b5d8-60abd83f9c7a"

func PutStock(c *gin.Context) {
	// TODO 写入数据库

	markdownMessage := wechat.Markdown{
		Content: wechat.MarkdownTemplate01,
	}
	err := wechat.SendMarkdownMessage(c, wechatBotWebHook, markdownMessage)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "入库申请提交成功，发送消息失败（请勿重复申请）"})
		return
	}

	textMessage := wechat.Text{
		Content:       "有新的入库申请，请审批",
		MentionedList: []string{"yangtao"},
	}
	err = wechat.SendTextMessage(c, wechatBotWebHook, textMessage)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "入库申请提交成功，发送消息失败（请勿重复申请）"})
		return
	}

	c.JSON(http.StatusOK, nil)
}
