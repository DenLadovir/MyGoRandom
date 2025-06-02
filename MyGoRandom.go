package MyGoRandom

import (
	cryptoRand "crypto/rand"
	"math/big"
	mathRand "math/rand"
)

func RandFloat(min, max float64) float64 {
	return min + mathRand.Float64()*(max-min)
}

func RandInt(min, max int) int {
	var maxBig *big.Int = big.NewInt(int64(max - min + 1))
	num, err := cryptoRand.Int(cryptoRand.Reader, maxBig)
	if err != nil {
		panic(err)
	}
	return int(num.Int64()) + min
}

func Choice(arr []int) int {
	var index int = RandInt(0, len(arr)-1)
	return arr[index]
}

func Shuffle(arr []int) {
	mathRand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}
