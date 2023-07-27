package main

import (
	"github.com/tuneinsight/lattigo/v4/ckks"
	"github.com/tuneinsight/lattigo/v4/rlwe"
)

func main() {
	paramDef := ckks.PN13QP218
	params, _ := ckks.NewParametersFromLiteral(paramDef)
	keyGenerator := ckks.NewKeyGenerator(params)
	secretKey, _ := keyGenerator.GenKeyPair()
	rKey := keyGenerator.GenRelinearizationKey(secretKey, 1)

	encoder := ckks.NewEncoder(params)
	encryptor := ckks.NewEncryptor(params, secretKey)
	data := []complex128{76.6}
	dataPlaintext := ckks.NewPlaintext(params, params.MaxLevel())
	encoder.Encode(data, dataPlaintext, params.LogSlots())
	dataCiphertext := encryptor.EncryptNew(dataPlaintext)

	evaluator := ckks.NewEvaluator(params, rlwe.EvaluationKey{Rlk: rKey})
	anotherCiphertext := dataCiphertext.CopyNew()
	resultCiphertext := evaluator.AddNew(dataCiphertext, anotherCiphertext)
	for i := 0; i < 999; i++ {
		evaluator.Add(dataCiphertext, resultCiphertext, resultCiphertext)
	}
	return
}
