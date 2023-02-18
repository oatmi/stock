package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oatmi/stock/data"
	"github.com/oatmi/stock/data/sqlite"
)

type ApproveINRequest struct {
	ID        int    `json:"id"`
	Status    int    `json:"status"`
	BatchNoIn string `json:"batch_no_in"`
}

// ApproveIN 入库申请
func ApproveIN(c *gin.Context) {
	var req ApproveINRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E100] 参数错误"})
		return
	}

	query := sqlite.New(data.Sqlite3)
	if req.Status == 1 {
		updateStockParam := sqlite.UpdateStocksParams{
			Status:    1,
			BatchNoIn: req.BatchNoIn,
			Status_2:  2,
		}
		err := query.UpdateStocks(c, updateStockParam)
		if err != nil {
			c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E101] 更新库存失败"})
			return
		}

		updateApproveParam := sqlite.UpdateApplicationINParams{
			Status: 3,
			ID:     int32(req.ID),
		}
		err = query.UpdateApplicationIN(c, updateApproveParam)
		if err != nil {
			c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E102] 更新库存失败"})
			return
		}
	} else {
		updateStockParam := sqlite.UpdateStocksParams{
			Status:    4,
			BatchNoIn: req.BatchNoIn,
			Status_2:  2,
		}
		err := query.UpdateStocks(c, updateStockParam)
		if err != nil {
			c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E103] 更新库存失败"})
			return
		}

		updateApproveParam := sqlite.UpdateApplicationINParams{
			Status: 4,
			ID:     int32(req.ID),
		}
		err = query.UpdateApplicationIN(c, updateApproveParam)
		if err != nil {
			c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E104] 更新库存失败"})
			return
		}
	}

	c.JSON(http.StatusOK, nil)
}
