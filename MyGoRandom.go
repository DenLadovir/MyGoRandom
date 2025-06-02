package MyGoRandom

import (
	mathRand "math/rand"
)

type Integers interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Floats interface {
	~float64 | ~float32
}

func RandFloat[T Floats](min, max T) float64 {
	if max < min {
		panic("RandFloat: max must be >= min")
	}
	return float64(min) + mathRand.Float64()*float64((max-min)+1)
}

func RandInt[T Integers](min, max T) T {
	if max < min {
		panic("RandInt: max must be >= min")
	}
	return T(int(min) + mathRand.Intn(int(max-min)))
}

func Choice[T any](arr []T) (T, int) {
	if len(arr) == 0 {
		panic("Empty slice")
	}
	var index int = RandInt(0, len(arr)-1)
	return arr[index], index
}

func Shuffle[T any](arr []T) {
	mathRand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}
