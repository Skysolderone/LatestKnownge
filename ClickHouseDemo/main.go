package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func main() {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"localhost:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		DialTimeout: 5 * time.Second,
		Debug:       false,
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	if err := conn.Ping(ctx); err != nil {
		log.Fatalf("无法连接到 ClickHouse: %v", err)
	}

	// 创建表
	err = conn.Exec(ctx, `
        CREATE TABLE IF NOT EXISTS example_table (
            id UInt64,
            name String,
            age UInt8
        ) ENGINE = Memory
    `)
	if err != nil {
		log.Fatal(err)
	}

	// 插入数据
	batch, err := conn.PrepareBatch(ctx, "INSERT INTO example_table")
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= 5; i++ {
		err = batch.Append(uint64(i), fmt.Sprintf("User%d", i), uint8(20+i))
		if err != nil {
			log.Fatal(err)
		}
	}
	if err = batch.Send(); err != nil {
		log.Fatal(err)
	}

	// 查询数据
	rows, err := conn.Query(ctx, `SELECT id, name, age FROM example_table`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id   uint64
			name string
			age  uint8
		)
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d, name: %s, age: %d\n", id, name, age)
	}

	// 关闭连接
	if err := conn.Close(); err != nil {
		log.Fatal(err)
	}
}
