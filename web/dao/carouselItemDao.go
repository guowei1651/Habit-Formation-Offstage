package dao

import (
    "log"
    "strconv"
    "database/sql"
    db "hf/database"
)

type CarouselItem struct {
	Order		int		`json:"order" description:"Carousel Item on carousel order" default:"1"`
	Genus		string	`json:"type" description:"type of the CarouselItem" default:"image"`
	RelationsId	int		`json:"relations_id" description:"relations_id of the CarouselItem" default:"0"`
	AlertLevel	string	`json:"alert_level" description:"alert_level of the CarouselItem" default:"norme"`
	TriggerTime	string	`json:"trigger_time" description:"trigger_time of the CarouselItem" default:""`
	Duration	int		`json:"duration" description:"duration of the CarouselItem" default:"30"`
	ChartUrl	string 	`json:"chartUrl" description:"chartUrl of the CarouselItem" default:""`
}

func FindAllCarouselItemsByCarouselId(carouselId int) ([]CarouselItem, error) {
    log.Printf("sqlSelectAllCarouselItemsByCarouselId param->", carouselId)
    rows, err := db.DBConnectPool.Query(`
SELECT carousel_item.order, carousel_item.type, carousel_item.relations_id, carousel_item.alert_level,
        carousel_item.trigger_time, carousel_item.duration, carousel_item.chart_url
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
        var order sql.NullInt32
        var genus sql.NullString
        var relationsId sql.NullInt32
        var alertLevel sql.NullString
        var triggerTime sql.NullString
        var duration sql.NullInt32
        var chartUrl sql.NullString
        if err := rows.Scan(&order, &genus, &relationsId, &alertLevel, &triggerTime, &duration, &chartUrl); err != nil {
            log.Fatal("db row next error. err->", err)
            return carouselItems, err
        }
        log.Printf("row ->", order, genus, duration, chartUrl)
        ci.Order = (order.Valid? order.value : 0)
        ci.Genus = (genus.Valid? genus.value : "0")
        ci.RelationsId = (relationsId.Valid? relationsId.value : 0)
        ci.AlertLevel = (alertLevel.Valid? alertLevel.value : "0")
        ci.TriggerTime = (triggerTime.Valid? triggerTime.value : "")
        ci.Duration = (duration.Valid? duration.value : 0)
        ci.ChartUrl = (chartUrl.Valid? chartUrl.value : "")
        carouselItems = append(carouselItems, ci)
    }
    if err = rows.Err(); err != nil {
        log.Fatal("db row next error. err->", err)
        return nil, err
    }
    return carouselItems, err
}
