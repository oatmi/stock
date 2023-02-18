package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oatmi/stock/data"
	"github.com/oatmi/stock/data/sqlite"
	"github.com/oatmi/stock/lib"
	"github.com/spf13/cast"
)

const wechatBotWebHook = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=3abf3cf5-ba21-44e8-b5d8-60abd83f9c7a"

type PutStockRequest struct {
	Stock     []Stock `json:"stock"`
	BatchNoIn string  `json:"batch_no_in"`
}

type Stock struct {
	Name           string `json:"name"`
	ProductType    string `json:"product_type"`
	Type           string `json:"type"`
	Supplier       string `json:"supplier"`
	Model          string `json:"model"`
	Unit           string `json:"unit"`
	Price          string `json:"price"`
	WayIn          string `json:"way_in"`
	Location       string `json:"location"`
	BatchNoProduce string `json:"batch_no_produce"`
	ProduceDate    string `json:"produce_date"`
	StockNum       string `json:"stock_num"`
}

func PutStock(c *gin.Context) {
	var stocks PutStockRequest
	err := c.BindJSON(&stocks)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E100] 参数错误"})
		return
	}

	query := sqlite.New(data.Sqlite3)

	application := sqlite.CreateStockApplicationParams{
		ApplicationDate: lib.CurrentDate(),
		BatchNoIn:       stocks.BatchNoIn,
		Status:          1,       // 1: initiate, 2: wait approve, 3: prooved, 4: rejected
		ApplicationUser: "admin", // TODO
		CreateDate:      lib.CurrentDate(),
	}
	err = query.CreateStockApplication(c, application)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E101] 提交申请失败"})
		return
	}

	var countSucc, countFail int

	for _, s := range stocks.Stock {
		createParam := sqlite.CreateStockParams{
			Name:           s.Name,
			ProductType:    cast.ToInt32(s.ProductType),
			Type:           cast.ToInt32(s.Type),
			Supplier:       s.Supplier,
			Model:          s.Model,
			Unit:           s.Unit,
			Price:          cast.ToInt32(s.Price),
			BatchNoIn:      stocks.BatchNoIn,
			WayIn:          s.WayIn,
			Location:       s.Location,
			BatchNoProduce: s.BatchNoProduce,
			ProduceDate:    s.ProduceDate, // TODO 转成日期
			StockDate:      lib.CurrentDate(),
			StockNum:       cast.ToInt32(s.StockNum),
			Status:         2, // 1: ok, 2: waitIN, 3: outed
		}
		err := query.CreateStock(c, createParam)
		if err == nil {
			countSucc += 1
		} else {
			fmt.Printf("debug: %+v\n", err)
			countFail += 1
		}
	}

	message := fmt.Sprintf("成功提交%d条申请，失败%d条", countSucc, countFail)

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

	c.JSON(http.StatusOK, AisudaiResponse{Status: 0, Message: message})
}
