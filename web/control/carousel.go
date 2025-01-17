package control

import (
	"log"
	"strconv"

	utils "hf/web/utils"
	service "hf/web/service"

	restful "github.com/emicklei/go-restful/v3"
)

type UserVO struct {
	UserName string `json:"username" description:"user name"`
	Password string `json:"password" description:"password"`
}

type CarouselResource struct {

}

func (carouselResource *CarouselResource)FindAllCarouselByOwnerId(request *restful.Request, response *restful.Response) {
	userId := utils.GetUserId(request)
	if userId == nil {
		response.WriteErrorString(http.StatusNotFound, "plases login")
		return
	}

	id,_ := strconv.Atoi(request.PathParameter("id"))
	if id == 0 {
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
	var respBody ResponseBody
	respBody.Code = 0
	respBody.Message = "success"
	respBody.Data = result
	response.WriteAsJson(respBody)
}

func (carouselResource *CarouselResource) loadRoute(ws *restful.WebService) (*restful.WebService) {
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
		Writes(ResponseBody{}).
		Returns(200, "OK", ResponseBody{}))
}