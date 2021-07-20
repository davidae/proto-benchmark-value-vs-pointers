package main

import (
	"math/rand"
	"testing"

	"google.golang.org/protobuf/proto"
)

func Benchmark_MessageOptional_Proto_Marshal(b *testing.B) {
	data := GenerateMessageOptional(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := proto.Marshal(data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_MessageOptional_Proto_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := GenerateMessageOptional(b.N)
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = proto.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	messages := make([]*MessageOptional, 0, len(data))
	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &MessageOptional{}
		err := proto.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
		messages = append(messages, o)
	}

	ParseMessageOptionals(messages)
}

func Benchmark_MessageValue_Proto_Marshal(b *testing.B) {
	data := GenerateMessageValue(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := proto.Marshal(data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_MessageValue_Proto_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := GenerateMessageValue(b.N)
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = proto.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	messages := make([]*MessageValue, 0, len(data))
	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &MessageValue{}
		err := proto.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
		messages = append(messages, o)
	}

	ParseMessageValues(messages)
}

func ParseMessageOptionals(ms []*MessageOptional) {
	for _, m := range ms {
		ParseRandomOptional(m.Address.Random)
	}
}

func ParseRandomOptional(r *RandomOptional) {
	r.NestedRandom = ParseRandomNestedOptional(r.NestedRandom)
}

func ParseRandomNestedOptional(r *NestedRandomOptional) *NestedRandomOptional {
	x := "xxx"
	r.FieldA = &x
	return r
}

func ParseMessageValues(ms []*MessageValue) {
	for _, m := range ms {
		ParseRandomValue(m.Address.Random)
	}
}

func ParseRandomValue(r *RandomValue) {
	r.NestedRandom = ParseRandomNestedValue(r.NestedRandom)
}

func ParseRandomNestedValue(r *NestedRandomValue) *NestedRandomValue {
	r.FieldA = "xxx"
	return r
}
