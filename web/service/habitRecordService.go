package service

import (
	"log"
	dao "hf/web/dao"
)

type HabitRecord struct {
	Type int64
	RelationsId int64
	Serial string
	RecordAt time.Time
	Remark string
}

func Record(Type int64, RelationsId int64, Serial string, Remark string) (error) {
	log.Printf("习惯计数 Type:%v RelationsId:%v Serial:%v Remark:%v", Type, RelationsId, Serial, Remark)
	dao.Record(Type, RelationsId, Serial, Remark)
}