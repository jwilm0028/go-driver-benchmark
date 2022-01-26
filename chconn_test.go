package gotests_test

import (
	"context"
	"testing"

	"github.com/vahid-sohrabloo/chconn"
	"github.com/vahid-sohrabloo/chconn/column"
)

func BenchmarkTestChconnSelect100MUint64(b *testing.B) {
	// return
	ctx := context.Background()
	c, err := chconn.Connect(ctx, "password=salam")
	if err != nil {
		b.Fatal(err)
	}
	// var datStr [][]byte
	colRead := column.NewUint64(false)
	for n := 0; n < b.N; n++ {
		s, err := c.Select(ctx, "SELECT number FROM system.numbers_mt LIMIT 100000000")
		if err != nil {
			b.Fatal(err)
		}

		// colReadStr := column.NewString(false)
		for s.Next() {
			if err := s.NextColumn(colRead); err != nil {
				b.Fatal(err)
			}
			colRead.GetAllUnsafe()
			// if err := s.NextColumn(colReadStr); err != nil {
			// 	b.Fatal(err)
			// }
			// datStr = datStr[:0]
			// colReadStr.ReadAll(&datStr)
		}
		if err := s.Err(); err != nil {
			b.Fatal(err)
		}
		s.Close()
	}
}

func BenchmarkTestChconnSelect1MString(b *testing.B) {
	// return
	ctx := context.Background()
	c, err := chconn.Connect(ctx, "password=salam")
	if err != nil {
		b.Fatal(err)
	}
	// var datStr [][]byte
	colRead := column.NewString(false)
	var data [][]byte
	for n := 0; n < b.N; n++ {
		s, err := c.Select(ctx, "SELECT randomString(20) FROM system.numbers_mt LIMIT 1000000")
		if err != nil {
			b.Fatal(err)
		}

		for s.Next() {
			data = data[:0]
			if err := s.NextColumn(colRead); err != nil {
				b.Fatal(err)
			}
			colRead.ReadAll(&data)
		}
		if err := s.Err(); err != nil {
			b.Fatal(err)
		}
		s.Close()
	}
}

func BenchmarkTestChconnInsert10M(b *testing.B) {
	// return
	ctx := context.Background()
	c, err := chconn.Connect(ctx, "password=salam")
	if err != nil {
		b.Fatal(err)
	}
	_, err = c.Exec(ctx, "DROP TABLE IF EXISTS test_insert_chconn")
	if err != nil {
		b.Fatal(err)
	}
	_, err = c.Exec(ctx, "CREATE TABLE test_insert_chconn (id UInt64,v String) ENGINE = Null")
	if err != nil {
		b.Fatal(err)
	}

	const (
		rowsInBlock = 10_000_000
	)

	idColumns := column.NewUint64(false)
	vColumns := column.NewString(false)
	for n := 0; n < b.N; n++ {
		idColumns.Reset()
		vColumns.Reset()
		for y := 0; y < rowsInBlock; y++ {
			idColumns.Append(1)
			vColumns.Append([]byte("test"))
		}
		err := c.Insert(ctx, "INSERT INTO test_insert_chconn VALUES", idColumns, vColumns)
		if err != nil {
			b.Fatal(err)
		}

	}
}
