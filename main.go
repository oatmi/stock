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
	}

	api := router.Group("/api")
	{
		api.GET("/home", handlers.GetStocks)
		api.GET("/instock", handlers.GetApplications)
		api.POST("/put/stock", handlers.PutStock)
	}

	router.Static("/sdk", "./sdk")

	// start sqlite connection
	data.SqliteMustInit()
	// query := sqlite.New(data.Sqlite3)
	// l, err := query.ListAuthors(ctx)
	// fmt.Printf("debug: %+v\n", l)
	// fmt.Printf("debug: %+v\n", err)

	router.Run("0.0.0.0:8888")
}
