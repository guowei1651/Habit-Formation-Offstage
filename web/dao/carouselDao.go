package dao

import (
    "log"
	"time"
    db "hf/database"
)

type Carousel struct {
	ID          int     `json:"id" description:"Carousel id"`
	Name        string  `json:"name" description:"Carousel name"`
	Description string  `json:"description" description:"Carousel Description"`
	OwnerId     int     `json:"ownerId" description:"Carousel owner user id"`
	BrownserUrl string  `json:"brownserUrl" description:"Carousel brownser URL"`
	CreatedAt   time.Time   `json:"createAt" description:"Carousel create At"`
	UpdatedAt   time.Time   `json:"updatedAt" description:"Carousel update At"`
}

func FindAllCarouselByOwnerId(ownerId int) ([]Carousel, error) {
    log.Printf("sqlSelectAllCarouselItemsByCarouselId param->", ownerId)
    rows, err := db.DBConnectPool.Query(`
select c.id, c.name, c.description, c.owner_id, c.brownser_url, c.created_at, c.updated_at 
from carousel c 
where owner_id = $1 and delete_flag = delete_flag ;`, ownerId)
    if err != nil {
        log.Fatal("sqlSelectAllCarouselItemsByCarouselId db query error. err->", err)
        return nil, err
    }
    defer rows.Close()

    log.Printf("sqlSelectAllCarouselItemsByCarouselId result->", rows)

    var carousels []Carousel

    for rows.Next() {
        var c Carousel
		var id int
		var name string
		var description string
		var ownerId int
		var brownserUrl string
		var createdAt time.Time
		var updatedAt time.Time
        if err := rows.Scan(&id, &name, &description, &ownerId, &brownserUrl, &createdAt, &updatedAt); err != nil {
            log.Fatal("db row next error. err->", err)
            return carousels, err
        }
        log.Printf("row ->", id, name, description, ownerId, brownserUrl, createdAt, updatedAt)
		c.ID = id
		c.Name = name
		c.Description = description
		c.OwnerId = ownerId
		c.BrownserUrl = brownserUrl
		c.CreatedAt = createdAt
		c.UpdatedAt = updatedAt
        carousels = append(carousels, c)
    }
    if err = rows.Err(); err != nil {
        log.Fatal("db row next error. err->", err)
        return nil, err
    }


    return carousels, err
}