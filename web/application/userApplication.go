package application

import (
	control "hf/web/control"

	restful "github.com/emicklei/go-restful/v3"
)

type UserApplication struct {

}

func (ua *UserApplication)LoadRoute() (*restful.WebService) {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	loginResource := control.LoginResource{}
	loginResource.LoadRoute(ws)
	
	return ws
}