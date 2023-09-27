package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oatmi/stock/data"
	"github.com/oatmi/stock/data/sqlite"
	"github.com/oatmi/stock/lib"
	"github.com/spf13/cast"
)

type AisudaiResponse struct {
	Status  int         `json:"status"` // 状态码，为0表示成功
	Message string      `json:"msg"`    // 文案信息
	Data    interface{} `json:"data"`   // 返回值内容
}

type AisudaiCRUDData struct {
	Count int         `json:"count"`
	Rows  interface{} `json:"rows"`
}

type StockItem struct {
	sqlite.Stock
	ShowValue int `json:"show_value"`
	ShowPrice int `json:"show_price"`
}

// GetStocks 获取库存数据
func GetStocks(c *gin.Context) {
	query := sqlite.New(data.Sqlite3)

	listParam := buildListStockParams(c)
	list, err := query.ListStocks(c, listParam)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Message: "无数据 " + err.Error()})
		return
	}
	if len(list) == 0 {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "无数据"})
		return
	}

	var countParam sqlite.CountStocksParams
	listParamByte, _ := json.Marshal(listParam)
	json.Unmarshal(listParamByte, &countParam)

	count, err := query.CountStocks(c, countParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	resp := []StockItem{}
	for _, stock := range list {
		stockItem := StockItem{
			Stock:     stock,
			ShowPrice: 0,
			ShowValue: 0,
		}

		if strings.Contains("chengnanyu,wangbo,likaihou,chenyu", lib.UserName(c)) {
			stockItem.ShowValue = 1
			stockItem.ShowPrice = 1
		}

		if stock.ProductType == 4 && lib.UserName(c) == "zhangling" {
			stockItem.ShowValue = 1
			stockItem.ShowPrice = 1
		}

		resp = append(resp, stockItem)
	}

	c.JSON(http.StatusOK, AisudaiCRUDData{Count: int(count), Rows: resp})
	return
}

func buildListStockParams(c *gin.Context) sqlite.ListStocksParams {
	arg := sqlite.ListStocksParams{
		Limit: sql.NullInt32{
			Int32: 10,
			Valid: true,
		},
		Offset: sql.NullInt32{
			Int32: (cast.ToInt32(c.DefaultQuery("page", "0")) - 1) * 10,
			Valid: true,
		},
	}

	if val, ok := c.GetQuery("name"); ok && val != "" {
		arg.Name = sql.NullString{
			String: "%" + val + "%",
			Valid:  true,
		}
	}
	if val, ok := c.GetQuery("product_type"); ok && val != "" {
		arg.ProductType = sql.NullInt32{
			Int32: cast.ToInt32(val),
			Valid: true,
		}
	}
	if val, ok := c.GetQuery("product_attr"); ok && val != "" {
		arg.ProductAttr = sql.NullInt32{
			Int32: cast.ToInt32(val),
			Valid: true,
		}
	}
	if val, ok := c.GetQuery("status"); ok && val != "" {
		arg.Status = sql.NullInt32{
			Int32: cast.ToInt32(val),
			Valid: true,
		}
	} else {
		arg.Status = sql.NullInt32{
			Int32: 1,
			Valid: true,
		}
	}

	return arg
}
