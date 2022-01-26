We use ClickkHouse for time-series databases, and the driver's performance is very important to us (Especially when inserting data). Here is a comparison of the three Go drivers that use Native protocol.
([chconn](https://github.com/vahid-sohrabloo/chconn), [gofaster](https://github.com/go-faster/ch), [goclickhhouse](https://github.com/ClickHouse/clickhouse-go))

Obviously, these tests are meant to help us decide and not to tell you which option is best for your project.

Using the following command, you can test the project on your computer
```
go test  -run=. -bench=. -benchtime=20x   -benchmem 

```

You  can also see [ch-bench](https://github.com/go-faster/ch-bench#benchmarks)


## Result
| lib          	| Function                     	| ns/op      	    | B/op       	    | allocs/op 	|
|--------------	|------------------------------	|------------------	|-------------------|-----------	|
| chconn       	| Select 100M uint64           	| **159957761**  	| **33604**      	| **14**        |
| go-faster    	| Select 100M uint64           	| **161677557**  	| 124695     	    | 6392      	|
| goclickhouse 	| Select 100M uint64           	| 619731971  	    | 804420292  	    | 18418     	|
|              	|                              	|               	|            	    |           	|
| chconn       	| Select 1M string            	| **39138253**  	| **243380**     	| 3288      	|
| go-faster    	| Select 1M string            	| 50995805  	    | 684746     	    | **175**       |
| goclickhouse 	| Select 1M string            	| 82825676      	| 107890787 	    | 2000752  	|
|              	|                              	|               	|            	    |           	|
| chconn       	| Insert 10m uint64 and string 	| **192278818**  	| **34727773**   	| **22**        |
| go-faster    	| Insert 10m uint64 and string 	| 275889697  	    | 96464431   	    | **55**        |
| goclickhouse 	| Insert 10m uint64 and string 	| 985197338 	    | 1654669134 	    | 10000157  	|