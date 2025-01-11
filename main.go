package main

import (
	"log"
	"context"
	"time"
	"strconv"
	"net/http"

	"database/sql"
	_ "github.com/lib/pq"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	restful "github.com/emicklei/go-restful/v3"
	"github.com/go-openapi/spec"
)

var db *sql.DB

type CarouselItem struct {
	Order		int    	`json:"order" description:"Carousel Item on carousel order" default:"1"`
	Genus		string 	`json:"type" description:"type of the CarouselItem" default:"image"`
	Duration	int 	`json:"duration" description:"duration of the CarouselItem" default:"30"`
	ChartUrl	string 	`json:"chartUrl" description:"chartUrl of the CarouselItem" default:""`
}

type ResponseBody struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type CarouselItemResource struct {
	// normally one would use DAO (data access object)
	users map[string]CarouselItem
}

func Ping(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	log.Printf("db ping db-> ", db)
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}

func sqlOpen() {
    log.Printf("open db start")
    var err error
    //db, err = sql.Open("postgres", "port=5432 user=appsmith password=appsmith dbname=appsmith sslmode=disable")
    db, err = sql.Open("postgres", "postgres://appsmith:appsmith@172.25.1.22:5432/appsmith?sslmode=disable")
    log.Printf("open db complete -> ", db)
    db.SetConnMaxIdleTime(30*1000)
    db.SetConnMaxLifetime(10*1000)
    db.SetMaxIdleConns(10)
    db.SetMaxOpenConns(20)
    if err != nil {
        log.Fatal("open db connect fail -> ", err)
        panic(err)
    }

    ctx, stop := context.WithCancel(context.Background())
    defer stop()
    Ping(ctx)
}

func sqlSelectAllCarouselItemsByCarouselId(carouselId int) ([]CarouselItem, error) {
    log.Printf("sqlSelectAllCarouselItemsByCarouselId param->", carouselId)
    rows, err := db.Query(`SELECT carousel_item.order, carousel_item.type, carousel_item.duration, carousel_item.chart_url
FROM carousel_item 
WHERE carousel_id = $1 AND delete_flag = FALSE ORDER BY carousel_item.order;`, carouselId)
    if err != nil {
	log.Fatal("sqlSelectAllCarouselItemsByCarouselId db query error. err->", err)
        return nil, err
    }
    defer rows.Close()

    var carouselItems []CarouselItem

    for rows.Next() {
        var ci CarouselItem
	var order string
	var genus string
	var duration string
	var chartUrl string
        if err := rows.Scan(&order, &genus, &duration, &chartUrl); err != nil {
	    log.Fatal("db row next error. err->", err)
            return carouselItems, err
        }
	log.Printf("row ->", order, genus, duration, chartUrl)
	ci.Order, _ = strconv.Atoi(order)
	ci.Genus = genus
	ci.Duration, _ = strconv.Atoi(duration)
	ci.ChartUrl = chartUrl
        carouselItems = append(carouselItems, ci)
    }
    if err = rows.Err(); err != nil {
	log.Fatal("db row next error. err->", err)
        return nil, err
    }
    return carouselItems, err
}

func (ci CarouselItemResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/carousel").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well

	tags := []string{"carouselItems"}

	ws.Route(ws.GET("{id}/carouselItems").To(ci.findAllCarouselItemsByCarouseId).
		// docs
		Doc("get all carousel items in carousel").
		Param(ws.PathParameter("id", "identifier of the Carousel").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]CarouselItem{}).
		Returns(200, "OK", []CarouselItem{}))

	return ws
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

func main() {
	sqlOpen()

	u := CarouselItemResource{map[string]CarouselItem{}}
	restful.DefaultContainer.Add(u.WebService())

	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))

	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("/Users/emicklei/Projects/swagger-ui/dist"))))

	// Optionally, you may need to enable CORS for the UI to work.
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET"},
		CookiesAllowed: false,
		Container:      restful.DefaultContainer}
	restful.DefaultContainer.Filter(cors.Filter)

	log.Printf("Get the API using http://xxx/apidocs.json")
	log.Printf("Open Swagger UI using http://xxx/apidocs/?url=http://xxx/apidocs.json")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
