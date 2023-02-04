package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oatmi/stock/data"
	"github.com/oatmi/stock/handlers"
)

func main() {
	// https://gin-gonic.com/docs/examples/grouping-routes/
	router := gin.Default()

	router.LoadHTMLGlob("template/*")

	view := router.Group("/view")
	{
		view.GET("/stock", func(c *gin.Context) { c.HTML(http.StatusOK, "stock.html", nil) })
	}

	api := router.Group("/api")
	{
		api.GET("/home", handlers.GetStocks)
	}

	router.Static("/sdk", "./sdk")

	// start sqlite connection
	data.SqliteMustInit()

	router.Run("0.0.0.0:8080")
}
