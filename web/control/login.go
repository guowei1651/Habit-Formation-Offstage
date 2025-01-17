package control

import (
	"log"
	"net/http"
	
	"hf/web/common"
	service "hf/web/service"

	restful "github.com/emicklei/go-restful/v3"
)

type UserVO struct {
	UserName string `json:"username" description:"user name"`
	Password string `json:"password" description:"password"`
}

type UserResource struct {

}

// POST http://localhost:8080/users/login
func (u *UserResource) login(request *restful.Request, response *restful.Response) {
	log.Println("User Login")
	usr := UserVO{}
	
	if err := request.ReadEntity(&usr); err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	respBody := ResponseBody{}
	log.Println("User Login name is : %v", usr.UserName)
	
	loginInfo, err := service.Login(usr.UserName, usr.Password)
	if (err != nil) {
		respBody.code = -1
		respBody.message = err.Error()
		response.WriteError(http.StatusInternalServerError, respBody)
		return
	}
	respBody.code = 0
	respBody.message = "success"
	respBody.data = loginInfo
	response.WriteHeaderAndEntity(http.StatusCreated, usr)
}

func (userResource *UserResource) loadRoute() (*restful.WebService) {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	tags := []string{"hf"}

	ws.Route(ws.POST("/login").To(userResource.findAllCarouselItemsByCarouseId).
		Doc("user login").
		Reads(User{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseBody{}).
		Returns(200, "OK", ResponseBody{}))

	return ws
}