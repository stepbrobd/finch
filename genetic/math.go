package genetic

import (
	"crypto/rand"
	"math/big"
)

func Mult(matrixA, vectorB, vectorC []float32, n, m int) int {
	rowCol := 0
	for row := 0; row < n; row++ {
		vectorC[row] = 0.0
		for col := 0; col < m; col++ {
			vectorC[row] += matrixA[rowCol] * vectorB[col]
			rowCol++
		}
	}
	return n * m
}

func Add(vectorA, vectorB, vectorC []float32, n int) int {
	for row := 0; row < n; row++ {
		vectorC[row] = vectorA[row] + vectorB[row]
	}
	return n
}

func ReLU(vectorA []float32, n int) {
	for row := 0; row < n; row++ {
		if vectorA[row] < 0.0 {
			vectorA[row] = 0.0
		}
	}
}

func Abs(value float32) float32 {
	if value < 0.0 {
		value *= -1.0
	}
	return value
}

func RandIntRange(min, max int) int {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(max+1-min)))
	n := nBig.Int64()
	return int(n) + min
}

func RandFloatRange(min, max float32) float32 {
	minInt := int(min * 1000000)
	maxInt := int(max * 1000000)
	return float32(float64(RandIntRange(minInt, maxInt)) / 1000000)
}
