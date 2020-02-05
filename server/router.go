package server

import (
	"borderland/api"
	"borderland/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		/**
		* showdoc
		* @catalog 测试文档/运行测试
		* @title 服务器运行
		* @description 测试服务器是否运行
		* @method post
		* @url /api/v1/ping
		* @return {"code":0,"msg":"Pong"}
		 */
		v1.POST("ping", api.Ping)

		/**
		* showdoc
		* @catalog 用户相关
		* @title 用户注册
		* @description 用户注册api
		* @method post
		* @url /api/v1/user/register
		* @param nickname 必选 string 昵称
		* @param user_name 必选 string 用户名
		* @param password 必选 string 密码
		* @param password_confirm 必选 string 再次确认密码
		* @return {"code":0,"data":{"id":1,"user_name":"explorer","nickname":"lovelyjax","status":"active","avatar":"","created_at":1580879384},"msg":""}
		 */
		v1.POST("user/register", api.UserRegister)

		/**
		* showdoc
		* @catalog 用户相关
		* @title 用户登录
		* @description 用户登录api
		* @method post
		* @url /api/v1/user/login
		* @param user_name 必选 string 用户名
		* @param password 必选 string 密码
		* @return  {"code":0,"data":{"id":1,"user_name":"explorer","nickname":"lovelyjax","status":"active","avatar":"","created_at":1580879384},"msg":""}
		 */
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			/**
			* showdoc
			* @catalog 用户相关
			* @title 获取当前用户信息
			* @description 获取当前用户信息api
			* @method get
			* @url /api/v1/user/me
			* @return  {"code":0,"data":{"id":1,"user_name":"explorer","nickname":"lovelyjax","status":"active","avatar":"","created_at":1580879384},"msg":""}
			 */
			auth.GET("user/me", api.UserMe)

			/**
			* showdoc
			* @catalog 用户相关
			* @title 注销登录
			* @description 注销登录api
			* @method delete
			* @url /api/v1/user/logout
			* @return {"code":0,"msg":"登出成功"}
			 */
			auth.DELETE("user/logout", api.UserLogout)
		}
	}
	return r
}
