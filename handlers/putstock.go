package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oatmi/stock/data"
	"github.com/oatmi/stock/data/sqlite"
	"github.com/oatmi/stock/lib"
	"github.com/spf13/cast"
)

const wechatBotWebHook = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=3abf3cf5-ba21-44e8-b5d8-60abd83f9c7a"

type PutStockRequest struct {
	BatchNoIn        string `json:"batch_no_in"`
	Name             string `json:"name"`
	ProductType      string `json:"product_type"`
	ProductAttr      string `json:"product_attr"`
	Supplier         string `json:"supplier"`
	Model            string `json:"model"`
	Unit             string `json:"unit"`
	Price            int    `json:"price"`
	WayIn            string `json:"way_in"`
	Location         string `json:"location"`
	BatchNoProduce   string `json:"batch_no_produce"`
	ProduceDate      string `json:"produce_date"`
	DisinfectionNo   string `json:"disinfection_no"`
	DisinfectionDate string `json:"disinfection_date"`
	StockNum         int    `json:"stock_num"`
}

func PutStock(c *gin.Context) {
	var stock PutStockRequest
	err := c.BindJSON(&stock)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E100] 参数错误"})
		return
	}

	query := sqlite.New(data.Sqlite3)
	createParam := sqlite.CreateStockParams{
		Status:           2, // 1: ok, 2: waitIN, 3: outed 4: declined
		Name:             stock.Name,
		ProductType:      cast.ToInt32(stock.ProductType),
		ProductAttr:      cast.ToInt32(stock.ProductAttr),
		Supplier:         stock.Supplier,
		Model:            stock.Model,
		Unit:             stock.Unit,
		Price:            cast.ToInt32(stock.Price),
		BatchNoIn:        stock.BatchNoIn,
		WayIn:            cast.ToInt32(stock.WayIn),
		Location:         cast.ToInt32(stock.Location),
		BatchNoProduce:   stock.BatchNoProduce,
		ProduceDate:      cast.ToInt32(stock.ProduceDate),
		DisinfectionNo:   stock.DisinfectionNo,
		DisinfectionDate: cast.ToInt32(stock.DisinfectionDate),
		StockDate:        int32(time.Now().Unix()),
		StockNum:         cast.ToInt32(stock.StockNum),
		CurrentNum:       cast.ToInt32(stock.StockNum),
		Value:            cast.ToInt32(stock.StockNum) * cast.ToInt32(stock.Price),
	}
	newStock, err := query.CreateStock(c, createParam)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E101] 提交申请失败" + err.Error()})
		return
	}

	application := sqlite.CreateStockApplicationParams{
		StockID:         newStock.ID,
		ApplicationDate: lib.CurrentDate(),
		BatchNoIn:       lib.NewBatchNO(),
		Status:          1,       // 1: initiate, 2: wait approve, 3: prooved, 4: rejected
		ApplicationUser: "admin", // TODO
		CreateDate:      lib.CurrentDate(),
	}
	err = query.CreateStockApplication(c, application)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E101] 提交申请失败" + err.Error()})
		return
	}

	// markdownMessage := wechat.Markdown{
	// 	Content: wechat.MarkdownTemplate01,
	// }
	// err := wechat.SendMarkdownMessage(c, wechatBotWebHook, markdownMessage)
	// if err != nil {
	// 	c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "入库申请提交成功，发送消息失败（请勿重复申请）"})
	// 	return
	// }

	// textMessage := wechat.Text{
	// 	Content:       "有新的入库申请，请审批",
	// 	MentionedList: []string{"yangtao"},
	// }
	// err = wechat.SendTextMessage(c, wechatBotWebHook, textMessage)
	// if err != nil {
	// 	c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "入库申请提交成功，发送消息失败（请勿重复申请）"})
	// 	return
	// }

	c.JSON(http.StatusOK, AisudaiResponse{Status: 0, Message: "提交成功"})
}
