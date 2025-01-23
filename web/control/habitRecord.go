package control

import (
	"log"
	"net/http"
	
	common "hf/web/common"
	service "hf/web/service"

	restful "github.com/emicklei/go-restful/v3"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
)

type HabitRecordVO struct {
	Type int64			`json:"type" description:"type"`
	RelationsId int64	`json:"relations_id" description:"relations id"`
	Serial string		`json:"serial" description:"serial"`
	Remark string		`json:"remark" description:"remark"`
}

type HabitRecordResource struct {
}

// POST http://localhost:8080/habit/${id}/record
func (h *HabitRecordResource) Record(request *restful.Request, response *restful.Response) {
	log.Println("User Login")
	habitRecord := HabitRecordVO{}
	if err := request.ReadEntity(&habitRecord); err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	// 轮播项的类型。1：代表美图，2：代表提醒，3：代表习惯，4：代表长日程
	if (habitRecord.Type == 1 || habitRecord.Type == 2 || habitRecord.Type == 3 || habitRecord.Type == 4) {
		response.WriteErrorString(http.StatusInternalServerError, "不支持的类型")
		return
	}

	respBody := common.ResponseBody{}
	err := service.Record(habitRecord.Type, habitRecord.RelationsId, habitRecord.Serial, habitRecord.Remark)
	if (err != nil) {
		respBody.Code = -1
		respBody.Message = err.Error()
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	respBody.Code = 0
	respBody.Message = "success"
	
	response.WriteHeaderAndEntity(http.StatusCreated, respBody)
}

func (h *HabitRecordResource) LoadRoute(ws *restful.WebService) {
	tags := []string{"hf"}

	ws.Route(ws.GET("{id}/record").To(h.Record).
		// docs
		Doc("Habit execution record").
		Param(ws.PathParameter("id", "identifier of the habit").DataType("integer").DefaultValue("0")).
		Reads(HabitRecordVO{}).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(common.ResponseBody{}).
		Returns(200, "OK", common.ResponseBody{}))
}