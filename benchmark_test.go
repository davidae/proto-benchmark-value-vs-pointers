package main

import (
	"math/rand"
	"testing"

	"proto-benchmark-value-vs-pointers/generator"
	"proto-benchmark-value-vs-pointers/proto"

	goproto "google.golang.org/protobuf/proto"
)

func Benchmark_MessageOptional_Proto_Marshal(b *testing.B) {
	data := generator.GenerateMessageOptional(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := goproto.Marshal(data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_MessageOptional_Proto_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generator.GenerateMessageOptional(b.N)
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = goproto.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}

	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	messages := make([]*proto.MessageOptional, 0, len(data))
	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &proto.MessageOptional{}
		err := goproto.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
		messages = append(messages, o)
	}

	ParseMessageOptionals(messages)
}

func Benchmark_MessageValue_Proto_Marshal(b *testing.B) {
	data := generator.GenerateMessageValue(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := goproto.Marshal(data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_MessageValue_Proto_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generator.GenerateMessageValue(b.N)
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = goproto.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	messages := make([]*proto.MessageValue, 0, len(data))
	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &proto.MessageValue{}
		err := goproto.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
		messages = append(messages, o)
	}

	ParseMessageValues(messages)
}

func ParseMessageOptionals(ms []*proto.MessageOptional) {
	for _, m := range ms {
		ParseRandomOptional(m.Address.Random)
	}
}

func ParseRandomOptional(r *proto.RandomOptional) {
	r.NestedRandom = ParseRandomNestedOptional(r.NestedRandom)
}

func ParseRandomNestedOptional(r *proto.NestedRandomOptional) *proto.NestedRandomOptional {
	x := "xxx"
	r.FieldA = &x
	return r
}

func ParseMessageValues(ms []*proto.MessageValue) {
	for _, m := range ms {
		ParseRandomValue(m.Address.Random)
	}
}

func ParseRandomValue(r *proto.RandomValue) {
	r.NestedRandom = ParseRandomNestedValue(r.NestedRandom)
}

func ParseRandomNestedValue(r *proto.NestedRandomValue) *proto.NestedRandomValue {
	r.FieldA = "xxx"
	return r
}
