package main

import (
	"douyin-api/controller"
	"douyin-api/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func initRouter() {
	r := gin.New() // without any middleware attached
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// public directory is used to serve static resources
	r.Static("/static", "./public")

	// the login and register does not need token(middleware)
	apiRouterLoginRegister := r.Group("/douyin")
	apiRouterLoginRegister.POST("/user/register/", controller.Register)
	apiRouterLoginRegister.POST("/user/login/", controller.Login)
	apiRouterLoginRegister.GET("/feed/", controller.Feed)

	// register middleware
	apiRouter := r.Group("/douyin").Use(jwt.JWT())

	apiRouter.GET("/user/", controller.UserInfo)
	//apiRouter.POST("/user/register/", controller.Register)
	//apiRouter.POST("/publish/action/", controller.Publish)
	//apiRouter.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	//apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	//apiRouter.GET("/favorite/list/", controller.FavoriteList)
	//apiRouter.POST("/comment/action/", controller.CommentAction)
	//apiRouter.GET("/comment/list/", controller.CommentList)
	//
	// extra apis - II
	//apiRouter.POST("/relation/action/", controller.RelationAction)
	//apiRouter.GET("/relation/follow/list/", controller.FollowList)
	//apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	//apiRouter.GET("/relation/friend/list/", controller.FriendList)
	//apiRouter.GET("/message/chat/", controller.MessageChat)
	//apiRouter.POST("/message/action/", controller.MessageAction)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
