package common

import (
	"math"
)
// 分页器
//  paging:=Paging{}
//  paging.GetPages(page, pagesize, total)
//@author Zhiqing Guo
type Paging struct {
	Page      int64  `json:"page" form:"page"` //当前页
	PageSize  int64    `json:"pagesize" form:"pagesize"` //每页条数
	Total     int64   `json:"total" form:"total"` //总条数
	PageCount int64   `json:"pagecount" form:"pagecount"` //总页数
	StartNums int64   `json:"startnums" form:"startnums"` //起始条数
	//Nums      []int64  `json:"nums" form:"Nums"`//分页序数
	//NumsCount int64   `json:"numscount" form:"numscount"` //总页序数
}
//获取分页信息
func( p *Paging)GetPages() *Paging {

	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 10
	}

	page_count := math.Ceil(float64(p.Total) / float64(p.PageSize))

	paging := new(Paging)
	paging.Page=p.Page
	paging.PageSize=p.PageSize
	paging.StartNums = p.PageSize*(p.Page-1)
	paging.Total = p.Total
	paging.PageCount = int64(page_count)
	return paging
}

