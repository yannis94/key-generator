package password

import (
    "time"
    "math/rand"
)

func pickRandomNtime(n int64, target string) string {
    var result string
    var i int64
    rand.Seed(time.Now().UnixNano())
    maxInt := len(target)

    for i = 0; i < n; i++ {
        randIndex := rand.Intn(maxInt)
        pick := byte(target[randIndex])
        result += string(pick)
    }

    return result
}

func Generate(config map[string]int64) string {
    letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    digits := "0123456789"
    char := "-_/\\|[]{}()?!.,;:'\"!@#$%^&*=+"
    var result string

    passwordLetters := pickRandomNtime(config["letter"], letters)
    passwordNumbers := pickRandomNtime(config["number"], digits)
    passwordChar := pickRandomNtime(config["char"], char)
    
    dumbPassword := passwordLetters + passwordNumbers + passwordChar
    rand.Seed(time.Now().Unix())
    prov := []rune(dumbPassword)
    rand.Shuffle(len(prov), func (i, j int) {
        prov[i], prov[j] = prov[j], prov[i]
    })

    result = string(prov)

    return result
}

