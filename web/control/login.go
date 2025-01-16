package control

import (
	"log"
	"net/http"
	service "hf/web/service"

	"github.com/go-openapi/spec"
)

type UserVO struct {
	UserName string
	Password string
}

// POST http://localhost:8080/users
// <User><Id>1</Id><Name>Melissa</Name></User>
func (u *UserResource) login(request *restful.Request, response *restful.Response) {
	log.Println("createUser")
	usr := User{}
	
	if err := request.ReadEntity(&usr); err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	respBody := ResponseBody{}
	
	loginInfo, err := service.Login(usr.UserName, usr.Password)
	if (err != null) {
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

func aa () {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well

	tags := []string{"hf"}

	ws.Route(ws.POST("").To(webServer.findAllCarouselItemsByCarouseId).
		// docs
		Doc("get all carousel items in carousel").
		Param(ws.PathParameter("id", "identifier of the Carousel").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseBody{}).
		Returns(200, "OK", ResponseBody{}))

	ws.Route(ws.POST("/login").To(u.createUser).
		// docs
		Doc("create a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(User{})) // from the request


	return ws
}