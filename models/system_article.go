package models

import (
	"go-admin/public/common"
	"time"
)

type SystemArticle struct {
	Id              int64     `json:"id" xorm:"pk autoincr comment('主键') BIGINT(20)"`
	Author          int       `json:"author" xorm:"not null default 0 comment('作者') INT(10)"`
	Importance      int       `json:"importance" xorm:"not null default 1 comment('重要级别') TINYINT(4)"`
	Status          int       `json:"status" xorm:"not null default 0 comment('状态') TINYINT(4)"`
	Title           string    `json:"title" xorm:"not null default '' comment('标题') VARCHAR(200)"`
	Content         string    `json:"content" xorm:"not null comment('内容') TEXT"`
	ContentShort    string    `json:"content_short" xorm:"not null default '' comment('摘要') VARCHAR(500)"`
	SourceUri       string    `json:"source_uri" xorm:"not null default '' comment('来源') VARCHAR(200)"`
	Ctime           int       `json:"ctime" xorm:"not null default 0 comment('创建时间') INT(11)"`
	ImageUri        string    `json:"image_uri" xorm:"not null default '' comment('图片') VARCHAR(200)"`
	CommentDisabled int       `json:"comment_disabled" xorm:"not null default 0 comment('是否展示评论') TINYINT(4)"`
	Display_time           time.Time `json:"display_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('发布时间') DATETIME"`
	Mtime           time.Time `json:"mtime" xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
}
type SystemArticlePage struct {
	SystemArticle
	AuthorName          string       `json:"authorname"`
}
//var systemarticle = "system_article"
//var SystemStatus = map[string]int{""}
func (m *SystemArticle) GetRow() bool {
	has, err := mEngine.Get(m)
	if err == nil && has {
		return true
	}
	return false
}
func (m *SystemArticle) Add() (int64 ,error){
	return  mEngine.Insert(m)
}
func (m *SystemArticle) AddBatch(beans ...interface{}) (int64 ,error){
	return mEngine.Insert(beans...)
}

func (u *SystemArticle) GetAllPage(paging *common.Paging,filters map[string]string)([]SystemArticle,error) {
	var systemarticles []SystemArticle
	var err error
	session:=mEngine.Where("1=1")
	if filters["status"]!="" {
		session.Where("status=?",filters["status"])
	}
	if filters["importance"]!="" {
		session.Where("importance=?", filters["importance"])
	}
	if filters["title"]!="" {
		session.Where("title like ?","%"+filters["title"]+"%")
	}
	if  filters["start_time"]!="" &&  filters["end_time"]!="" {
		session.Where("mtime>=?",common.StrToTimes(filters["start_time"])).Where("mtime<=?",common.StrToTimes(filters["end_time"]))
	}
	sessionRows:=*session
	paging.Total,err=session.Count(u)
	paging.GetPages()
	if paging.Total<1 {
		return systemarticles,err
	}
	err=sessionRows.Limit(int(paging.PageSize),int(paging.StartNums)).Find(&systemarticles)
	return systemarticles,err
}
func (a *SystemArticle) Update() error {
	if _, err := mEngine.Where("id = ?", a.Id).Update(a); err != nil {
		return err
	}
	return nil
}

