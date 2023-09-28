package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oatmi/stock/data"
	"github.com/oatmi/stock/data/sqlite"
	"github.com/oatmi/stock/lib"
	"github.com/oatmi/stock/wechat"
)

type AlterPriceRequest struct {
	ID    int `json:"id"`
	Price int `json:"price"`
}

// AlterPrice 修改定价
//
// 定价修改之后，对应的价值也需要调整
func AlterPrice(c *gin.Context) {
	var req AlterPriceRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E100] 参数错误"})
		return
	}

	query := sqlite.New(data.Sqlite3)

	stock, err := query.StocksByID(c, int32(req.ID))
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E100] 参数错误"})
		return
	}

	param := sqlite.UpdateStockPriceByIDParams{
		Price: int32(req.Price),
		ID:    int32(req.ID),
		Value: stock.CurrentNum * int32(req.Price),
	}
	err = query.UpdateStockPriceByID(c, param)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E101] 更新价格失败"})
		return
	}

	message := wechat.Text{
		Content: fmt.Sprintf("【通知】%s修改了%s的价格。", lib.UserName(c), stock.Name),
	}

	wechat.SendTextMessage(c, wechatBotWebHook, message)

	c.JSON(http.StatusOK, nil)
}
