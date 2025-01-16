package control

import (
	"log"
	"net/http"
	"strconv"
	"github.com/emicklei/go-restful"
	"net/http"
	service "hf/web/service"

	"github.com/go-openapi/spec"
)



type WEBResource struct {
	getRoute() ()
}

func (ci CarouselItemResource) findAllCarouselItemsByCarouseId(request *restful.Request, response *restful.Response) {
	id,_ := strconv.Atoi(request.PathParameter("id"))
	if id == 0 {
		log.Printf("findAllCarouselItemsByCarouseId param error id->", id)
		response.WriteErrorString(http.StatusNotFound, "find CarouselItems params error.")
		return
	}

	result, err := sqlSelectAllCarouselItemsByCarouselId(id)
        if err != nil {
		log.Printf("findAllCarouselItemsByCarouseId db query error! err->", err)
		response.WriteErrorString(http.StatusNotFound, "CarouselItems could not be found.")
		return 
	}

	log.Printf("findAllCarouselItemsByCarouseId db query result->", result)
	var respBody ResponseBody
	respBody.Code = 0
	respBody.Message = "success"
	respBody.Data = result
	response.WriteAsJson(respBody)
}

func aa () {
	tags := []string{"carouselItems"}

	ws.Route(ws.GET("{id}/carouselItems").To(webServer.findAllCarouselItemsByCarouseId).
		// docs
		Doc("get all carousel items in carousel").
		Param(ws.PathParameter("id", "identifier of the Carousel").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ResponseBody{}).
		Returns(200, "OK", ResponseBody{}))
}