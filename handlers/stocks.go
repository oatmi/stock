package handlers

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oatmi/stock/data"
	"github.com/oatmi/stock/data/sqlite"
	"github.com/oatmi/stock/lib"
	"github.com/spf13/cast"
)

var (
	funcProductType = func(t int32) string {
		switch t {
		case 1:
			return "主材"
		case 2:
			return "辅材"
		case 3:
			return "半成品"
		case 4:
			return "成品"
		case 5:
			return "鸡货"
		case 6:
			return "报废品"
		default:
			return "其他"
		}
	}

	funcProductAttr = func(t int32) string {
		switch t {
		case 1:
			return "医疗器械"
		case 2:
			return "非医疗器械"
		default:
			return "其他"
		}
	}

	funcWayIn = func(t int32) string {
		switch t {
		case 1:
			return "客供"
		case 2:
			return "期初"
		case 3:
			return "外发"
		case 4:
			return "外购"
		default:
			return "其他"
		}
	}
)

type AisudaiResponse struct {
	Status  int         `json:"status"` // 状态码，为0表示成功
	Message string      `json:"msg"`    // 文案信息
	Data    interface{} `json:"data"`   // 返回值内容
}

type AisudaiCRUDData struct {
	Count    int         `json:"count"`
	Download int         `json:"download"`
	Rows     interface{} `json:"rows"`
}

type StockItem struct {
	sqlite.Stock
}

func ExportStocks(c *gin.Context) {
	query := sqlite.New(data.Sqlite3)

	listParam := buildListStockParams(c)
	list, err := query.ListStocks(c, listParam)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Message: "无数据 " + err.Error()})
		return
	}
	if len(list) == 0 {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "无数据"})
		return
	}

	fileName := "/tmp/stock.csv"
	file, err := os.Create(fileName)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "无数据"})
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	// 写入CSV文件的标题行
	headers := []string{
		"货品名称",
		"货品类型",
		"货品属性",
		"供应商",
		"规格型号",
		"单位",
		"入库批号",
		"入库方式",
		"存放仓库",
		"生产批号",
		"生产日期",
		"灭菌批号",
		"灭菌日期",
		"入库时间",
		"入库数量",
		"货品单价",
		"当前价值",
	}
	if err := writer.Write(headers); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	// 将数据写入CSV文件
	for _, stock := range list {
		record := []string{
			stock.Name,
			funcProductType(stock.ProductType),
			funcProductAttr(stock.ProductAttr),
			stock.Supplier,
			stock.Model,
			stock.Unit,
			stock.BatchNoIn,
			funcWayIn(stock.WayIn),
			cast.ToString(stock.Location),
			stock.BatchNoProduce,
			lib.TimestampToDate(int64(stock.ProduceDate)),
			stock.DisinfectionNo,
			lib.TimestampToDate(int64(stock.DisinfectionDate)),
			lib.TimestampToDate(int64(stock.StockDate)),
			cast.ToString(stock.StockNum),
			cast.ToString(stock.CurrentNum),
			cast.ToString(stock.Price),
			cast.ToString(stock.Value),
		}

		if err := writer.Write(record); err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
	}

	writer.Flush()

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=stock.csv")
	c.File(fileName)
}

// GetStocks 获取库存数据
func GetStocks(c *gin.Context) {
	query := sqlite.New(data.Sqlite3)

	listParam := buildListStockParams(c)
	list, err := query.ListStocks(c, listParam)
	if err != nil {
		c.JSON(http.StatusOK, AisudaiResponse{Message: "无数据 " + err.Error()})
		return
	}
	if len(list) == 0 {
		c.JSON(http.StatusOK, AisudaiResponse{Status: 1, Message: "无数据"})
		return
	}

	var countParam sqlite.CountStocksParams
	listParamByte, _ := json.Marshal(listParam)
	json.Unmarshal(listParamByte, &countParam)

	count, err := query.CountStocks(c, countParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	resp := []StockItem{}
	for _, stock := range list {
		stockItem := StockItem{
			Stock: stock,
		}

		if !strings.Contains("chengnanyu,wangbo,likaihou,chenyu", lib.UserName(c)) {
			if stock.ProductType == 4 && lib.UserName(c) == "zhangling" {
				// Do Noting ...
			} else {
				stockItem.Value = 0
				stockItem.Price = 0
			}
		}

		resp = append(resp, stockItem)
	}
	download := 0
	if strings.Contains("likaihou", lib.UserName(c)) {
		download = 1
	}

	c.JSON(http.StatusOK, AisudaiCRUDData{Count: int(count), Download: download, Rows: resp})
}

func buildListStockParams(c *gin.Context) sqlite.ListStocksParams {
	arg := sqlite.ListStocksParams{
		Limit: sql.NullInt32{
			Int32: 10,
			Valid: true,
		},
		Offset: sql.NullInt32{
			Int32: (cast.ToInt32(c.DefaultQuery("page", "0")) - 1) * 10,
			Valid: true,
		},
	}

	if val, ok := c.GetQuery("name"); ok && val != "" {
		arg.Name = sql.NullString{
			String: "%" + val + "%",
			Valid:  true,
		}
	}
	if val, ok := c.GetQuery("product_type"); ok && val != "" {
		arg.ProductType = sql.NullInt32{
			Int32: cast.ToInt32(val),
			Valid: true,
		}
	}
	if val, ok := c.GetQuery("product_attr"); ok && val != "" {
		arg.ProductAttr = sql.NullInt32{
			Int32: cast.ToInt32(val),
			Valid: true,
		}
	}
	if val, ok := c.GetQuery("status"); ok && val != "" {
		arg.Status = sql.NullInt32{
			Int32: cast.ToInt32(val),
			Valid: true,
		}
	} else {
		arg.Status = sql.NullInt32{
			Int32: 1,
			Valid: true,
		}
	}

	return arg
}
