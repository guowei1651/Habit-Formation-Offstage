package utils

func GetUserId(request *restful.Request) string {
	return request.HeaderParameter("HF-User-Id")
}