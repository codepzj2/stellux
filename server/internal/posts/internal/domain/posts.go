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
	PageNo int64 `json:"page_no" binding:"required"`
	// 每页数量
	PageSize int64 `json:"page_size" binding:"required"`
	// 排序字段
	Field string `json:"sort_field,omitempty"`
	// 排序规则
	Order string `json:"sort_order,omitempty"`
	// 搜索内容
	Keyword string `json:"keyword,omitempty"`
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
