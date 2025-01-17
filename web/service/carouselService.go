package service

import (
	"log"
	dao "hf/web/dao"
)

func FindAllCarouselByOwnerId(ownerId int) ([]dao.Carousel, error) {
	log.Printf("根据所有者查询轮播。 OwnerId:%v", ownerId)
	return dao.FindAllCarouselByOwnerId(ownerId)
}