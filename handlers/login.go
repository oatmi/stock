package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oatmi/stock/data/account"
)

const (
	CorpID = "ww0980d2723fb39cdb"
)

type LoginRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// Login 系统入口
//
// https://developer.work.weixin.qq.com/document/path/98151
func Login(ctx *gin.Context) {
	var login LoginRequest
	err := ctx.BindJSON(&login)
	if err != nil {
		ctx.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E100] 参数错误"})
		return
	}

	for u, p := range account.Accounts {
		if u == login.User {
			if p == login.Password {
				ctx.SetCookie("stock_un", login.User, 86400, "/", ctx.Request.Host, false, true)
				ctx.JSON(http.StatusOK, AisudaiResponse{})
				return
			} else {
				ctx.JSON(http.StatusUnauthorized, AisudaiResponse{Status: 1, Message: "账号密码错误"})
				return
			}
		}
	}

	ctx.JSON(http.StatusUnauthorized, AisudaiResponse{Status: 1, Message: "账号密码错误"})
}
