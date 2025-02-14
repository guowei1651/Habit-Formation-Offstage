package control

import (
	"log"
	"net/http"
	
	common "hf/web/common"
	service "hf/web/service"

	restful "github.com/emicklei/go-restful/v3"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
)

type UserVO struct {
	UserName string `json:"username" description:"user name"`
	Password string `json:"password" description:"password"`
}

type LoginResource struct {
}

// POST http://localhost:8080/users/login
func (u *LoginResource) Login(request *restful.Request, response *restful.Response) {
	log.Println("User Login")
	usr := UserVO{}
	
	if err := request.ReadEntity(&usr); err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	if len(usr.UserName) == 0 || len(usr.Password) ==0 {
		response.WriteErrorString(http.StatusInternalServerError, "请填写用户名密码之后再点击登录")
		return
	}

	respBody := common.ResponseBody{}
	log.Println("User Login name is : ", usr.UserName)
	
	loginInfo, err := service.Login(usr.UserName, usr.Password)
	if (err != nil) {
		respBody.Code = -1
		respBody.Message = err.Error()
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	log.Println("Login loginInfo->", loginInfo)
	if (loginInfo == nil) {
		respBody.Code = -1
		respBody.Message = "not found"
	} else {
		respBody.Code = 0
		respBody.Message = "success"
		respBody.Data = loginInfo
	}
	response.WriteHeaderAndEntity(http.StatusCreated, respBody)
}

func (loginResource *LoginResource) LoadRoute(ws *restful.WebService) {
	tags := []string{"hf"}

	ws.Route(ws.POST("/login").To(loginResource.Login).
		Doc("user login").
		Reads(UserVO{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(common.ResponseBody{}).
		Returns(200, "OK", common.ResponseBody{}))
}