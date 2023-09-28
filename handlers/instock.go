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
	CanApprove      int    `json:"can_approve"`
}

// GetStocks 获取库存数据
func GetApplications(c *gin.Context) {
	query := sqlite.New(data.Sqlite3)

	listParam := buildApplicationParams(c)
	list, err := query.ListApplications(c, listParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if len(list) == 0 {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "无数据"})
		return
	}

	var countParam sqlite.CountApplicationsParams
	listParamByte, _ := json.Marshal(listParam)
	json.Unmarshal(listParamByte, &countParam)

	count, err := query.CountApplications(c, countParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if count == 0 {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "无数据"})
		return
	}

	var resp []ApplicationItem
	for _, s := range list {
		stock, err := query.StocksByID(c, s.StockID)
		if err != nil {
			continue
		}

		item := ApplicationItem{
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
			CanApprove:      0,
		}

		if lib.UserName(c) == s.ApproveUser {
			item.CanApprove = 1
		}

		resp = append(resp, item)
	}

	c.JSON(http.StatusOK, AisudaiCRUDData{Count: int(count), Rows: resp})
	return
}

// buildApplicationParams 构建入库申请查询参数
//
// page=1&date=1675180800,1677599999&user_name=test&status=1&perPage=10
func buildApplicationParams(ctx *gin.Context) sqlite.ListApplicationsParams {
	arg := sqlite.ListApplicationsParams{
		Limit: sql.NullInt32{
			Int32: 10,
			Valid: true,
		},
		Offset: sql.NullInt32{
			Int32: (cast.ToInt32(ctx.DefaultQuery("page", "0")) - 1) * 10,
			Valid: true,
		},
	}
	if val, ok := ctx.GetQuery("application_date"); ok && val != "" {
		arrData := strings.Split(val, ",")
		if len(arrData) == 2 {
			start := cast.ToInt(arrData[0])
			if start > 0 {
				arg.ApplicationDateS = sql.NullString{
					String: lib.TimestampToDate(cast.ToInt64(arrData[0])),
					Valid:  true,
				}
			}

			end := cast.ToInt(arrData[1])
			if end > 0 {
				arg.ApplicationDateE = sql.NullString{
					String: lib.TimestampToDate(cast.ToInt64(arrData[1])),
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
