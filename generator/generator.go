package generator

import (
	"fmt"
	"math/rand"
	"proto-benchmark-value-vs-pointers/proto"
)

func randString(l int) string {
	buf := make([]byte, l)
	for i := 0; i < (l+1)/2; i++ {
		buf[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", buf)[:l]
}

func GenerateMessageValue(n int) []*proto.MessageValue {
	out := make([]*proto.MessageValue, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, &proto.MessageValue{
			Name:     randString(10),
			BirthDay: rand.Int63n(5),
			Phone:    randString(10),
			Siblings: rand.Int31n(10),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
			Type:     proto.TypeValue(rand.Intn(4)),
			Address: &proto.AddressValue{
				Street:   randString(10),
				Number:   rand.Int31n(10),
				PostCode: rand.Int31n(10),
				Floor:    rand.Int31n(10),
				Random: &proto.RandomValue{
					FieldA: randString(10),
					FieldB: randString(10),
					FieldC: randString(10),
					FieldD: randString(10),
					FieldE: randString(10),
					FieldF: randString(10),
					FieldG: randString(10),
					FieldH: randString(10),
					FieldI: randString(10),
					FieldJ: randString(10),
					FieldK: rand.Int63n(5),
					FieldL: rand.Int63n(5),
					FieldM: rand.Int63n(5),
					FieldN: rand.Int63n(5),
					FieldO: randString(10),
					FieldP: randString(10),
					FieldQ: rand.Int31n(10),
					FieldR: randString(10),
					FieldS: randString(10),
					FieldT: randString(10),
					FieldU: rand.Int31n(10),
					FieldV: rand.Int31n(10),
					FieldW: rand.Int31n(10),
					FieldX: rand.Int31n(10),
					FieldY: randString(10),
					FieldZ: rand.Intn(2) == 1,
					NestedRandom: &proto.NestedRandomValue{
						FieldA: randString(10),
						FieldB: randString(10),
						FieldC: randString(10),
						FieldD: randString(10),
						FieldE: randString(10),
						FieldF: randString(10),
						FieldG: rand.Float64(),
						FieldH: rand.Float64(),
						FieldI: rand.Float64(),
						FieldJ: rand.Float64(),
						FieldK: rand.Float64(),
						FieldL: rand.Float64(),
						FieldM: randString(10),
						FieldN: randString(10),
						FieldO: randString(10),
						FieldP: randString(10),
						FieldQ: rand.Int63n(5),
						FieldR: rand.Int63n(5),
						FieldS: rand.Int63n(5),
						FieldT: rand.Int63n(5),
						FieldU: randString(10),
						FieldV: randString(10),
						FieldW: randString(10),
						FieldX: randString(10),
						FieldY: rand.Intn(2) == 1,
						FieldZ: rand.Intn(2) == 1,
					},
				},
			},
		})
	}
	return out
}

func GenerateMessageOptional(n int) []*proto.MessageOptional {
	var (
		Name     = randString(10)
		BirthDay = rand.Int63n(5)
		Phone    = randString(10)
		Siblings = rand.Int31n(10)
		Spouse   = rand.Intn(2) == 1
		Money    = rand.Float64()
		Type     = proto.TypeOptional(rand.Intn(4))
		Street   = randString(10)
		Number   = rand.Int31n(10)
		PostCode = rand.Int31n(10)
		Floor    = rand.Int31n(10)
		FieldA   = randString(10)
		FieldB   = randString(10)
		FieldC   = randString(10)
		FieldD   = randString(10)
		FieldE   = randString(10)
		FieldF   = randString(10)
		FieldG   = randString(10)
		FieldH   = randString(10)
		FieldI   = randString(10)
		FieldJ   = randString(10)
		FieldK   = rand.Int63n(5)
		FieldL   = rand.Int63n(5)
		FieldM   = rand.Int63n(5)
		FieldN   = rand.Int63n(5)
		FieldO   = randString(10)
		FieldP   = randString(10)
		FieldQ   = rand.Int31n(10)
		FieldR   = randString(10)
		FieldS   = randString(10)
		FieldT   = randString(10)
		FieldU   = rand.Int31n(10)
		FieldV   = rand.Int31n(10)
		FieldW   = rand.Int31n(10)
		FieldX   = rand.Int31n(10)
		FieldY   = randString(10)
		FieldZ   = rand.Intn(2) == 1
		FieldA1  = randString(10)
		FieldB1  = randString(10)
		FieldC1  = randString(10)
		FieldD1  = randString(10)
		FieldE1  = randString(10)
		FieldF1  = randString(10)
		FieldG1  = rand.Float64()
		FieldH1  = rand.Float64()
		FieldI1  = rand.Float64()
		FieldJ1  = rand.Float64()
		FieldK1  = rand.Float64()
		FieldL1  = rand.Float64()
		FieldM1  = randString(10)
		FieldN1  = randString(10)
		FieldO1  = randString(10)
		FieldP1  = randString(10)
		FieldQ1  = rand.Int63n(5)
		FieldR1  = rand.Int63n(5)
		FieldS1  = rand.Int63n(5)
		FieldT1  = rand.Int63n(5)
		FieldU1  = randString(10)
		FieldV1  = randString(10)
		FieldW1  = randString(10)
		FieldX1  = randString(10)
		FieldY1  = rand.Intn(2) == 1
		FieldZ1  = rand.Intn(2) == 1
	)

	out := make([]*proto.MessageOptional, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, &proto.MessageOptional{
			Name:     &Name,
			BirthDay: &BirthDay,
			Phone:    &Phone,
			Siblings: &Siblings,
			Spouse:   &Spouse,
			Money:    &Money,
			Type:     &Type,
			Address: &proto.AddressOptional{
				Street:   &Street,
				Number:   &Number,
				PostCode: &PostCode,
				Floor:    &Floor,
				Random: &proto.RandomOptional{
					FieldA: &FieldA,
					FieldB: &FieldB,
					FieldC: &FieldC,
					FieldD: &FieldD,
					FieldE: &FieldE,
					FieldF: &FieldF,
					FieldG: &FieldG,
					FieldH: &FieldH,
					FieldI: &FieldI,
					FieldJ: &FieldJ,
					FieldK: &FieldK,
					FieldL: &FieldL,
					FieldM: &FieldM,
					FieldN: &FieldN,
					FieldO: &FieldO,
					FieldP: &FieldP,
					FieldQ: &FieldQ,
					FieldR: &FieldR,
					FieldS: &FieldS,
					FieldT: &FieldT,
					FieldU: &FieldU,
					FieldV: &FieldV,
					FieldW: &FieldW,
					FieldX: &FieldX,
					FieldY: &FieldY,
					FieldZ: &FieldZ,
					NestedRandom: &proto.NestedRandomOptional{
						FieldA: &FieldA1,
						FieldB: &FieldB1,
						FieldC: &FieldC1,
						FieldD: &FieldD1,
						FieldE: &FieldE1,
						FieldF: &FieldF1,
						FieldG: &FieldG1,
						FieldH: &FieldH1,
						FieldI: &FieldI1,
						FieldJ: &FieldJ1,
						FieldK: &FieldK1,
						FieldL: &FieldL1,
						FieldM: &FieldM1,
						FieldN: &FieldN1,
						FieldO: &FieldO1,
						FieldP: &FieldP1,
						FieldQ: &FieldQ1,
						FieldR: &FieldR1,
						FieldS: &FieldS1,
						FieldT: &FieldT1,
						FieldU: &FieldU1,
						FieldV: &FieldV1,
						FieldW: &FieldW1,
						FieldX: &FieldX1,
						FieldY: &FieldY1,
						FieldZ: &FieldZ1,
					},
				},
			},
		})
	}
	return out
}
