package api

import (
	"borderland/serializer"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

// OwnerService 网站静态信息的服务
type OwnerService struct {
	Name       string `form:"name" json:"name"`
	Profession string `form:"profession" json:"profession"`
	School     string `form:"school" json:"school"`
	Address    string `form:"address" json:"address"`
	Email      string `form:"email" json:"email"`
	Hobby      string `form:"hobby" json:"hobby"`
}

//StaticOwner 返回owner的静态信息
func StaticOwner(c *gin.Context) {

	file, err := os.Open("./OwnerInfo.json")
	defer file.Close()

	//有bug 不能成功返回错误
	if err != nil {
		fmt.Println("I am here")
		c.JSON(200, serializer.Response{
			Code: serializer.CodeEnv,
			Msg:  err.Error(),
		})

	} else {
		content, _ := ioutil.ReadAll(file)

		var service OwnerService
		json.Unmarshal([]byte(content), &service)

		c.JSON(200, serializer.Response{
			Code: 0,
			Data: service,
		})
	}

}
