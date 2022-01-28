We use ClickkHouse for time-series databases, and the driver's performance is very important to us (Especially when inserting data). Here is a comparison of the three Go drivers that use Native protocol.
([chconn](https://github.com/vahid-sohrabloo/chconn), [gofaster](https://github.com/go-faster/ch), [goclickhhouse](https://github.com/ClickHouse/clickhouse-go))

See their website if you are unfamiliar with ClickHouse:
[https://clickhouse.com/](https://clickhouse.com/)

Obviously, these tests are meant to help us decide and not to tell you which option is best for your project.

Using the following command, you can test the project on your computer
```
go test  -run=. -bench=. -benchtime=20x   -benchmem 

```

You  can also see [ch-bench](https://github.com/go-faster/ch-bench#benchmarks)


## Result
| Lib           | Function                     | ns/op         | B/op          | allocs/op |
| ------------- | ---------------------------- | ------------- | ------------- | --------- |
| chconn        | Select 100M uint64           | **156132671** | **33611**     | **14**    |
| go-faster     | Select 100M uint64           | **156665895** | 124732        | 6394      |
| go-clickhouse | Select 100M uint64           | 795474396     | 801195961     | 18829     |
|               |                              |               |               |           |
| chconn        | Select 1M string             | **34021204**  | **243267**    | 3288      |
| go-faster     | Select 1M string             | 52143253      | 684490        | **176**   |
| go-clickhouse | Select 1M string             | 116804874     | 137189293     | 2000766   |
|               |                              |               |               |           |
| chconn        | Insert 10m uint64 and string | **183496685** | **34727882**  | **22**    |
| go-faster     | Insert 10m uint64 and string | 258597566     | 96464450      | **55**    |
| go-clickhouse | Insert 10m uint64 and string | 346469462     | 402921766     | **66**    |

