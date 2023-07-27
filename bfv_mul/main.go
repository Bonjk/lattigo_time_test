package main

import (
	"github.com/tuneinsight/lattigo/v4/bfv"
	"github.com/tuneinsight/lattigo/v4/rlwe"
)

func main() {
	paramDef := bfv.PN13QP218
	paramDef.T = 0x3ee0001
	params, _ := bfv.NewParametersFromLiteral(paramDef)
	keyGenerator := bfv.NewKeyGenerator(params)
	secretKey, _ := keyGenerator.GenKeyPair()

	encoder := bfv.NewEncoder(params)
	encryptor := bfv.NewEncryptor(params, secretKey)
	data := []uint64{76}
	dataPlaintext := bfv.NewPlaintext(params, params.MaxLevel())
	encoder.Encode(data, dataPlaintext)
	dataCiphertext := encryptor.EncryptNew(dataPlaintext)

	evaluator := bfv.NewEvaluator(params, rlwe.EvaluationKey{})
	for i := 0; i < 1000; i++ {
		evaluator.MulNew(dataCiphertext, dataCiphertext)
	}

	return
}
