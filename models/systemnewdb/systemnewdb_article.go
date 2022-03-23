package systemnewdb

import "go-admin/lib/common"

type SystemArticlePage struct {
	SystemArticle
	AuthorName string `json:"authorname"`
}

//根据非零值获取一条数据
func (a *SystemArticle) GetRow() (int64, error) {
	result := db.First(&a)
	return result.RowsAffected, result.Error
}
func (a *SystemArticle) Update() error {
	result := db.Model(&a).Where("id=?", a.ID).Updates(a)
	return result.Error
}

func (a *SystemArticle) GetAllPage(paging *common.Paging, filters map[string]string) ([]SystemArticle, error) {
	var systemarticles []SystemArticle
	var err error
	tx := db.Where("1=1")
	if filters["status"] != "" {
		tx.Where("state=?", filters["state"])
	}
	if filters["importance"] != "" {
		tx.Where("importance=?", filters["importance"])
	}
	if filters["title"] != "" {
		tx.Where("title like ?", "%"+filters["title"]+"%")
	}
	if filters["start_time"] != "" && filters["end_time"] != "" {
		tx.Where("mtime>=?", common.StrToTimes(filters["start_time"])).Where("mtime<=?", common.StrToTimes(filters["end_time"]))
	}
	sessionRows := *tx
	var count int64
	tx.Model(&a).Count(&count)
	paging.Total = count
	paging.GetPages()
	if paging.Total < 1 {
		return systemarticles, err
	}
	err = sessionRows.Limit(int(paging.PageSize)).Offset(int(paging.StartNums)).Find(&systemarticles).Error
	return systemarticles, err
}
