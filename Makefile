all: build-proto bench

bench:
	go test -count=1 -timeout 5m -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out | grep Benchmark_Message > results.txt
	@cat results.txt

build-proto:
	go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
	protoc --go_out=paths=source_relative:. *.proto

cpu-prof:
	go tool pprof profile.out

mem-prof:
	go tool pprof memprofile.out