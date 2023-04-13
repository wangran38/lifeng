package apic

import (
	// "fmt"
	"linfeng/models"
	_ "time"
	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

type Cityserch struct {
	Id int64 `json:"id"`
	Shortname  string `json:"shortname"`
	Pid  int `json:"pid"`
    Level  int `json:"level"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Order string `json:"sort"`
}
// type Any interface{}
//获取当前用户信息
func Getcitylist(c *gin.Context) {
	//从header中获取到token
	var searchdata Cityserch
	c.BindJSON(&searchdata)
	// //读取数据库
	
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	// title := searchdata.Title
	order := searchdata.Order
	search := &models.City{
		Id:        searchdata.Id,
		Pid:       searchdata.Pid,
        Level:       searchdata.Level,
		Shortname:     searchdata.Shortname,
	}
	listdata := models.GetCityList(limit, page, search, order)
	listnum := models.GetCitytotal(search)
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



