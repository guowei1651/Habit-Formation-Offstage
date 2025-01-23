package dao

import (
	"log"
	"time"

	"context"
	"database/sql"
	db "hf/database"
)

type HabitRecord struct {
	Type int64
	RelationsId int64
	Serial string
	RecordAt time.Time
	Remark string
}

func Record(Type int64, RelationsId int64, Serial string, Remark string) (error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := db.DBConnectPool.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = tx.ExecContext(ctx, `INSERT INTO habit_raw_records ("type",relations_id,serial,remark) VALUES ($1,$2,'$3','$4');`, 
			Type, RelationsId, Serial, Remark)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("insert drivers: unable to rollback: %v", rollbackErr)
		}
		log.Fatalf("ExecContext error: %v", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("Commit error: %v", err)
		return err
	}
	return nil
}