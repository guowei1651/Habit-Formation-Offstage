package control

import (
	"log"
	"strconv"
	"net/http"

	common "hf/web/common"
	utils "hf/web/utils"
	service "hf/web/service"

	restful "github.com/emicklei/go-restful/v3"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
)

type CarouselResource struct {

}

func (carouselResource *CarouselResource)FindAllCarouselByOwnerId(request *restful.Request, response *restful.Response) {
	userId := utils.GetUserId(request)
	if len(userId) == 0 {
		response.WriteErrorString(http.StatusNotFound, "plases login")
		log.Printf("findAllCarouselItemsByCarouseId param error userId->", userId)
		return
	}

	id, err := strconv.Atoi(userId)
	if err != nil {
		log.Printf("findAllCarouselItemsByCarouseId conv atoi error id->", userId)
		response.WriteErrorString(http.StatusNotFound, "plases login")
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

func (carouselResource *CarouselResource) LoadRoute() (*restful.WebService) {
	ws := new(restful.WebService)
	ws.
		Path("/carousels").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	tags := []string{"hf"}

	ws.Route(ws.GET("").To(carouselResource.FindAllCarouselByOwnerId).
		// docs
		Doc("get all carousel by owner").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(common.ResponseBody{}).
		Returns(200, "OK", common.ResponseBody{}))
	
	return ws
}