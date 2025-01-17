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

type UserResource struct {

}

// POST http://localhost:8080/users/login
func (u *UserResource) Login(request *restful.Request, response *restful.Response) {
	log.Println("User Login")
	usr := UserVO{}
	
	if err := request.ReadEntity(&usr); err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	respBody := common.ResponseBody{}
	log.Println("User Login name is : %v", usr.UserName)
	
	loginInfo, err := service.Login(usr.UserName, usr.Password)
	if (err != nil) {
		respBody.Code = -1
		respBody.Message = err.Error()
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	respBody.Code = 0
	respBody.Message = "success"
	respBody.Data = loginInfo
	response.WriteHeaderAndEntity(http.StatusCreated, usr)
}

func (userResource *UserResource) LoadRoute() (*restful.WebService) {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	tags := []string{"hf"}

	ws.Route(ws.POST("/login").To(userResource.Login).
		Doc("user login").
		Reads(UserVO{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(common.ResponseBody{}).
		Returns(200, "OK", common.ResponseBody{}))

	return ws
}