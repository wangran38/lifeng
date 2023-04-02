package apic

import (
	// "fmt"
	"linfeng/models"
	_ "time"
	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

type CategoryJserch struct {
	Id int64 `json:"id"`
	Title  string `json:"title"`
	Pid  int64 `json:"pid"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Order string `json:"sort"`
}
// type Any interface{}
//获取当前用户信息
func Getcategorylist(c *gin.Context) {
	//从header中获取到token
	var searchdata CategoryJserch
	c.BindJSON(&searchdata)
	// //读取数据库
	
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	// title := searchdata.Title
	order := searchdata.Order
	search := &models.Category{
		Id:        searchdata.Id,
		Pid:       searchdata.Pid,
		Title:     searchdata.Title,
	}
	listdata := models.GetCList(limit, page, search, order)
	listnum := models.GetCNum(search)
result := make(map[string]interface{})
	result["page"] = page
	result["totalnum"] = listnum
	result["limit"] = limit
	if listdata == nil {
		c.JSON(200, gin.H{
			"code":    201,
			"message": "获取分类失败",
			"data":    "",
		})
		return
	} else {
		result["listdata"] = listdata
		c.JSON(200, gin.H{
			"code":    200,
			"message": "数据获取成功",
			"data":    result,
		})
		return
	}
}
//获取全部上下级



