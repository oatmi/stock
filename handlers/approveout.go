package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oatmi/stock/data"
	"github.com/oatmi/stock/data/sqlite"
	"github.com/oatmi/stock/lib"
)

type ApproveOutRequest struct {
	ID      int `json:"id"`
	StockID int `json:"stock_id"`
	Status  int `json:"status"`
}

// ApproveOut 出库审核
//
// 1. 分别获取出库表和库存表的数据
// 2. 基础的数据判断，如出库数量、剩余数量等
// 3. 减少库存数量
// 4. 修改出库表状态
func ApproveOut(c *gin.Context) {
	var req ApproveOutRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E100] 参数错误"})
		return
	}

	query := sqlite.New(data.Sqlite3)

	stock, err := query.StocksByID(c, int32(req.StockID))
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E100] 参数错误"})
		return
	}

	out, err := query.OutApplicationByID(c, int32(req.ID))
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E100] 参数错误"})
		return
	}

	left := stock.CurrentNum - out.Number
	if left < 0 {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E100] 剩余库存不足"})
		return
	}

	if req.Status == 1 {
		updateStockParam := sqlite.UpdateStockNumberParams{
			CurrentNum: left,
			ID:         int32(req.StockID),
			Value:      stock.Price * left,
		}
		err := query.UpdateStockNumber(c, updateStockParam)
		if err != nil {
			c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E101] 更新库存失败"})
			return
		}

		updateApproveParam := sqlite.UpdateApplicationOUTParams{
			Status:      3,
			ID:          int32(req.ID),
			ApproveUser: lib.UserName(c),
		}
		err = query.UpdateApplicationOUT(c, updateApproveParam)
		if err != nil {
			c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E102] 更新库存失败"})
			return
		}
	} else {
		updateApproveParam := sqlite.UpdateApplicationOUTParams{
			Status:      4,
			ID:          int32(req.ID),
			ApproveUser: lib.UserName(c),
		}
		err = query.UpdateApplicationOUT(c, updateApproveParam)
		if err != nil {
			c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "[E104] 更新库存失败"})
			return
		}
	}

	c.JSON(http.StatusOK, AisudaiResponse{})
}
