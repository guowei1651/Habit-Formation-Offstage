package dao

import (
    "log"
    db "../database/postgres"
)

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


