package application

import (
	"log"
	"net/http"

	control "hf/web/control"

	restful "github.com/emicklei/go-restful/v3"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
)

type UserApplication strucet {

}

func (ua *UserApplication)LoadRoute() {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
		
	userResource := control.UserResource{}
	userResource.LoadRoute(ws)
}