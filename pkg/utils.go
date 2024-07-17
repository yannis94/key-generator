package pkg

import (
	"math/rand"
	"time"
)

var _rand *rand.Rand

func init() {
	_rand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GetRandomNbr(min, max int) int {
	return _rand.Intn(max-min+1) + min
}

func ShuffleString(str string) string {
	strChars := []rune(str)
	_rand.Shuffle(len(strChars), func(i, j int) {
		strChars[i], strChars[j] = strChars[j], strChars[i]
	})

	return string(strChars)
}
