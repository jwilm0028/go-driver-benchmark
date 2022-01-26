package gotests_test

import (
	"context"
	"io"
	"testing"

	"github.com/go-faster/ch"
	"github.com/go-faster/ch/proto"
)

func BenchmarkTestGofasterSelect100MUint64(b *testing.B) {
	ctx := context.Background()
	c, err := ch.Dial(ctx, ch.Options{
		Password: "salam",
		Address:  "localhost:9000",
	})
	if err != nil {
		b.Fatal(err)
	}
	defer func() { _ = c.Close() }()
	var (
		data proto.ColUInt64
		// data2 proto.ColStr
	)
	for n := 0; n < b.N; n++ {

		if err := c.Do(ctx, ch.Query{
			Body: "SELECT number FROM system.numbers_mt LIMIT 100000000",
			OnProgress: func(ctx context.Context, p proto.Progress) error {
				// gotBytes += p.Bytes
				return nil
			},
			OnResult: func(ctx context.Context, block proto.Block) error {
				// gotRows += uint64(block.Rows)
				return nil
			},
			Result: proto.Results{
				{Name: "number", Data: &data},
				// {Name: "st", Data: &data2},
			},
		}); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkTestGofasterSelect1MString(b *testing.B) {
	ctx := context.Background()
	c, err := ch.Dial(ctx, ch.Options{
		Password: "salam",
		Address:  "localhost:9000",
	})
	if err != nil {
		b.Fatal(err)
	}
	defer func() { _ = c.Close() }()
	var (
		data proto.ColStr
		// data2 proto.ColStr
	)
	var dataStr [][]byte
	for n := 0; n < b.N; n++ {

		if err := c.Do(ctx, ch.Query{
			Body: "SELECT randomString(20) as number FROM system.numbers_mt LIMIT 1000000",
			OnProgress: func(ctx context.Context, p proto.Progress) error {
				// gotBytes += p.Bytes
				return nil
			},
			OnResult: func(ctx context.Context, block proto.Block) error {
				dataStr = dataStr[:0]
				data.ForEachBytes(func(i int, b []byte) error {
					// dataStr = append(dataStr, b)
					return nil
				})
				return nil
			},
			Result: proto.Results{
				{Name: "number", Data: &data},
				// {Name: "st", Data: &data2},
			},
		}); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkTestGofasterInsert10M(b *testing.B) {
	// return
	ctx := context.Background()
	c, err := ch.Dial(ctx, ch.Options{
		Password: "salam",
		Address:  "localhost:9000",
	})
	if err != nil {
		b.Fatal(err)
	}
	defer func() { _ = c.Close() }()

	if err := c.Do(ctx, ch.Query{
		Body: "DROP TABLE IF EXISTS test_insert_gofaster",
	}); err != nil {
		b.Fatal(err)
	}
	if err := c.Do(ctx, ch.Query{
		Body: "CREATE TABLE test_insert_gofaster (id UInt64,v String) ENGINE = Null",
	}); err != nil {
		b.Fatal(err)
	}

	const (
		rowsInBlock = 10_000_000
	)

	var idColumns proto.ColUInt64
	var vColumns proto.ColStr
	for n := 0; n < b.N; n++ {
		idColumns.Reset()
		vColumns.Reset()
		for i := 0; i < rowsInBlock; i++ {
			idColumns = append(idColumns, 1)
			vColumns.AppendBytes([]byte("test"))
		}
		if err := c.Do(ctx, ch.Query{
			Body: "INSERT INTO test_insert_gofaster VALUES",
			OnInput: func(ctx context.Context) error {
				return io.EOF
			},
			Input: []proto.InputColumn{
				{Name: "id", Data: idColumns},
				{Name: "v", Data: vColumns},
			},
		}); err != nil {
			b.Fatal(err)
		}
	}
}
