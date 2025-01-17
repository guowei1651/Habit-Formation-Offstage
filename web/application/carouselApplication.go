package application

import (
	"log"
	"strconv"
	"net/http"

	control "hf/web/control"

	restful "github.com/emicklei/go-restful/v3"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
)

type CarouselApplication struct {

}

func (c *CarouselApplication)LoadRoute() (*restful.WebService) {
	ws := new(restful.WebService)
	ws.
		Path("/carousels").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	carouselResource := control.CarouselResource{}
	carouselResource.LoadRoute(ws)

	carouselItemResource := control.CarouselItemResource{}
	carouselItemResource.LoadRoute(ws)
	return ws
}