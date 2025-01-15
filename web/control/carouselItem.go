package control

import (
	"log"
	"net/http"
	"strconv"
	"github.com/emicklei/go-restful"
)

type CarouselItem struct {
	Order		int    	`json:"order" description:"Carousel Item on carousel order" default:"1"`
	Genus		string 	`json:"type" description:"type of the CarouselItem" default:"image"`
	Duration	int 	`json:"duration" description:"duration of the CarouselItem" default:"30"`
	ChartUrl	string 	`json:"chartUrl" description:"chartUrl of the CarouselItem" default:""`
}

type WEBResource struct {
	

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

