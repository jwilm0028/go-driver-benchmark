We use ClickkHouse for time-series databases, and the driver's performance is very important to us (Especially when inserting data). Here is a comparison of the three Go drivers that use Native protocol.
([chconn](https://github.com/vahid-sohrabloo/chconn), [gofaster](https://github.com/go-faster/ch), [goclickhhouse](https://github.com/ClickHouse/clickhouse-go))

Obviously, these tests are meant to help us decide and not to tell you which option is best for your project.

Using the following command, you can test the project on your computer
```
go test  -run=. -bench=. -benchtime=20x   -benchmem 

```

You  can also see [ch-bench](https://github.com/go-faster/ch-bench#benchmarks)

## Result
| lib          	| Function                     	| ns/op      	| B/op       	| allocs/op 	|
|--------------	|------------------------------	|------------	|------------	|-----------	|
| chconn       	| Select 100M uint64           	| *161013288*  	| *33637*      	| *14*        	|
| go-faster    	| Select 100M uint64           	| *177033114*  	| 129064     	| 6419      	|
| goclickhouse 	| Select 100M uint64           	| 615800261  	| 804415334  	| 804415334 	|
|              	|                              	|            	|            	|           	|
| chconn       	| Select 10M string            	| *3243895219* 	| *248436*     	| 3320      	|
| go-faster    	| Select 10M string            	| 487056596  	| 788240     	| *1348*      	|
| goclickhouse 	| Select 10M string            	| 894460988  	| 1053260923 	| 20007941  	|
|              	|                              	|            	|            	|           	|
| chconn       	| Insert 10m uint64 and string 	| *196375920*  	| *34727680*   	| *21*        	|
| go-faster    	| Insert 10m uint64 and string 	| 287801818  	| 96465054   	| *56*        	|
| goclickhouse 	| Insert 10m uint64 and string 	| 1088248059 	| 1654666726 	| 10000155  	|