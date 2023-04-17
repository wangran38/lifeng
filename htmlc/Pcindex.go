package htmlc

import (
	// "fmt"
	"net/http"
	_ "linfeng/models"
	_ "time"
	_ "strconv"
	// "linfeng/utils"
	"github.com/gin-gonic/gin"
)

// type Newsserch struct {
//     Limit string `uri:"limit" binding:"required"`
//     Page string `uri:"page" binding:"required"`
//     Categoryid string `uri:"categroy_id"`
// 	// Id int64 `json:"id"`
// 	// Title  string `json:"title"`
// 	// Categoryid  int `json:"categroy_id"`
// 	// Limit int    `json:"limit"`
// 	// Page  int    `json:"page"`
// 	// Order string `json:"sort"`
// }

//获取展会信息
func Pcindex(c *gin.Context) {
	//从header中获取到token
	// 	page, ok := gintool.QueryInt64(ctx, "id")
	// if ok {
	// 	fmt.Printf("id: %v\n", id)
	// } else {
	// 	fmt.Printf("\"no\": %v\n", "no")
	// }
//     page := c.Query("page")
// 	pageint, err := strconv.Atoi(page)
// 	if err!=nil {
// 		pageint = 1
// 	}
// 	limit := c.Query("limit")
// 	limitint, err := strconv.Atoi(limit)
// 	if err!=nil {
// 		limitint = 20
// 	}
// 	Categoryid := c.Query("categoryid")
// 	Categoryidint, err := strconv.Atoi(Categoryid)
// 	if err!=nil {
// 		Categoryidint = 0
// 	}
// 	// title := searchdata.Title
// 	order := "id desc"
// 	search := &models.News{
// 		Categoryid:  Categoryidint,
// 	}
// 	listdata := models.GetNewsList(limitint, pageint, search, order)
// 	listnum := models.GetNewstotal(search)
// result := make(map[string]interface{})
// 	result["page"] = page
// 	result["totalnum"] = listnum
// 	result["limit"] = limit
	// if listdata == nil {
	// 	c.HTML(http.StatusOK, "news_list.html", gin.H{
	// 		"code":    201,
	// 		"message": "获取分类失败",
	// 		"data":    "",
	// 	})
	// 	return
	// } else {
	// 	result["listdata"] = listdata
	// 	c.HTML(http.StatusOK, "news_list.html", gin.H{
	// 		"code":    200,
	// 		"message": "数据获取成功",
	// 		"data":    result,
	// 	})
	// 	return
	// }
    		c.HTML(http.StatusOK, "index.html", gin.H{
			"code":    200,
			"message": "数据获取成功",
			"data":    "世展会",
		})
}
func Pczhanhui(c *gin.Context) {
    c.HTML(http.StatusOK, "ex_list.html", gin.H{
			"code":    200,
			"message": "数据获取成功",
			"data":    "世展会",
		})
}