# Borderland

个人使用博客Borderland，意为无主之地的意思，也意为无聊之地。
主要使用go -> gin ->singo框架进行开发。

## 框架介绍

### gin地址
https://github.com/gin-gonic/gin

### Singo地址
https://github.com/Gourouting/singo

## 后端开发文档
https://www.showdoc.cc/jaxchan?page_id=3846945987031560

## 框架简介

本项目划分出下列模块:

1. api文件夹就是MVC框架的controller，负责协调各部件完成任务
2. model文件夹负责存储数据库模型和数据库操作相关的代码
3. service负责处理比较复杂的业务，把业务代码模型化可以有效提高业务代码的质量（比如用户注册，充值，下单等）
4. serializer储存通用的json模型，把model得到的数据库模型转换成api需要的json对象
5. cache负责redis缓存相关的代码
6. auth权限控制文件夹
7. util一些通用的小工具
8. conf放一些静态存放的配置文件，其中locales内放置翻译相关的配置文件

## Godotenv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env文件设置环境变量便于使用(建议开发环境使用)

```shell
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接地址
REDIS_ADDR="127.0.0.1:6379" # Redis端口和地址
REDIS_PW="" # Redis连接密码
REDIS_DB="" # Redis库从0到10
SESSION_SECRET="setOnProducation" # Seesion密钥，必须设置而且不要泄露
GIN_MODE="debug"
```

## Go Mod

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```shell
go mod init borderland
go env -w GOPROXY=https://goproxy.cn,direct
go run main.go // 将会自动包依赖安装
```

## 运行

1. 安装Mysql和[Redis](https://blog.csdn.net/weixin_40976261/article/details/89854520)
2. 设置数据库参数

在项目根目录创建.env文件设置环境变量便于使用。
需要手动创库，但是会自动创表。

```shell
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接地址
REDIS_ADDR="127.0.0.1:6379" # Redis端口和地址
REDIS_PW="" # Redis连接密码
REDIS_DB="" # Redis库从0到10
SESSION_SECRET="setOnProducation" # Seesion密钥，必须设置而且不要泄露
GIN_MODE="debug"
```

3. 在任意文件夹下clone项目（不要在GOPATH和GOROOT的src下clone）并且运行
```shell
git clone https://github.com/JaxChan25/borderland.git
go env -w GOPROXY=https://goproxy.cn,direct
cd borderland
go run main.go
```

## 关于自动生成后端api文档
请查看[showdocc 自动生成API文档](https://www.showdoc.cc/page/741656402509783)

### 标准语法
```
        /**
        * showdoc
        * @catalog 测试文档/用户相关
        * @title 测试api文档自动生成
        * @description 测试api文档自动生成
        * @method get
        * @url https://www.showdoc.cc/home/user/login
        * @header token 可选 string 设备token 
        * @param username 必选 string 用户名 
        * @param password 必选 string 密码  
        * @param name 可选 string 用户昵称  
        * @return {"error_code":0,"data":{"uid":"1","username":"12154545","name":"吴系挂","groupid":2,"reg_time":"1436864169","last_login_time":"0"}}
        * @return_param groupid int 用户组id
        * @return_param name string 用户昵称
        * @remark 这里是备注信息
        * @number 99
        */
```
