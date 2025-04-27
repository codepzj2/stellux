package domain

import "time"

// Post 文章
type Post struct {
	ID          string    // 文章ID
	CreatedAt   time.Time // 创建时间
	UpdatedAt   time.Time // 更新时间
	DeletedAt   time.Time // 删除时间
	Title       string    // 标题
	Content     string    // 内容
	Description string    // 描述
	Author      string    // 作者
	Category    string    // 分类
	Tags        []string  // 标签
	IsPublish   bool      // 是否发布
	IsTop       bool      // 是否置顶
	Thumbnail   string    // 缩略图
	LikeCount   int       // 点赞数
	ViewCount   int       // 浏览数
	ShareCount  int       // 分享数
}

// Page 分页
type Page struct {
	PageNo   int64  // 页码
	PageSize int64  // 每页条数
	Keyword  string // 关键词
	Field    string // 排序字段
	Order    string // 排序方向
}
