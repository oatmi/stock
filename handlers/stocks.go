package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oatmi/stock/data"
	"github.com/oatmi/stock/data/sqlite"
)

// GetStocks 获取库存数据
func GetStocks(c *gin.Context) {
	query := sqlite.New(data.Sqlite3)
	list, err := query.ListStocks(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, list)
	return
}
