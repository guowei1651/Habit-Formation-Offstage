package utils

import (
	restful "github.com/emicklei/go-restful/v3"
)

func GetUserId(request *restful.Request) string {
	for k, v := range r.Header {
        fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
    }
	return request.Header.Get("HF-User-Id")
}