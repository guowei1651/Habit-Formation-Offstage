package dao

import (
    "log"
    "database/sql"
    db "hf/database"
)

type CarouselItem struct {
	Order		int64	`json:"order" description:"Carousel Item on carousel order" default:"1"`
	Genus		string	`json:"type" description:"type of the CarouselItem" default:"image"`
	RelationsId	int64	`json:"relations_id" description:"relations_id of the CarouselItem" default:"0"`
	AlertLevel	string	`json:"alert_level" description:"alert_level of the CarouselItem" default:"norme"`
	TriggerTime	string	`json:"trigger_time" description:"trigger_time of the CarouselItem" default:""`
	Duration	int64	`json:"duration" description:"duration of the CarouselItem" default:"30"`
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
        var order sql.NullInt64
        var genus sql.NullString
        var relationsId sql.NullInt64
        var alertLevel sql.NullString
        var triggerTime sql.NullString
        var duration sql.NullInt64
        var chartUrl sql.NullString
        if err := rows.Scan(&order, &genus, &relationsId, &alertLevel, &triggerTime, &duration, &chartUrl); err != nil {
            log.Fatal("db row next error. err->", err)
            return carouselItems, err
        }
        log.Printf("row ->", order, genus, duration, chartUrl)
        if (order.Valid) { tmp, _ := order.Value(); ci.Order = tmp.(int64) } else { ci.Order = 0 }
        if (genus.Valid) { tmp, _ := genus.Value(); ci.Genus = tmp.(string) } else{ ci.Genus = "0" }
        if (relationsId.Valid) { tmp, _ := relationsId.Value(); ci.RelationsId = tmp.(int64) } else{ ci.RelationsId = 0 }
        if (alertLevel.Valid) { tmp, _ := alertLevel.Value(); ci.AlertLevel = tmp.(string) } else{ ci.AlertLevel = "0" }
        if (triggerTime.Valid) { tmp, _ := triggerTime.Value(); ci.TriggerTime = tmp.(string) } else{ ci.TriggerTime = "" }
        if (duration.Valid) { tmp, _ := duration.Value(); ci.Duration = tmp.(int64) } else{ ci.Duration = 0 }
        if (chartUrl.Valid) { tmp, _ := chartUrl.Value(); ci.ChartUrl = tmp.(string) } else{ ci.ChartUrl = "" }
        carouselItems = append(carouselItems, ci)
    }
    if err = rows.Err(); err != nil {
        log.Fatal("db row next error. err->", err)
        return nil, err
    }
    return carouselItems, err
}
