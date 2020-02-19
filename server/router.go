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

		/**
		* showdoc
		* @catalog 用户相关
		* @title 站长信息
		* @description 获取站长信息的api
		* @method get
		* @url /api/v1/user/owner
		* @return  {"code":0,"data":{"name":"陈亮","profession":"在读三年级本科生","school":"华南理工大学","address":"广东广州","email":"501892778@qq.com","hobby":"任何有趣的事"},"msg":""}
		 */
		v1.GET("user/owner", api.StaticOwner)

		/**
		* showdoc
		* @catalog 云存储相关
		* @title 获取阿里云存储签名
		* @description 获取阿里云存储签名的api。
		* @method post
		* @url /api/v1/upload/token
		* @return {"code":0,"data":{"filename":"avatar/06d62bcc-fa8c-44f8-9f18-5f7716927350.jpg","getapi":"http://borderland.oss-cn-shenzhen.aliyuncs.com/avatar%2F06d62bcc-fa8c-44f8-9f18-5f7716927350.jpg?Expires=1581143926\u0026OSSAccessKeyId=LTAI4FodY2q3iXWCJKXHG17u\u0026Signature=f851azSCvkhl1AD50gBsLVrl9ts%3D","putapi":"http://borderland.oss-cn-shenzhen.aliyuncs.com/avatar%2F06d62bcc-fa8c-44f8-9f18-5f7716927350.jpg?Expires=1581143926\u0026OSSAccessKeyId=LTAI4FodY2q3iXWCJKXHG17u\u0026Signature=BZsx4NbU4AhDjmsPuOcdinyJDGc%3D"},"msg":""}
		* @return_param filename string 将要上传的文件的文件名
		* @return_param putapi string 真正进行上传文件的url地址
		* @return_param getapi string 上传文件后，访问文件的api地址
		* @remark 有必要阐述一下云存储流程。本api实际上只是进行签名的，访问本api不需要带任何访问参数，只是告诉阿里云，我将要上传什么信息了。拿到putapi后，需要使用put请求，header需要指定Content-Type为image/jpg，body里面放应该上传的二进制文件，成功的话不会返回任何信息。然后可以通过getapi访问上传后的文件。 两个api签名都只有十分钟的时效，过了之后需要重新申请签名。

		 */
		v1.POST("upload/token", api.UploadToken)

		/**
		* showdoc
		* @catalog 文章相关
		* @title 添加文章
		* @description 添加文章的api。
		* @method post
		* @url /api/v1/article
		* @param title 必选 string 文章标题
		* @param catalog 必选 string 文章列别
		* @return {"code":0,"data":{"id":1,"title":"自序","content":"","catalog":"建站日记","created_at":1581488171},"msg":""}
		 */
		v1.POST("article", api.ArticlePost)

		/**
		* showdoc
		* @catalog 文章相关
		* @title 具体内容文章
		* @description 查看文章具体内容的api。
		* @method post
		* @url /api/v1/article/id
		* @param id 必选 int 主键
		* @return {"code":0,"data":{"id":1,"title":"自序","content":"# 为什么要搭博客","catalog":"建站日记","view": 4,"created_at":"2020年2月12日 14:16:11"},"msg":""}
		 */
		v1.GET("article/:id", api.ShowArticle)

		/**
		* showdoc
		* @catalog 点赞相关
		* @title 文章点赞
		* @description 给文章点赞数加一
		* @method post
		* @url /api/v1/article/like/id
		* @param id 必选 int 主键
		* @return {"code":0,"data":{"id":1,"title":"自序","content":"","catalog":"建站日记","view":4,"like":1,"created_at":"2020年2月12日 14:16:11"},"msg":""}
		 */
		v1.POST("article/like/:id", api.AddArticleLike)

		/**
		* showdoc
		* @catalog 文章相关
		* @title 文章列表
		* @description 查看文章具体内容的api。
		* @method post
		* @url /api/v1/articles
		* @return {"code":0,"data":[{"id":1,"title":"自序","content":"","catalog":"建站日记","view":4,"like":0,"created_at":"2020年2月12日 14:16:11"},{"id":2,"title":"测试","content":"","catalog":"建站日记","view":0,"like":0,"created_at":"2020年2月12日 17:50:44"},{"id":3,"title":"测试2","content":"","catalog":"建站日记","view":0,"like":0,"created_at":"2020年2月12日 17:52:30"}],"msg":""}
		 */
		v1.GET("articles", api.ListArticle)

		v1.GET("catalogs", api.ListCatalog)

		/**
		* showdoc
		* @catalog 排行相关
		* @title 文章总排行top10文章
		* @description 查看文章总排行top10的api。
		* @method get
		* @url /api/v1/rank/total
		* @return {"code":0,"data":[{"id":1,"title":"自序","content":"","catalog":"建站日记","view":4,"created_at":"2020年2月12日 14:16:11"},{"id":2,"title":"测试","content":"","catalog":"建站日记","view":3,"created_at":"2020年2月12日 17:50:44"},{"id":3,"title":"测试2","content":"","catalog":"建站日记","view":1,"created_at":"2020年2月12日 17:52:30"}],"msg":""}
		* @remark 有点击的量才会返回，点击量为0的不会返回。
		 */
		v1.GET("rank/total", api.TotalRank)

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
			* @return {"code":0,"data":{"id":1,"user_name":"explorer","nickname":"lovelyjax","status":"active","avatar":"","created_at":1580879384},"msg":""}
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
