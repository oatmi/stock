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

	c.JSON(http.StatusOK, AisudaiCRUDData{Count: int(count), Rows: list})
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
