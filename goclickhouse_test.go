package gotests_test

import (
	"context"
	"testing"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func BenchmarkTestGoclickhouseSelect100MUint64(b *testing.B) {
	c, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "salam",
		},
	})
	if err != nil {
		b.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		rows, err := c.Query(context.Background(), "SELECT number FROM system.numbers_mt LIMIT 100000000")
		if err != nil {
			b.Fatal(err)
		}
		var count int
		for rows.Next() {
			count++
		}
	}
}
func BenchmarkTestGoclickhouseSelect1MString(b *testing.B) {
	c, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "salam",
		},
	})
	if err != nil {
		b.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		rows, err := c.Query(context.Background(), "SELECT randomString(20) FROM system.numbers_mt LIMIT 1000000")
		if err != nil {
			b.Fatal(err)
		}
		var count int
		for rows.Next() {
			count++
		}
	}
}

func BenchmarkTestGoclickhouseInsert10M(b *testing.B) {
	// return
	c, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "salam",
		},
	})
	if err != nil {
		b.Fatal(err)
	}

	ctx := context.Background()

	err = c.Exec(ctx, `
		DROP TABLE IF EXISTS test_insert_go_goclickhouse
	`)
	if err != nil {
		b.Fatal(err)
	}
	err = c.Exec(ctx, `
			CREATE TABLE test_insert_go_goclickhouse (id UInt64,v String) ENGINE = Null
	`)
	if err != nil {
		b.Fatal(err)
	}

	const (
		rowsInBlock = 10_000_000
	)
	var (
		col1 []uint64
		col2 []string
	)
	for n := 0; n < b.N; n++ {
		col1 = col1[:0]
		col2 = col2[:0]
		for i := 0; i < rowsInBlock; i++ {
			col1 = append(col1, 1)
			col2 = append(col2, "test")
		}
		batch, err := c.PrepareBatch(ctx, "INSERT INTO test_insert_go_goclickhouse VALUES")
		if err != nil {
			b.Fatal(err)
		}
		if err := batch.Column(0).Append(col1); err != nil {
			b.Fatal(err)
		}
		if err := batch.Column(1).Append(col2); err != nil {
			b.Fatal(err)
		}

		if err = batch.Send(); err != nil {
			b.Fatal(err)
		}
	}
}
