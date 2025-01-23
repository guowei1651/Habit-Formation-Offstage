package application

import (
	control "hf/web/control"

	restful "github.com/emicklei/go-restful/v3"
)

type HabbitApplication struct {

}

func (ha *HabbitApplication)LoadRoute() (*restful.WebService) {
	ws := new(restful.WebService)
	ws.
		Path("/habits").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	habitRecordResource := control.HabitRecordResource{}
	habitRecordResource.LoadRoute(ws)
	
	return ws
}