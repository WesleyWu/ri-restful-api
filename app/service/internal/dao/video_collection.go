// Code generated by gf-codegen. DO NOT EDIT.
// 数据操作 dao 包装类
// 生成日期：2022-12-23 18:36:40
// 生成人：Wumengye
package dao

import (
	"github.com/WesleyWu/ri-restful-api/app/service/internal/dao/internal"
)

// videoCollectionDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type videoCollectionDao struct {
	*internal.VideoCollectionDao
}

var (
    // VideoCollection is globally public accessible object for table tools_gen_table operations.
    VideoCollection = videoCollectionDao{
        internal.NewVideoCollectionDao(),
    }
	VideoCollectionDbType = VideoCollection.DB().GetConfig().Type
)