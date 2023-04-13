package apic

import (
	// "fmt"
	"linfeng/models"
	_ "time"
	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

type Exserch struct {
	Id int64 `json:"id"`
	Title  string `json:"title"`
	Categoryid  int64 `json:"categroy_id"`
    Cityid  int64 `json:"city_id"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Order string `json:"sort"`
}

//获取展会信息
func GetExlist(c *gin.Context) {
	//从header中获取到token
	var searchdata Exserch
	c.BindJSON(&searchdata)
	// //读取数据库
	
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	// title := searchdata.Title
	order := searchdata.Order
	search := &models.Exhibition{
		Id:        searchdata.Id,
		Categoryid:       searchdata.Categoryid,
        Cityid:       searchdata.Cityid,
		Title:     searchdata.Title,
	}
	listdata := models.GetexhibitionList(limit, page, search, order)
	listnum := models.Getexhibitiontotal(search)
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