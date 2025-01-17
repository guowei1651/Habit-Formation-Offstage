package service

import (
	"log"
	dao "hf/web/dao"
)

func FindAllCarouselItemsByCarouseId(carouseId int) ([]dao.CarouselItem, error) {
	log.Printf("查询轮播项 carouseId:%v", carouseId)
	return dao.FindAllCarouselItemsByCarouselId(carouseId)
}