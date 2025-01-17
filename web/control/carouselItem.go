package control

import (
	"log"
	"net/http"
	"strconv"

	common "hf/web/common"
	utils "hf/web/utils"
	service "hf/web/service"

	restful "github.com/emicklei/go-restful/v3"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
)

type CarouselItemResource struct {
}

func (ci *CarouselItemResource) FindAllCarouselItemsByCarouseId(request *restful.Request, response *restful.Response) {
	userId := utils.GetUserId(request)
	if len(userId) == 0 {
		response.WriteErrorString(http.StatusNotFound, "plases login")
		return
	}

	carouselId := request.PathParameter("id")
	id, err := strconv.Atoi(carouselId)
	if err != nil {
		log.Printf("findAllCarouselItemsByCarouseId param error id->", id)
		response.WriteErrorString(http.StatusNotFound, "find CarouselItems params error.")
		return
	}

	result, err := service.FindAllCarouselItemsByCarouseId(id)
	if err != nil {
		log.Printf("findAllCarouselItemsByCarouseId db query error! err->", err)
		response.WriteErrorString(http.StatusNotFound, "CarouselItems could not be found.")
		return 
	}

	log.Printf("findAllCarouselItemsByCarouseId db query result->", result)
	respBody := common.ResponseBody{}
	respBody.Code = 0
	respBody.Message = "success"
	respBody.Data = result
	response.WriteAsJson(respBody)
}

func (carouselItemResource *CarouselItemResource) LoadRoute(ws *restful.WebService) {
	tags := []string{"hf"}

	ws.Route(ws.GET("{id}/carouselItems").To(carouselItemResource.FindAllCarouselItemsByCarouseId).
		// docs
		Doc("get all carousel items in carousel").
		Param(ws.PathParameter("id", "identifier of the Carousel").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(common.ResponseBody{}).
		Returns(200, "OK", common.ResponseBody{}))
}