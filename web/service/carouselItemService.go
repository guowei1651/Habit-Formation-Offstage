package services

import (
	"log"
	dao "hf/web/dao"
)

func findAllCarouselItemsByCarouseId(carouseId int) ([]CarouselItem, error) {
	log.Printf("查询轮播项 carouseId:%v", carouseId)
	return dao.FindAllCarouselItemsByCarouselId(carouseId)
}