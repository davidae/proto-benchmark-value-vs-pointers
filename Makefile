all: build-proto bench

bench:
	go test -count=1 -timeout 5m -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out | grep Benchmark_Message > results.txt
	@cat results.txt

build-proto:
	go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
	protoc --go_out=paths=source_relative:. proto/*.proto

mem-prof:
	x-www-browser http://localhost:6060/debug/pprof
	go run cmd/main.go
