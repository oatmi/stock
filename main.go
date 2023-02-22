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
		view.GET("/in", func(c *gin.Context) { c.HTML(http.StatusOK, "in.html", nil) })
		view.GET("/app", func(c *gin.Context) { c.HTML(http.StatusOK, "application.html", nil) })
		view.GET("/out", func(c *gin.Context) { c.HTML(http.StatusOK, "out.html", nil) })
		view.GET("/out_approve", func(c *gin.Context) { c.HTML(http.StatusOK, "out_approve.html", nil) })
		view.GET("/price", func(c *gin.Context) { c.HTML(http.StatusOK, "price.html", nil) })
	}

	api := router.Group("/api")
	{
		api.GET("/home", handlers.GetStocks)
		api.GET("/instock", handlers.GetApplications)
		api.GET("/outstock", handlers.OutStockList)

		api.POST("/put/stock", handlers.PutStock)
		api.POST("/out/stock", handlers.OutStockCreate)
		api.POST("/approvein", handlers.ApproveIN)
		api.POST("/approveout", handlers.ApproveOut)
		api.POST("/alterprice", handlers.AlterPrice)
	}

	router.Static("/sdk", "./sdk")

	data.SqliteMustInit()

	router.Run("0.0.0.0:8888")
}
