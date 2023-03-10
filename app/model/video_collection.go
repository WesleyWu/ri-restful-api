package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VideoCollection 数据对象
type VideoCollection struct {
	Id          string      `json:"id"`          // 视频集ID，字符串格式
	Name        string      `json:"name"`        // 视频集名称
	ContentType int         `json:"contentType"` // 内容类型
	FilterType  int         `json:"filterType"`  // 筛选类型
	Count       uint32      `json:"count"`       // 集合内视频数量
	IsOnline    bool        `json:"isOnline"`    // 是否上线：0 未上线|1 已上线
	CreatedAt   *gtime.Time `json:"createdAt"`   // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"`   // 更新时间
}

// VideoCollectionInput 用于Insert、Update、Upsert的输入数据对象结构
type VideoCollectionInput struct {
	Id          interface{} `p:"id" v:"required#视频集ID，字符串格式不能为空" json:"id"` // 视频集ID，字符串格式
	Name        interface{} `p:"name" json:"name"`                          // 视频集名称
	ContentType interface{} `p:"contentType" json:"contentType"`            // 内容类型
	FilterType  interface{} `p:"filterType" json:"filterType"`              // 筛选类型
	Count       interface{} `p:"count" json:"count"`                        // 集合内视频数量
	IsOnline    interface{} `p:"isOnline" json:"isOnline"`                  // 是否上线：0 未上线|1 已上线
	CreatedAt   interface{} `p:"createdAt" json:"createdAt"`                // 创建时间
	UpdatedAt   interface{} `p:"updatedAt" json:"updatedAt"`                // 更新时间
}

// VideoCollectionQuery 用于 Query By Example 模式的查询条件数据结构
type VideoCollectionQuery struct {
	Id          interface{} `p:"id" json:"id"`                                         // 视频集ID，字符串格式
	Name        interface{} `p:"name" wildcard:"none" json:"name,omitempty"`           // 视频集名称
	ContentType interface{} `p:"contentType" json:"contentType,omitempty"`             // 内容类型
	FilterType  interface{} `p:"filterType" json:"filterType,omitempty"`               // 筛选类型
	Count       interface{} `p:"count" json:"count,omitempty"`                         // 集合内视频数量
	IsOnline    interface{} `p:"isOnline" json:"isOnline,omitempty"`                   // 是否上线：0 未上线|1 已上线
	CreatedAt   interface{} `p:"createdAt" multi:"between" json:"createdAt,omitempty"` // 创建时间
	UpdatedAt   interface{} `p:"updatedAt" multi:"between" json:"updatedAt,omitempty"` // 更新时间
}

// VideoCollectionCountReq 查询记录总条数的条件数据结构
type VideoCollectionCountReq struct {
	g.Meta `json:"-" path:"/count" method:"get"`
	VideoCollectionQuery
}

// VideoCollectionCountRes 查询记录总条数的返回结果
type VideoCollectionCountRes struct {
	Total int `json:"total"`
}

// VideoCollectionOneReq 查询单一记录的条件数据结构
type VideoCollectionOneReq struct {
	g.Meta `json:"-" path:"/one" method:"get"`
	VideoCollectionQuery
	OrderBy string `json:"orderBy,omitempty"` // 排序方式
}

// VideoCollectionOneRes 查询单一记录的返回结果
type VideoCollectionOneRes struct {
	VideoCollection
}

// VideoCollectionListReq 用于列表查询的查询条件数据结构，支持翻页和排序参数，支持查询条件参数类型自动转换
type VideoCollectionListReq struct {
	g.Meta `json:"-" path:"/list" method:"get"`
	VideoCollectionQuery
	Page     uint32 `d:"1" v:"min:0#分页号码错误" json:"page,omitempty"`          // 当前页码
	PageSize uint32 `d:"10" v:"max:50#分页数量最大50条" json:"pageSize,omitempty"` // 每页记录数
	OrderBy  string `json:"orderBy,omitempty"`                              // 排序方式
}

// VideoCollectionListRes 分页返回结果
type VideoCollectionListRes struct {
	Total   uint64             `json:"total"`   // 记录总数
	Current uint32             `json:"current"` // 当前页码
	Items   []*VideoCollection `json:"items"`   // 当前页记录列表
}

// VideoCollectionCreateReq 插入操作请求参数
type VideoCollectionCreateReq struct {
	g.Meta `orm:"do:true" json:"-" path:"/" method:"post"`
	VideoCollectionInput
}

// VideoCollectionCreateRes 插入操作返回结果
type VideoCollectionCreateRes struct {
	Message      string `json:"message"`      // 提示信息
	LastInsertId int64  `json:"lastInsertId"` // 上一条INSERT插入的记录主键，当主键为自增长时有效
	RowsAffected int64  `json:"rowsAffected"` // 影响的条数
}

// VideoCollectionUpdateReq 更新操作请求参数
type VideoCollectionUpdateReq struct {
	g.Meta `orm:"do:true" json:"-" path:"/:id" method:"patch"`
	VideoCollectionInput
}

// VideoCollectionUpdateRes 更新操作返回结果
type VideoCollectionUpdateRes struct {
	Message      string `json:"message"` // 提示信息
	RowsAffected int64  `json:"rowsAffected"`
}

// VideoCollectionUpsertReq 更新插入操作请求参数
type VideoCollectionUpsertReq struct {
	g.Meta `orm:"do:true" json:"-" path:"/" method:"put"`
	VideoCollectionInput
}

// VideoCollectionUpsertRes 更新插入操作返回结果
type VideoCollectionUpsertRes struct {
	Message      string `json:"message"`      // 提示信息
	LastInsertId int64  `json:"lastInsertId"` // 上一条INSERT插入的记录主键，当主键为自增长时有效
	RowsAffected int64  `json:"rowsAffected"` // 影响的条数
}

// VideoCollectionDeleteReq 删除操作请求参数
type VideoCollectionDeleteReq struct {
	g.Meta `json:"-" path:"/*id" method:"delete"`
	VideoCollectionQuery
}

// VideoCollectionDeleteRes 删除操作返回结果
type VideoCollectionDeleteRes struct {
	Message      string `json:"message"`      // 提示信息
	RowsAffected int64  `json:"rowsAffected"` // 影响的条数
}

// IsEmpty 判断删除请求参数是否为空
func (q *VideoCollectionDeleteReq) IsEmpty() bool {
	return g.IsEmpty(q.Id) &&
		q.Name == nil &&
		q.ContentType == nil &&
		q.FilterType == nil &&
		q.Count == nil &&
		q.IsOnline == nil &&
		q.CreatedAt == nil &&
		q.UpdatedAt == nil
}
