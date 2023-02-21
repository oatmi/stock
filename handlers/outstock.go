package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oatmi/stock/data"
	"github.com/oatmi/stock/data/sqlite"
	"github.com/oatmi/stock/lib"
	"github.com/spf13/cast"
)

type OutStockRequest struct {
	IDs string `json:"ids"`
}

// OutStockCreate 创建出库申请单
func OutStockCreate(c *gin.Context) {
	var stocks OutStockRequest
	err := c.BindJSON(&stocks)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E100] 参数错误"})
		return
	}

	query := sqlite.New(data.Sqlite3)
	param := sqlite.CreateOutApplicationParams{
		Stockids:        stocks.IDs,
		Status:          2,
		ApplicationUser: "admin",
		ApproveUser:     "yangtao",
		CreateDate:      lib.CurrentDate(),
	}
	err = query.CreateOutApplication(c, param)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E101] 数据写入错误"})
		return
	}

	c.JSON(http.StatusOK, AisudaiResponse{Status: 0, Message: "出库申请提交成功"})
}

func OutStockList(c *gin.Context) {
	query := sqlite.New(data.Sqlite3)

	count, err := query.CountStocks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if count == 0 {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "无数据"})
		return
	}

	list, err := query.ListOutApplications(c, buildOutApplicationParams(c))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if len(list) == 0 {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "无数据"})
		return
	}

	c.JSON(http.StatusOK, AisudaiCRUDData{Count: int(count), Rows: list})
	return
}

func buildOutApplicationParams(ctx *gin.Context) sqlite.ListOutApplicationsParams {
	arg := sqlite.ListOutApplicationsParams{}
	if val, ok := ctx.GetQuery("application_user"); ok && val != "" {
		arg.ApplicationUser = sql.NullString{
			String: val,
			Valid:  true,
		}
	}

	if val, ok := ctx.GetQuery("approve_user"); ok && val != "" {
		arg.ApproveUser = sql.NullString{
			String: val,
			Valid:  true,
		}
	}

	if val, ok := ctx.GetQuery("status"); ok && val != "" {
		arg.Status = sql.NullInt32{
			Int32: cast.ToInt32(val),
			Valid: true,
		}
	}

	return arg
}
