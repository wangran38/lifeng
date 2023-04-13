package apic

import (
	// "fmt"
	"net/http"
	"linfeng/models"
	_ "time"
	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

type Newsserch struct {
	Id int64 `json:"id"`
	Title  string `json:"title"`
	Categoryid  int `json:"categroy_id"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Order string `json:"sort"`
}

//获取展会信息
func GetNewslist(c *gin.Context) {
	//从header中获取到token
	var searchdata Newsserch
	c.BindJSON(&searchdata)
	// //读取数据库
	
	// name:=""
	limit := searchdata.Limit
	page := searchdata.Page
	// title := searchdata.Title
	order := searchdata.Order
	search := &models.News{
		Id:        searchdata.Id,
		Categoryid:       searchdata.Categoryid,
		Title:     searchdata.Title,
	}
	listdata := models.GetNewsList(limit, page, search, order)
	listnum := models.GetNewstotal(search)
result := make(map[string]interface{})
	result["page"] = page
	result["totalnum"] = listnum
	result["limit"] = limit
	if listdata == nil {
		c.HTML(http.StatusOK, "news_list.html", gin.H{
			"code":    201,
			"message": "获取分类失败",
			"data":    "",
		})
		return
	} else {
		result["listdata"] = listdata
		c.HTML(http.StatusOK, "news_list.html", gin.H{
			"code":    200,
			"message": "数据获取成功",
			"data":    result,
		})
		return
	}
}