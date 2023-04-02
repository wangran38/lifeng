package controllers

import (
	"time"
	"strconv"
	"strings"
	"fmt"
	"log"
	"linfeng/models"
    "io/ioutil"
	"net/http"
	"regexp"
	// "ginstudy/utils"
	"github.com/gin-gonic/gin"
)
const (
    DIV_CATEGORY = `<div class="news_list_cate_list_title">(?s:(.*?))</div>` //栏目的UL
    // HOME_TITLE_LI = `.*?<a.*?href=".*?".*?>(.+?)</a>.*?` //栏目的LI
    // LUNWENDIV = `.*?<dd>(?s:(.*?))</dd>.*?` //
    // LUNWENA = `.*?<a.*?href=".*?".*?class="nav_left">(.+?)</a>.*?`
    // LUNWENAPAGEUL = `.*?<ul class="lwbox">(?s:(.*?))</ul>.*?` //
    // LUNWENAPAGEULH4 = `.*?<h4><a.*?href="(?s:(.*?))".*?>.*?</a></h4>.*?` //page
    // LUNWENILW = `.*?<div class="info ilw">(?s:(.*?))</div>.*?`
    // INFO = `.*?<div class="info">(?s:(.*?))</div>.*?`
    H2 = `<h2>(?s:(.*?))</h2>` //page
    // LWUL = `.*?<ul class="lwbox">(?s:(.*?))</ul>.*?` //page
    // LWP = `.*?<p class="tsign">(?s:(.*?))</p>.*?` //page
    TITLEA = `<a.*?href=".*?".*?>(.+?)</a>.*?` //a
    // IMG = `.*?<img.*?src="(?s:(.*?))".*?/>.*?` //a
    // TD = `.*?<td.*?td="" width="250">(?s:(.*?))</td>.*?`
    // RE_EMAIL = `(\w+?)(\.\w+)?@(\w+)?\.(\w{2,5})(\.\w{2,3})?`
    // RE_LINK  = `<a[\s\S]+?href="(http[\s\S]+?)"`
	//展会详情页的解析
	 DIV_EX = `<div class="top_content" .*?>(?s:(.*?))</div>` //栏目的UL
	 DHA = `<a.*?>(?s:(.*?))</a>` //a
	 SPANT = `<span .*?>(?s:(.*?))</span>` //获取栏目的标题
	 EXH_DETAIL = `<div class="exh_detail_right" .*?>(?s:(.*?))</div>` //栏目的UL
	 L_VALUE = `<span class="line_value" .*?>(?s:(.*?))</span>` //栏目的UL
	 SPAN_SPAN = `<span style=".*?" .*?><span .*?>.*?</span>(?s:(.*?))</span>` //栏目的UL
	 SPAN_LINE = `<span style="margin-right: 20px" .*?>(?s:(.*?))</span>` //栏目的UL
	//   <div style="margin-top: 12px" data-v-98726318>
	//列表页采集
	EXH_LIST = `.*?<div class="goods-item-container goods_exh" .*?>(?s:(.*?))</div></div></div></div>.*?` //栏目的UL
	NAME_REMIND = `<a href=".*?" title=".*?" target="_blank" .*?>(?s:(.*?))</a>` //栏目的UL
	ALINK = `.*?<a href="(.*?)" title=".*?" target="_blank" .*?>.*?</a>.*?` //a
	EN_NAME = `<div class="En_name" .*?>(?s:(.*?))</div>` //a
	SCALE = `<div class="scale-remind" .*?><div .*?>(?s:(.*?))</div>` //取得会展面积
	ETIME = `<time .*?>(?s:(.*?))</time>` //取得会展的时间
      
	 

)
func Copycategory(c *gin.Context) {
    resp, _ := http.Get("https://www.shifair.com/information/2-0/")
    defer resp.Body.Close() //go的特殊语法，main函数执行结束前会执行 resp.Body.Close()
    //fmt.Println(resp.StatusCode)          //有http的响应码输出
    if resp.StatusCode == http.StatusOK { //如果响应码为200
        body, err := ioutil.ReadAll(resp.Body) //把响应的body读出
        if err != nil {                        //如果有异常
            fmt.Println(err) //把异常打印
            log.Fatal(err)   //日志
        }
    //   fmt.Println(string(body)) //把响应的文本输出到console
	//   fmt.Println("=======================================================")
		    //解释正则表达式提取UL
    uldata := regexp.MustCompile(DIV_CATEGORY)
	        //提取关键信息
    result := uldata.FindAllStringSubmatch(string(body),-1)
    for _, v1 := range result {

	//fmt.Println(v1[0])
	// 	fmt.Println(result[0][0])
	lidata := regexp.MustCompile(H2)
	result1 := lidata.FindAllStringSubmatch(string(v1[0]),-1)
	for _, v2 := range result1 {
categorydata := new(models.Category)
	categorydata.Title = v2[1]
	categorydata.Keywords = v2[1]
	categorydata.Description = v2[1]
	categorydata.Content = v2[1]
	categorydata.Created = time.Now()
	info, _ := models.SelectcategoryByTitle(v2[1]) //判断大类是否存在！
	if info == nil {
err := models.Addcategory(categorydata) //判断账号是否存在！
		if err != nil {
			fmt.Println(err)
	}
	lidata1 := regexp.MustCompile(TITLEA)
	result2 := lidata1.FindAllStringSubmatch(string(v1[0]),-1)
	for _, v3 := range result2 {
		categorydata1 := new(models.Category)
		categorydata1.Pid = categorydata.Id
	categorydata1.Title = v3[1]
	categorydata1.Keywords = v3[1]
	categorydata1.Description = v3[1]
	categorydata1.Content = v3[1]
	categorydata1.Created = time.Now()
	err := models.Addcategory(categorydata1) //判断账号是否存在！
		if err != nil {
			fmt.Println(err)
	}
// fmt.Println(v3[1])
	}
	}

//fmt.Println(id)

// fmt.Println(v2[1])//大类的名称
	}

    }


	}
}

//获取当前用户 所有的权限控制
/**
  以GET的方式请求
  **/
func Copyexhibition(url string) string {
	upurl:= url
	datainfo, _ := models.SelectexhibitionByLink(upurl) //判断账号是否存在！
	if datainfo == nil {
		return "为空"
	}
	resp, _ := http.Get(upurl)
    defer resp.Body.Close() //go的特殊语法，main函数执行结束前会执行 resp.Body.Close()
    //fmt.Println(resp.StatusCode)          //有http的响应码输出
    if resp.StatusCode == http.StatusOK { //如果响应码为200
        body, err := ioutil.ReadAll(resp.Body) //把响应的body读出
        if err != nil {                        //如果有异常
            fmt.Println(err) //把异常打印
            log.Fatal(err)   //日志
        }
		
    //  fmt.Println(string(body)) //把响应的文本输出到console
	//    fmt.Println("==========================================================================================")
	uldata := regexp.MustCompile(DIV_EX)
	//提取关键信息
    result := uldata.FindAllStringSubmatch(string(body),-1)
	//fmt.Println(string(result[0][1]))
	lidata := regexp.MustCompile(DHA)
	result1 := lidata.FindAllStringSubmatch(string(result[0][1]),-1)
	
	// s1:=len([]rune(string(result1[2][1])))
	// snum := s1-2
	// fmt.Println(int(snum))
	// str := string([]rune(result1[2][1])[0:2])
	s := strings.TrimSpace(result1[2][1]) //去除首尾空格，靠
	st := "展会"
	str := strings.TrimSuffix(s, st) //去除展会

	//fmt.Println(string(str))
	var categpryid int64 //
		info, _ := models.SelectcategoryByTitle(str) //判断大类是否存在！
	if info == nil {
		categpryid = 344 //如果没找到该分类
	} else {
		categpryid = info.Id //如果找到该分类
	}
	//fmt.Println(int64(id))
	// categpryid, _ := strconv.ParseInt(id, 64, 10)
	exhibitiondata := new(models.Exhibition)
	exhibitiondata.Categoryid=categpryid

	//获取详情的DIV
	divdata1 := regexp.MustCompile(EXH_DETAIL)
	rs1 := divdata1.FindAllStringSubmatch(string(body),-1)
	//fmt.Println(string(rs1[0][1])) 
	divdata2 := regexp.MustCompile(SPANT)
	rs2 := divdata2.FindAllStringSubmatch(string(rs1[0][1]),-1)
	oldyear:= strings.TrimSpace(rs2[0][1])
	//fmt.Println(strings.TrimSpace(rs2[1][1])) //开馆时间
	opentimefg := strings.Split(strings.TrimSpace(rs2[1][1]),"-") //分割-
	// fmt.Println(string(opentimefg[0]))
	// fmt.Println(string(opentimefg[1]))
	opentimefg1:= strings.Split(opentimefg[0],":") //分割.
	// fmt.Println(string(opentimefg1[0]))
	// fmt.Println(string(opentimefg1[1]))
	// fmt.Println(string(opentimefg1[2]))
	fmt.Println("开馆准确时间--"+ string(opentimefg1[1]+":"+ opentimefg1[2])) //开馆开始时间
	opentimefg2:= strings.Split(opentimefg[1],":") //分割.
	// fmt.Println(string(opentimefg2[0]))
	// fmt.Println(string(opentimefg2[1]))
	fmt.Println("关馆准确时间--"+string(opentimefg2[0]+":"+ opentimefg2[1])) //开馆结束时间
	//日期字符分割再合并
	fgyear:= strings.Split(oldyear,"-") //分割-
	fenge1:= strings.Split(fgyear[0],".") //分割.
	fenge2:= strings.Split(fgyear[1],".") //分割.
	fmt.Println(string(fenge1[0]+"/"+fenge1[1]+"/"+fenge1[2])) //开始时间
	fmt.Println(string(fenge1[0]+"/"+fenge2[0]+"/"+fenge2[1])) //结束时间
	exhibitiondata.Openyear=fenge1[0] //开始年份
	exhibitiondata.Openbday=fenge1[1]+"/"+fenge1[2] //开始日期
	exhibitiondata.Openeday=fenge2[0]+"/"+fenge2[1]//结束日期
	exhibitiondata.Openbtime=opentimefg1[1]+":"+ opentimefg1[2]//开始日期
	exhibitiondata.Openetime=opentimefg2[0]+":"+ opentimefg2[1]//结束日期
	//时间戳开始
	//nowStr := time.Now().Format("2006/01/02 15:04:05") //根据指定的模板[ 2006/01/02 15:04:05 ]，返回时间。
	// nowUnix, err := strToUnix(nowStr, "2006/01/02 15:04:05")
	// nowStr = time.Now().Format(exhibitiondata.Openetime+ "00:00:01") //根据指定的模板[ 2006/01/02 15:04:05 ]，返回时间。
	btimetime := TimeToUnix(fenge1[0]+"/"+fenge1[1]+"/"+fenge1[2]+ " 00:00:01")
	etimetime := TimeToUnix(fenge1[0]+"/"+fenge2[0]+"/"+fenge2[1]+ " 23:59:59")
	fmt.Println(string(fenge1[0]+"/"+fenge1[1]+"/"+fenge1[2]+ " 23:59:59"))
	fmt.Println(int64(btimetime))
	fmt.Println(int64(etimetime))
	timeStr := fenge1[0]+"/"+fenge1[1]+"/"+fenge1[2]
    fmt.Println("timeStr:", timeStr)
    // t, _ := time.Parse("2006/01/02", timeStr)
    // fmt.Println(t.Format(time.UnixDate))
	exhibitiondata.Openbtimetime=btimetime
	exhibitiondata.Openetimetime=etimetime
	//时间戳结束
	//日期字符结束
	//获取主办方单位
		lt1 := regexp.MustCompile(L_VALUE)
	rss2 := lt1.FindAllStringSubmatch(string(body),-1)
	//fmt.Println(string(rss2[0][1])) 
	fmt.Println("主办方单位"+string(rss2[1][1])) //主办单位
    exhibitiondata.Sponsor=strings.TrimSpace(rss2[1][1]) //主办方单位
	fmt.Println("展会地址"+string(rss2[2][1])) //展会地址
	// fmt.Println("举办历史"+string(rss2[3][1])) //举办历史
	// if rss2[3][1]==
	// exhibitiondata.Periodhistory=strings.TrimSpace(rss2[3][1]) //举办历史

	lt2 := regexp.MustCompile(DHA)
	rss22 := lt2.FindAllStringSubmatch(string(rss2[2][1]),-1)
	fmt.Println("举办地址"+string(rss22[0][1])) //主办方地址
	exhibitiondata.Address=strings.TrimSpace(rss22[0][1]) //举办地址
	//
	//获取主办方单位结束
	//获取举办周期，展商数量
	spandata := regexp.MustCompile(SPAN_SPAN)
	spanspanrs := spandata.FindAllStringSubmatch(string(body),1)
fmt.Println("举办周期"+ string(spanspanrs[0][1])) //举办周期
    exhibitiondata.Period=strings.TrimSpace(spanspanrs[0][1]) //举办历史周期
//进行更新操作
exhibitiondata.Id=datainfo.Id
 infott, err := models.Upexhibition(exhibitiondata) //判断账号是否存在！
if err!=nil {
	fmt.Println(err)
} else {
	fmt.Println(int64(infott))
}
return ("成功执行")
///////////////////////////
	}
	return ("在https状态外部")

}
func Copyex(c *gin.Context) {
	for i := 1; i <= 200; i++ {
	resp, _ := http.Get("https://www.jufair.com/exhibition-0-0-1-0-0-0-"+ strconv.Itoa(i) +"/")
    defer resp.Body.Close() //go的特殊语法，main函数执行结束前会执行 resp.Body.Close()
    //fmt.Println(resp.StatusCode)          //有http的响应码输出
    if resp.StatusCode == http.StatusOK { //如果响应码为200
        body, err := ioutil.ReadAll(resp.Body) //把响应的body读出
        if err != nil {                        //如果有异常
            fmt.Println(err) //把异常打印
            log.Fatal(err)   //日志
        }
		
    //fmt.Println(string(body)) //把响应的文本输出到console
	listdata := regexp.MustCompile(EXH_LIST)
	        //提取关键信息
    result := listdata.FindAllStringSubmatch(string(body),-1)
	//fmt.Println(string(result[0][1]))
	for _, v1 := range result { //循环获取内容
	//    rq := httplib.Get(v1[1])
	// fmt.Println("======================================================================")
    strr:= string(v1[0])

	//提取展会标题信息
	titlename := regexp.MustCompile(NAME_REMIND)
    title_rs := titlename.FindAllStringSubmatch(strr,-1)

	// fmt.Println(string(title_rs[1][1]))
	// fmt.Println("---------------------------------------------------")
	exhibitiondata := new(models.Exhibition)
	exhibitiondata.Title=strings.TrimSpace(title_rs[1][1])
	//提取展会标题结束
	//提取展会的英文标题
	title_en := regexp.MustCompile(EN_NAME)
    title_en_rs := title_en.FindAllStringSubmatch(strr,-1)
	// fmt.Println(string(title_en_rs[0][1]))
	// fmt.Println("======================================================================")
	exhibitiondata.EnTitle=title_en_rs[0][1] //英文
    //提取展会的英文标题结束
	//提取链接URL
	a_link := regexp.MustCompile(ALINK)
    a_link_rs := a_link.FindAllStringSubmatch(strr,-1)
	exhibitiondata.Oldlink="https://www.jufair.com"+ string(a_link_rs[0][1]) //英文
	//fmt.Println("https://www.jufair.com"+ string(a_link_rs[0][1]))
	//提取链接URL结束
	//提取展会时间
	time_name := regexp.MustCompile(ETIME)
    time_rs := time_name.FindAllStringSubmatch(strr,-1)
	fmt.Println(string(time_rs[0][1]))
	//提取展会时间结束
	//提取展会面积
	scale_data := regexp.MustCompile(SCALE)
	scale_rs := scale_data.FindAllStringSubmatch(strr,-1)
	exhibitiondata.Area=strings.TrimSpace(scale_rs[0][1]) //展会面积
		err := models.Addexhibition(exhibitiondata) //判断账号是否存在！
		if err != nil {
			fmt.Println(err)
	}
	ch := make(chan string)
    // 开启一个并发匿名函数
    go func() {
    tturl := exhibitiondata.Oldlink
    // staticurl := "static/upload/wximg/"
    gaourl := Copyexhibition(tturl) //返回图片地址
        // 通过通道通知main的goroutine
        ch <- gaourl
        // fmt.Println("exit goroutine")
    }()
	//fmt.Println(string(scale_rs[0][1]))
	//SCALE
	//提取展会面积结束
		}

	}
time.Sleep(20 * time.Second) 
	}
}
func Maintest(c *gin.Context) {
	//1、时间戳转时间
	nowUnix := time.Now().Unix() //获取当前时间戳
	nowStr := unixToStr(nowUnix, "2006-01-02 15:04:05")
	fmt.Printf("1、时间戳转时间：%d => %s \n", nowUnix, nowStr)

	//2、时间转时间戳
	nowStr = time.Now().Format("2006/01/02 15:04:05") //根据指定的模板[ 2006/01/02 15:04:05 ]，返回时间。
	nowUnix, err := strToUnix(nowStr, "2006/01/02 15:04:05")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("2、时间转时间戳：%s => %d", nowStr, nowUnix)
}

//时间戳转时间
func unixToStr(timeUnix int64, layout string) string {
	timeStr := time.Unix(timeUnix, 0).Format(layout)
	return timeStr
}
// 时间（字符串）转时间戳
func TimeToUnix(str string) int64 {
	// go语言固定日期模版
	timeLayout := "2006/01/02 15:04:05"
	times, _ := time.Parse(timeLayout, str)
	return times.Unix()
}

//时间转时间戳
func strToUnix(timeStr, layout string) (int64, error) {
	local, err := time.LoadLocation("Asia/Shanghai") //设置时区
	if err != nil {
		return 0, err
	}
	tt, err := time.ParseInLocation(layout, timeStr, local)
	if err != nil {
		return 0, err
	}
	timeUnix := tt.Unix()
	return timeUnix, nil
}
