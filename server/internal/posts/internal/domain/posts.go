package domain

import "github.com/chenmingyong0423/go-mongox/v2"

type Posts struct {
	mongox.Model `bson:",inline"`
	Title        string
	Content      string
	Author       string
	Category     string
	Tags         []string
	Cover        string
	IsTop        bool
}

type PageDTO struct {
	// 当前页
	PageNo int64 `form:"pageNo" binding:"required"`
	// 每页数量
	PageSize int64 `form:"pageSize" binding:"required"`
	// 排序字段
	Field string `form:"sortField,omitempty"`
	// 排序规则
	Order string `form:"sortOrder,omitempty"`
	// 搜索内容
	Keyword string `form:"keyword,omitempty"`
}

func (p *PageDTO) OrderConvertToInt() int {
	switch p.Order {
	case "ASC":
		return 1
	case "DESC":
		return -1
	default:
		return -1
	}
}
