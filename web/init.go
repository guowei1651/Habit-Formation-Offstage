// https://github.com/emicklei/go-restful/blob/v3/examples/user-resource/restful-user-resource.go
package web

import (
	"log"
	"fmt"
	"net/http"

	hfConfig "hf/config"
	app "hf/web/application"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	restful "github.com/emicklei/go-restful/v3"
	"github.com/go-openapi/spec"
)

type WebServer struct {}

func (webServer WebServer) LoadWebService() {
	ca := app.CarouselApplication{}
	restful.DefaultContainer.Add(ca.LoadRoute())

	ua := app.UserApplication{}
	restful.DefaultContainer.Add(ua.LoadRoute())

	ha := app.HabbitApplication{}
	restful.DefaultContainer.Add(ha.LoadRoute())
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "UserService",
			Description: "Resource for managing Users",
			Contact: &spec.ContactInfo{
				ContactInfoProps: spec.ContactInfoProps{
					Name:  "wales",
					Email: "wales.kuo@gmail.com",
					URL:   "http://guowei1651.github.io",
				},
			},
			License: &spec.License{
				LicenseProps: spec.LicenseProps{
					Name: "MIT",
					URL:  "http://mit.org",
				},
			},
			Version: "1.0.0",
		},
	}
	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "CarouselItem",
		Description: "get CarouselItem"}}}
}

func OpenServer(ch chan string) {
	webServer := WebServer{}
	webServer.LoadWebService()
	
	restConfig := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(restConfig))

	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("/Users/emicklei/Projects/swagger-ui/dist"))))

	// Optionally, you may need to enable CORS for the UI to work.
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept", "HF-User-Id"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		CookiesAllowed: false,
		Container:      restful.DefaultContainer}
	restful.DefaultContainer.Filter(cors.Filter)

	log.Printf("Get the API using http://xxx/apidocs.json")
	log.Printf("Open Swagger UI using http://xxx/apidocs/?url=http://xxx/apidocs.json")
	portStr := fmt.Sprintf(":%d", hfConfig.Config.WEBConfig.Port)
	err := http.ListenAndServe(portStr, nil)
	log.Fatal(err)
	ch <- err.Error()
}