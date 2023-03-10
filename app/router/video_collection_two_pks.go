// Code generated by gf-codegen. DO NOT EDIT.
// http路由 router
// 生成日期：2022-12-23 18:36:40
// 生成人：Wumengye
package router

import (
	"github.com/WesleyWu/ri-restful-api/app/service"
	"github.com/WesleyWu/ri-restful-api/middleware"
	//"arco-demo/app/api"
    //"github.com/gogf/gf/v2/frame/g"
    //"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

//加载路由
func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Group("/app", func(group *ghttp.RouterGroup) {
			group.Group("/video-collection-2pks", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.ResponseJsonWrapper)
				group.Bind(
					service.VideoCollectionTwoPks,
					service.VideoCollectionTwoPksRepo,
				)
			})
		})
	})
}

