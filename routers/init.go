package routers

import (
	"fmt"
	"net/http"
	"strings"
	"linfeng/controllers"
	"linfeng/apic"
	"linfeng/htmlc"
	_ "linfeng/models"

	"github.com/gin-gonic/gin"
)

func init() {

	router := gin.Default()
	    //加载模版
    //router.LoadHTMLGlob("./view/*")
	router.LoadHTMLGlob("./view/**/*")
	router.StaticFS("/static", http.Dir("./static"))
	// router.GET("/news_list/:page/:limit/:cagegoryid", apic.GetNewslist)
	router.GET("/news_list/:page/:limit/:cagegoryid", htmlc.Newslist)
	router.GET("/index", htmlc.Pcindex)
	router.GET("/ex_list", htmlc.Pczhanhui)
	router.Use(Cors())
	// 版本v1
	admin := router.Group("/admin")
	{
		admin.POST("/login", controllers.LoginController) //登录
		admin.POST("/logout", controllers.Loginout)       //登录
		admin.POST("/add", controllers.AddAdmin)
		admin.POST("/del", controllers.Deladmin)
		admin.POST("/edit", controllers.EditAdmin)
		admin.POST("/getinfo", controllers.GetAdminInfo)
		admin.POST("/getrule", controllers.GetAdminRule)
		admin.POST("/getadminlist", controllers.GetAdminlist)
		//用户组别接口
		admin.POST("/getgrouplist", controllers.Getgrouplist)
		admin.POST("/getgrouptree", controllers.TreeGroup)
		admin.POST("/addgroup", controllers.AddGroup)
		admin.POST("/editgroup", controllers.EditGroup)
		admin.POST("/delgroup", controllers.Delgroup)
		//菜单接口
		admin.POST("/getruleslist", controllers.Getruleslist)
		admin.POST("/delrules", controllers.DelRules)
		admin.POST("/addrules", controllers.AddRules)
		admin.POST("/editrules", controllers.EditRules)
		admin.POST("/getallrule", controllers.GetAllRule)
		//城市接口
		admin.POST("/citylist", controllers.Getcitylist)
		//分类接口
		admin.POST("/categorylist", controllers.Getcategorylist)
		admin.POST("/addcategory", controllers.AddCategory)
		admin.POST("/editcategory", controllers.EditCategory)
		admin.POST("/delcategory", controllers.DelCategory)
		admin.POST("/treecategory", controllers.TreeCategory)
		//爬取接口
		admin.POST("/copy/copycategory", controllers.Copycategory)
		admin.POST("/copy/news", controllers.Copynews)
		admin.POST("/copy/copyex", controllers.Copyex)
		admin.POST("/copy/test", controllers.Maintest)
		//展会接口
		admin.POST("/exhibitionlist", controllers.GetExhibitionlist)
		admin.POST("/addexhibition", controllers.AddExhibition)
		// admin.POST("/editdevice", controllers.EditDevice)
		// admin.POST("/deldevice", controllers.DelDevice)
		// //SIM卡接口
		// admin.POST("/simcardlist", controllers.GetSimcardlist)
		// admin.POST("/addsimcard", controllers.AddSimcard)
		// admin.POST("/editsimcard", controllers.EditSimcard)
		// admin.POST("/delsimcard", controllers.DelSimcard)
		// //开启tcp服务
		// admin.POST("/opentcp", controllers.Opentcp)
		//新闻文章接口
		admin.POST("/newslist", controllers.Getnewslist)
		admin.POST("/addnews", controllers.AddNews)
		admin.POST("/editnews", controllers.EditNews)
		admin.POST("/delnews", controllers.DelNews)

	}
		api := router.Group("/api")
	{
api.POST("/cglist", apic.Getcategorylist) //登录
api.POST("/cgtree", apic.Getcategorytree) //登录
api.POST("/exlist", apic.GetExlist) //登录
api.POST("/citylist", apic.Getcitylist) //登录
api.POST("/newslist", apic.GetNewslist) //登录
	}


	//开启TCP服务啦
	// listen, err := net.Listen("tcp", "127.0.0.1:20000")
	// if err != nil {
	// 	fmt.Println("listen failed, err:", err)
	// 	return
	// }
	// for {
	// 	conn, err := listen.Accept() // 建立连接
	// 	if err != nil {
	// 		fmt.Println("accept failed, err:", err)
	// 		continue
	// 	}
	// 	go process(conn) // 启动一个goroutine处理连接
	// }
	// //开启TCP服务结束
	router.Run(":8088")
}

// //// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json,multipart/form-data,application/x-www-form-urlencoded")                                                                                                        // 设置返回格式是json
		}

		//放行所有OPTIONS方法

		//放行所有OPTIONS方法,防止vue握手2次的问题
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// if method == "OPTIONS" {
		// 	c.JSON(http.StatusOK, "Options Request!")
		// }
		// 处理请求
		c.Next() //  处理请求
	}
}
