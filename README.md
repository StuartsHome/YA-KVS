# Yet Another - Key Value Store

What is a database other than a hashmap? `YAKVS` provides an in-memory DB.

How advanced can we make this database (hashmap)?
- Client?
- Distributed transactions?
- Consensus? Raft?
- Spatial Indexing?
- Benchmarks?


### Design Principles
In-memory DB.
ACI - No durability as in-memory

### Concurrency
One process at a time.
Each process acquires a lock.

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

#### Thoughts
Transactions using MVCC.


There can only be a single concurrent writer, but any
number of readers.

