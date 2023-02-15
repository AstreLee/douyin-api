package main

import (
	"douyin-api/controller"
	"douyin-api/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func initRouter() {
	r := gin.New() // 不包含任何中间件的实例对象
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// public directory is used to serve static resources
	r.Static("/static", "./public")

	// 注册中间件
	apiRouter := r.Group("/douyin").Use(jwt.JWT())

	// basic apis
	//apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	//apiRouter.POST("/publish/action/", controller.Publish)
	//apiRouter.GET("/publish/list/", controller.PublishList)
	//
	//// extra apis - I
	//apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	//apiRouter.GET("/favorite/list/", controller.FavoriteList)
	//apiRouter.POST("/comment/action/", controller.CommentAction)
	//apiRouter.GET("/comment/list/", controller.CommentList)
	//
	//// extra apis - II
	//apiRouter.POST("/relation/action/", controller.RelationAction)
	//apiRouter.GET("/relation/follow/list/", controller.FollowList)
	//apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	//apiRouter.GET("/relation/friend/list/", controller.FriendList)
	//apiRouter.GET("/message/chat/", controller.MessageChat)
	//apiRouter.POST("/message/action/", controller.MessageAction)
}
