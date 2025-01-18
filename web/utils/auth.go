package utils

import (
	"fmt"
	restful "github.com/emicklei/go-restful/v3"
)

func GetUserId(request *restful.Request) string {
	for k, v := range request.Request.Header {
        fmt.Printf("Header field %q, Value %q\n", k, v)
    }
	return request.Request.Header.Get("HF-User-Id")
}