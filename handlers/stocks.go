package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oatmi/stock/data"
	"github.com/oatmi/stock/data/sqlite"
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

	c.JSON(http.StatusOK, AisudaiCRUDData{Count: int(count), Rows: list})
	return
}

func buildListStockParams(c *gin.Context) sqlite.ListStocksParams {
	arg := sqlite.ListStocksParams{
		Limit: sql.NullInt32{
			Int32: (cast.ToInt32(c.DefaultQuery("page", "0")) - 1) * 10,
			Valid: true,
		},
		Offset: sql.NullInt32{
			Int32: 10,
			Valid: true,
		},
	}

	if val, ok := c.GetQuery("name"); ok && val != "" {
		arg.Name = sql.NullString{
			String: "%" + val + "%",
			Valid:  true,
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
	fmt.Printf("debug: %+v\n", arg)

	return arg
}
