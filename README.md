## Yet Another Key-Value Store



### Design Principles
In-memory DB.
ACI - No durability as in-memory


### Benchmarks
Get()
```
goos: darwin
goarch: amd64
pkg: github.com/StuartsHome/YA-KVS/handler
cpu: Intel(R) Core(TM) M-5Y31 CPU @ 0.90GHz
BenchmarkHandleGet-4   	 4240266	       274.9 ns/op	     256 B/op	       2 allocs/op
PASS
coverage: [no statements]
ok  	github.com/StuartsHome/YA-KVS/handler	1.902s
```

### Concurrency
One process at a time.
Each process acquires a lock.

#### Thoughts
Transactions using MVCC.


There can only be a single concurrent writer, but any
number of readers.

