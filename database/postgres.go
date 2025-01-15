// https://jianghushinian.cn/2023/06/05/how-to-use-database-sql-to-operate-database-in-go/

package database

import (
    "log"
    "time"
    "context"
    "database/sql"
    _ "github.com/lib/pq"
)

type Postgres struct {
    Pool *sql.DB
}

func (p *Postgres) ConnectPing() {
    log.Printf("开始进行Postgres的Ping动作")
    ctx, stop := context.WithCancel(context.Background())
    defer stop()

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	log.Printf("Postgres ping db-> ", p.Pool)
	if err := p.Pool.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
        panic(err)
	}
    log.Printf("进行Postgres的Ping动作完成")
}

func (p *Postgres) ConnectOpen() {
    log.Printf("开始打开Postgres链接")
    var err error
    p.Pool, err = sql.Open("postgres", "host=172.25.1.22 port=5432 user=appsmith password=appsmith dbname=appsmith sslmode=disable")
    // db, err = sql.Open("postgres", "postgres://appsmith:appsmith@172.25.1.22:5432/appsmith?sslmode=disable")
    if err != nil {
        log.Fatal("open db connect fail -> ", err)
        panic(err)
    }
    p.Pool.SetConnMaxIdleTime(30*1000)
    p.Pool.SetConnMaxLifetime(10*1000)
    p.Pool.SetMaxIdleConns(10)
    p.Pool.SetMaxOpenConns(20)

    log.Printf("链接已经打开，链接池信息: -> ", p.Pool)
    log.Printf("打开Postgres链接完成")
}

func (p *Postgres) GetPool() (Pool *sql.DB) {
    return p.Pool
}