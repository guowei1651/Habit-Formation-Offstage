package dao

import (
    "log"
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
        if (order.Valid) { ci.Order, _ = order.Value() } else{ ci.Order = 0 }
        if (genus.Valid) { ci.Genus, _ = genus.Value() } else{ ci.Genus = "0" }
        if (relationsId.Valid) { ci.RelationsId, _ = relationsId.Value() } else{ ci.RelationsId = 0 }
        if (alertLevel.Valid) { ci.AlertLevel, _ = alertLevel.Value() } else{ ci.AlertLevel = "0" }
        if (triggerTime.Valid) { ci.TriggerTime, _ = triggerTime.Value() } else{ ci.TriggerTime = "" }
        if (duration.Valid) { ci.Duration, _ = duration.Value() } else{ ci.Duration = 0 }
        if (chartUrl.Valid) { ci.ChartUrl, _ = chartUrl.Value() } else{ ci.ChartUrl = "" }
        carouselItems = append(carouselItems, ci)
    }
    if err = rows.Err(); err != nil {
        log.Fatal("db row next error. err->", err)
        return nil, err
    }
    return carouselItems, err
}
