package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oatmi/stock/data"
	"github.com/oatmi/stock/data/sqlite"
	"github.com/oatmi/stock/lib"
	"github.com/spf13/cast"
)

type ApproveINRequest struct {
	ID        int    `json:"id"`
	Status    int    `json:"status"`
	Price     int    `json:"price"`
	Number    int    `json:"number"`
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

	application, err := query.ApplicationByID(c, int32(req.ID))
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E101] 参数错误"})
		return
	}
	if application.Status > 2 {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E102] 已审核"})
		return
	}

	if req.Status == 1 {
		updateStockParam := sqlite.UpdateStockStatusByIDParams{
			Status: 1,
			ID:     application.StockID,
			Price:  cast.ToInt32(req.Price),
			Value:  cast.ToInt32(req.Price * req.Number),
		}
		err := query.UpdateStockStatusByID(c, updateStockParam)
		if err != nil {
			c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E101] 更新库存失败"})
			return
		}

		updateApproveParam := sqlite.UpdateApplicationINParams{
			Status:      3,
			ID:          int32(req.ID),
			ApproveUser: lib.UserName(c),
		}
		err = query.UpdateApplicationIN(c, updateApproveParam)
		if err != nil {
			c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E102] 更新库存失败"})
			return
		}
	} else {
		updateStockParam := sqlite.UpdateStockStatusByIDParams{
			Status: 4,
			ID:     application.StockID,
		}
		err := query.UpdateStockStatusByID(c, updateStockParam)
		if err != nil {
			c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E101] 更新库存失败"})
			return
		}

		updateApproveParam := sqlite.UpdateApplicationINParams{
			Status:      4,
			ApproveUser: lib.UserName(c),
			ID:          int32(req.ID),
		}
		err = query.UpdateApplicationIN(c, updateApproveParam)
		if err != nil {
			c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E104] 更新库存失败"})
			return
		}
	}

	c.JSON(http.StatusOK, nil)
}
