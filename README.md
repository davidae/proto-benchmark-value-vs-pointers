# Basic benchmarking of protobuf struct with and without optional fields
What impact, in terms of performance and memory allocations does a pointer and value struct have?

## Results
```
Benchmark_MessageOptional_Proto_Marshal-12      	  579118	      2468 ns/op	       537.0 B/serial	     576 B/op	       1 allocs/op
Benchmark_MessageOptional_Proto_Unmarshal-12    	  277396	      4289 ns/op	       537.0 B/serial	    1976 B/op	     101 allocs/op
Benchmark_MessageValue_Proto_Marshal-12         	  547080	      2868 ns/op	       524.3 B/serial	     575 B/op	       1 allocs/op
Benchmark_MessageValue_Proto_Unmarshal-12       	  453637	      3002 ns/op	       524.3 B/serial	    1480 B/op	      37 allocs/op
```