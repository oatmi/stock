package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oatmi/stock/data"
	"github.com/oatmi/stock/data/sqlite"
	"github.com/spf13/cast"
)

type ApproveOutRequest struct {
	ID       int    `json:"id"`
	Stockids string `json:"stockids"`
	Status   int    `json:"status"`
}

func ApproveOut(c *gin.Context) {
	var req ApproveOutRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E100] 参数错误"})
		return
	}

	stockID := strings.Split(req.Stockids, ",")
	if len(stockID) == 0 {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E101] 参数错误"})
		return
	}

	query := sqlite.New(data.Sqlite3)
	if req.Status == 1 {
		for _, strID := range stockID {
			updateStockParam := sqlite.UpdateStockStatusByIDParams{
				Status: 3,
				ID:     int32(cast.ToInt(strID)),
			}
			err := query.UpdateStockStatusByID(c, updateStockParam)
			if err != nil {
				c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E101] 更新库存失败"})
				return
			}
		}

		updateApproveParam := sqlite.UpdateApplicationOUTParams{
			Status: 3,
			ID:     int32(req.ID),
		}
		err = query.UpdateApplicationOUT(c, updateApproveParam)
		if err != nil {
			c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E102] 更新库存失败"})
			return
		}
	} else {
		updateApproveParam := sqlite.UpdateApplicationOUTParams{
			Status: 4,
			ID:     int32(req.ID),
		}
		err = query.UpdateApplicationOUT(c, updateApproveParam)
		if err != nil {
			c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E104] 更新库存失败"})
			return
		}
	}

	c.JSON(http.StatusOK, nil)
}
