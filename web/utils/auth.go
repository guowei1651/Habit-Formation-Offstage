package utils

import (
	"fmt"
	restful "github.com/emicklei/go-restful/v3"
)

func GetUserId(request *restful.Request) string {
	return request.HeaderParameter("HF-User-Id")
}