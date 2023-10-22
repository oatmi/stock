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
	ID            int `json:"id"`
	Number        int `json:"number"`
	CurrentNumber int `json:"current_num"`
}

// OutStockCreate 创建出库申请单
func OutStockCreate(c *gin.Context) {
	var stocks OutStockRequest
	err := c.BindJSON(&stocks)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E100] 参数错误"})
		return
	}

	if stocks.Number <= 0 || stocks.CurrentNumber <= 0 {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E102] 库存数量错误"})
		return
	}

	if stocks.Number > stocks.CurrentNumber {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E103] 超过库存数量"})
		return
	}

	query := sqlite.New(data.Sqlite3)
	param := sqlite.CreateOutApplicationParams{
		Stockids:        cast.ToString(stocks.ID),
		Number:          int32(stocks.Number),
		Status:          2,
		ApplicationUser: lib.UserName(c),
		ApproveUser:     "",
		CreateDate:      lib.CurrentDate(),
	}
	err = query.CreateOutApplication(c, param)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E104] 数据写入错误"})
		return
	}

	c.JSON(http.StatusOK, AisudaiResponse{Status: 0, Message: "出库申请提交成功"})
}

type OutStockListItem struct {
	ID              int32  `json:"id"`
	Name            string `json:"name"`
	Number          int    `json:"number"`
	CurrentNumber   int    `json:"current_number"`
	StockID         int32  `json:"stock_id"`
	Status          int32  `json:"status"`
	ApplicationUser string `json:"application_user"`
	ApproveUser     string `json:"approve_user"`
	CreateDate      string `json:"create_date"`
}

func OutStockList(c *gin.Context) {
	query := sqlite.New(data.Sqlite3)

	listParam := buildOutApplicationParams(c)
	list, err := query.ListOutApplications(c, listParam)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: err.Error()})
		return
	}

	if len(list) == 0 {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "无数据"})
		return
	}

	var countParam sqlite.CountOutApplicationParams
	lib.ConvertTo(listParam, &countParam)
	count, err := query.CountOutApplication(c, countParam)

	if count == 0 || err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "无数据"})
		return
	}

	var resp []OutStockListItem
	for _, l := range list {
		s, err := query.StocksByID(c, cast.ToInt32(l.Stockids))
		if err != nil {
			continue
		}

		resp = append(resp, OutStockListItem{
			ID:              l.ID,
			Name:            s.Name,
			Number:          int(l.Number),
			CurrentNumber:   int(s.CurrentNum),
			StockID:         s.ID,
			Status:          l.Status,
			ApplicationUser: l.ApplicationUser,
			ApproveUser:     l.ApproveUser,
			CreateDate:      l.CreateDate,
		})
	}

	c.JSON(http.StatusOK, AisudaiCRUDData{Count: int(count), Rows: resp})
}

func buildOutApplicationParams(ctx *gin.Context) sqlite.ListOutApplicationsParams {
	arg := sqlite.ListOutApplicationsParams{
		Limit: sql.NullInt32{
			Int32: 10,
			Valid: true,
		},
		Offset: sql.NullInt32{
			Int32: (cast.ToInt32(ctx.DefaultQuery("page", "0")) - 1) * 10,
			Valid: true,
		},
	}
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
