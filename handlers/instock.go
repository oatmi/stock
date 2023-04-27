package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oatmi/stock/data"
	"github.com/oatmi/stock/data/sqlite"
	"github.com/spf13/cast"
)

type ApplicationItem struct {
	ID              int32  `json:"id"`
	StockName       string `json:"stock_name"`
	StockID         int32  `json:"stock_id"`
	Number          int32  `json:"number"`
	ApplicationDate string `json:"application_date"`
	BatchNoIn       string `json:"batch_no_in"`
	Status          int32  `json:"status"`
	ApplicationUser string `json:"application_user"`
	ApproveUser     string `json:"approve_user"`
	ApproveDate     string `json:"approve_date"`
	CreateDate      string `json:"create_date"`
}

// GetStocks 获取库存数据
func GetApplications(c *gin.Context) {
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

	list, err := query.ListApplications(c, buildApplicationParams(c))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if len(list) == 0 {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "无数据"})
		return
	}

	var resp []ApplicationItem
	for _, s := range list {
		stock, err := query.StocksByID(c, s.StockID)
		if err != nil {
			continue
		}

		resp = append(resp, ApplicationItem{
			ID:              s.ID,
			StockName:       stock.Name,
			StockID:         s.StockID,
			Number:          stock.StockNum,
			ApplicationDate: s.ApplicationDate,
			BatchNoIn:       s.BatchNoIn,
			Status:          s.Status,
			ApplicationUser: s.ApplicationUser,
			ApproveUser:     s.ApproveUser,
			ApproveDate:     s.ApproveDate,
			CreateDate:      s.CreateDate,
		})
	}

	c.JSON(http.StatusOK, AisudaiCRUDData{Count: int(count), Rows: resp})
	return
}

// buildApplicationParams 构建入库申请查询参数
//
// page=1&date=1675180800,1677599999&user_name=test&status=1&perPage=10
func buildApplicationParams(ctx *gin.Context) sqlite.ListApplicationsParams {
	arg := sqlite.ListApplicationsParams{}
	if val, ok := ctx.GetQuery("application_date"); ok && val != "" {
		arrData := strings.Split(val, ",")
		if len(arrData) == 2 {
			start := cast.ToInt(arrData[0])
			if start > 0 {
				arg.ApplicationDateS = sql.NullString{
					String: arrData[0],
					Valid:  true,
				}
			}

			end := cast.ToInt(arrData[1])
			if end > 0 {
				arg.ApplicationDateS = sql.NullString{
					String: arrData[1],
					Valid:  true,
				}
			}
		}
	}

	if val, ok := ctx.GetQuery("application_user"); ok && val != "" {
		arg.ApplicationUser = sql.NullString{
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
