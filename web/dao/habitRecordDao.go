package dao

import (
	"log"

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
	tx, err := db.DBConnectPool.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO appsmith.habit_raw_records ("type",relations_id,serial,remark) VALUES (?,?,?,?);", Type, RelationsId, Serial, Remark)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("update drivers: unable to rollback: %v", rollbackErr)
		}
		log.Fatal(err)
		return err
	}
	_, err = tx.ExecContext(ctx, "UPDATE pickups SET driver_id = $1;", id)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("update failed: %v, unable to back: %v", err, rollbackErr)
		}
		log.Fatal(err)
		return err
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
		return err
	}
}