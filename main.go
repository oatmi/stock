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

	v := router.Group("/v")
	{
		v.GET("/login", func(c *gin.Context) {
			cookie, err := c.Cookie("stock_un")
			if err == nil && cookie != "" {
				c.HTML(http.StatusOK, "home.html", nil)
			} else {
				c.HTML(http.StatusOK, "login.html", nil)
			}
		})
		v.GET("/logout", func(c *gin.Context) {
			c.SetCookie("stock_un", "", 86400, "/", c.Request.Host, false, true)
			c.HTML(http.StatusOK, "login.html", nil)
		})

		v.POST("/api/login", handlers.Login)
	}

	// view := router.Group("/view", gin.BasicAuth(account.Accounts))

	view := router.Group("/view", StockAuth())
	{
		view.GET("/home", func(c *gin.Context) { c.HTML(http.StatusOK, "home.html", nil) })
		view.GET("/stock", func(c *gin.Context) { c.HTML(http.StatusOK, "stock.html", nil) })
		view.GET("/in", func(c *gin.Context) { c.HTML(http.StatusOK, "in.html", nil) })
		view.GET("/app", func(c *gin.Context) { c.HTML(http.StatusOK, "application.html", nil) })
		view.GET("/out", func(c *gin.Context) { c.HTML(http.StatusOK, "out.html", nil) })
		view.GET("/out_approve", func(c *gin.Context) { c.HTML(http.StatusOK, "out_approve.html", nil) })
		view.GET("/price", func(c *gin.Context) { c.HTML(http.StatusOK, "price.html", nil) })
	}

	api := router.Group("/api", StockAuth())
	{
		api.GET("/home", handlers.GetStocks)
		api.GET("/instock", handlers.GetApplications)
		api.GET("/outstock", handlers.OutStockList)
		api.GET("/navs", handlers.Navs)

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

func StockAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("stock_un")
		if err != nil || cookie == "" {
			ctx.Redirect(http.StatusTemporaryRedirect, "/v/login")
			ctx.AbortWithStatus(http.StatusTemporaryRedirect)
		}
	}
}
