// https://jianghushinian.cn/2023/06/05/how-to-use-database-sql-to-operate-database-in-go/

package database

import (
    "log"
    "time"
    "context"
    "database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Ping(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	log.Printf("db ping db-> ", db)
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}

func sqlOpen() {
    log.Printf("open db start")
    var err error
    db, err = sql.Open("postgres", "host= port=5432 user=appsmith password=appsmith dbname=appsmith sslmode=disable")
    // db, err = sql.Open("postgres", "postgres://appsmith:appsmith@172.25.1.22:5432/appsmith?sslmode=disable")
    log.Printf("open db complete -> ", db)
    db.SetConnMaxIdleTime(30*1000)
    db.SetConnMaxLifetime(10*1000)
    db.SetMaxIdleConns(10)
    db.SetMaxOpenConns(20)
    if err != nil {
        log.Fatal("open db connect fail -> ", err)
        panic(err)
    }

    ctx, stop := context.WithCancel(context.Background())
    defer stop()
    Ping(ctx)
}