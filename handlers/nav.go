package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Nav struct {
	Label    string     `json:"label"`
	To       string     `json:"to"`
	Icon     string     `json:"icon"`
	Unfolded bool       `json:"unfolded"`
	Children []Children `json:"children"`
}
type Children struct {
	Label string `json:"label"`
	To    string `json:"to"`
}

func Navs(c *gin.Context) {
	userName, err := c.Cookie("stock_un")
	if err != nil || userName == "" {
		c.JSON(http.StatusOK, AisudaiResponse{})
		return
	}

	if userName == "yangtao" ||
		userName == "chengnanyu" ||
		userName == "wangbo" ||
		userName == "likaihou" ||
		userName == "chenyu" {
		nav := []Nav{
			{Label: userName, To: "/view/home", Icon: "fa fa-user", Unfolded: true, Children: []Children{}},
			{Label: "库存管理", To: "/view/stock", Icon: "fa fa-boxes-stacked"},
			{Label: "入库管理", To: "/view/in", Icon: "fa fa-inbox"},
			{Label: "出库管理", To: "/view/out", Icon: "fa fa-right-from-bracket"},
			{Label: "出库审核", To: "/view/out_approve", Icon: "fa fa-right-from-bracket"},
			{Label: "定价管理", To: "/view/price", Icon: "fa fa-tag"},
		}
		c.JSON(http.StatusOK, AisudaiResponse{Data: nav})
	} else if userName == "xiaoyang" {
		nav := []Nav{
			{Label: userName, To: "/view/home", Icon: "fa fa-user"},
			{Label: "库存管理", To: "/view/stock", Icon: "fa fa-boxes-stacked"},
			{Label: "入库管理", To: "/view/in", Icon: "fa fa-inbox"},
			{Label: "出库管理", To: "/view/out", Icon: "fa fa-right-from-bracket"},
			{Label: "出库审核", To: "/view/out_approve", Icon: "fa fa-right-from-bracket"},
		}
		c.JSON(http.StatusOK, AisudaiResponse{Data: nav})
	} else if userName == "zhouxiaoli" {
		nav := []Nav{
			{Label: userName, To: "/view/home", Icon: "fa fa-user"},
			{Label: "入库管理", To: "/view/in", Icon: "fa fa-inbox"},
		}
		c.JSON(http.StatusOK, AisudaiResponse{Data: nav})
	} else if userName == "zhangling" {
		nav := []Nav{
			{Label: userName, To: "/view/home", Icon: "fa fa-user"},
			{Label: "入库管理", To: "/view/in", Icon: "fa fa-inbox"},
			{Label: "出库审核", To: "/view/out_approve", Icon: "fa fa-right-from-bracket"},
			{Label: "定价管理", To: "/view/price", Icon: "fa fa-tag"},
		}
		c.JSON(http.StatusOK, AisudaiResponse{Data: nav})
	} else if userName == "chenhua" {
		nav := []Nav{
			{Label: userName, To: "/view/home", Icon: "fa fa-user"},
			{Label: "入库管理", To: "/view/in", Icon: "fa fa-inbox"},
		}
		c.JSON(http.StatusOK, AisudaiResponse{Data: nav})
	} else {
		c.JSON(http.StatusOK, AisudaiResponse{})
	}
}
