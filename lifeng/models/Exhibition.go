package models

import (
	"errors"
	"time"
)

type Exhibition struct {
	Id      int64     `json:"id"`
	Categoryid     int64       `json:"categroy_id"`
    Cityid     int64       `json:"city_id"`
	Cityname     string    `json:"city_name" xorm:"varchar(200)"`
	Title    string    `json:"title" xorm:"varchar(200)"`
	EnTitle    string    `json:"en_title" xorm:"varchar(200)"`
	Oldlink    string    `json:"oldlink" xorm:"varchar(200)"`
	Image   string    `json:"image" xorm:"TEXT "`
	Keywords   string  `json:"keywords" xorm:"TEXT "`
	Description   string  `json:"description" xorm:" TEXT "`
	Content   string  `json:"content" xorm:"LONGTEXT "` //内容详情
	Isshow     int       `json:"isshow" xorm:"not null default 1 comment('是否启用 默认1 是 0 无') TINYINT"`
    Openyear string `json:"openyear" xorm:"varchar(200)"` //展会年份
    Openbday string `json:"openbday" xorm:"varchar(200)"`  //展会开始日期
	Openeday string `json:"openeday" xorm:"varchar(200)"`  //展会结束日期
    Openbtime string `json:"openbtime" xorm:"varchar(200)"`  //展会的开始时间
    Openetime string `json:"openetime" xorm:"varchar(200)"`  //展会的结束时间
    Openbtimetime int64 `json:"opentime" ` //展会的实际开始日期时间戳
    Openetimetime int64 `json:"opentime" ` //展会的实际结束日期时间戳
    Sponsor string `json:"sponsor" xorm:"varchar(200)"`  //展会的主办方单位
	Period string `json:"period" xorm:"varchar(200)"`  //展会的举办周期
	Periodhistory string `json:"periodhistory" xorm:"varchar(200)"`  //展会的举办历史
    Address string `json:"address" xorm:"varchar(200)"`  //展会的举报地址
	Addressreal string `json:"addressreal" xorm:"varchar(200)"`  //展会的真实地址
    Lng string `json:"lng" xorm:"varchar(200)"`  //地址经度
    Lat string `json:"lat" xorm:"varchar(200)"`  //地址维度
    Area string `json:"area" xorm:"varchar(250)"`  //展会的面积
    Businessmannums float64 `json:"businessmannums" xorm:"Numeric"`  //展商数量
    Opennums float64 `json:"opennums" xorm:"Numeric"`  //观众数量
    Follownums float64 `json:"follownums" xorm:"Numeric"`  //关注数量
	Created time.Time `json:"createtime" xorm:"int"`
	Updated time.Time `json:"updatetime" xorm:"int"`
	Weigh   int  `json:"weigh"`
	Status  int       `json:"status"`
}

func (a *Exhibition) TableName() string {
	return "exhibition"
}

//根据用户id找用户返回数据
func Selectexhibitionid(Id int64) (*Exhibition, error) {
	a := new(Exhibition)
	has, err := orm.Where("id = ?", Id).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("组别菜单数据出错！")
	}
	return a, nil

}

//添加
func Addexhibition(a *Exhibition) error {
	_, err := orm.Insert(a)
	return err
}
//修改
func Upexhibition(a *Exhibition) (int64,error) {
	affected, err := orm.Id(a.Id).Update(a)
	return affected, err

}
func GetexhibitionList(limit int, pagesize int, search *Exhibition, order string) []*Exhibition {
	var page int
	listdata := []*Exhibition{}
	if pagesize-1 < 1 {
		page = 0
	} else {
		page = pagesize - 1
	}
	if limit <= 6 {
		limit = 6

	}
	session := orm.Table("exhibition")
	// stringid := strconv.FormatInt(search.Id, 10)
	if search.Id > 0 {
		session = session.And("id =?", search.Id)
	}
	// fmt.Println(stringid)

	if search.Title != "" {
		title := "%" + search.Title + "%"
		session = session.And("title LIKE ?", title)
		// session = session.And("pid", rules.Title)
	}
	if search.Cityid > 0 {
		session = session.And("cityid = ?", search.Cityid)
	}
	if search.Categoryid > 0 {
		session = session.And("categoryid = ?", search.Categoryid)
	}
	
	var byorder string
	byorder = "id ASC"
	if order != "" {
		byorder = "id DESC"
	}
	session.OrderBy(byorder).Limit(limit, limit*page).Find(&listdata)
	return listdata
}

func Getexhibitiontotal(search *Exhibition) int64 {
	var num int64
	num = 0
	session := orm.Table("exhibition")
	if search.Id > 0 {
		session = session.And("id", search.Id)
	}
	if search.Title != "" {
		name := "%" + search.Title + "%"
		session = session.And("title LIKE ?", name)
		// session = session.And("pid", rules.Title)
	}
	if search.Cityid > 0 {
		session = session.And("cityid = ?", search.Cityid)
	}
	if search.Categoryid > 0 {
		session = session.And("categoryid = ?", search.Categoryid)
	}
	if search.Isshow > 0 {
		session = session.And("isshow = ?", search.Isshow)
		// session = session.And("pid", rules.Title)
	}
	a := new(Exhibition)
	total, err := session.Count(a)
	if err == nil {
		num = total

	}
	return num
}

func Deleteexhibition(id int64) int {
	// intid, _ := strconv.ParseInt(id, 10, 64)
	a := new(Exhibition)
	outnum, _ := orm.ID(id).Delete(a)

	return int(outnum)

}
//根据
func SelectexhibitionByTitle(Title string) (*Exhibition, error) {
	a := new(Exhibition)
	has, err := orm.Where("title = ?", Title).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到！")
	}
	return a, nil

}

//根据
func SelectexhibitionByLink(Link string) (*Exhibition, error) {
	a := new(Exhibition)
	has, err := orm.Where("oldlink = ?", Link).Get(a)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未找到！")
	}
	return a, nil

}