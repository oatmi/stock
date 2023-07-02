package handlers

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/oatmi/stock/wechat"
)

const (
	CorpID = "ww0980d2723fb39cdb"
)

// Login 系统入口
//
// https://developer.work.weixin.qq.com/document/path/98151
func Login(ctx *gin.Context) {
	loginHandleURL := "http://test.stock.com:8888/view/login_callback"
	param := url.Values{}
	param.Add("login_type", "CorpApp")
	param.Add("agentid", "1000024")
	param.Add("appid", CorpID)
	param.Add("redirect_uri", loginHandleURL)
	URL := "https://login.work.weixin.qq.com/wwlogin/sso/login?" + param.Encode()
	fmt.Printf("debug: %+v\n", URL)
	ctx.Redirect(302, URL)
}

func LoginSuncess(ctx *gin.Context) {
	code := ctx.Query("code")
	fmt.Printf("debug: %+v\n", code)

	at, err := wechat.AccessToken(ctx)
	fmt.Printf("debug: %+v\n", at)
	fmt.Printf("debug: %+v\n", err)

	userID, err := wechat.GetUserID(ctx, at, code)
	fmt.Printf("debug: %+v\n", userID)
	fmt.Printf("debug: %+v\n", err)
}
