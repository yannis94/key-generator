package passphrase

import (
   "io/ioutil"
   "log"
   "time"
   "math/rand"
)

func Generate(config map[string]int64) string {
    n := config["words"]
    words := getWordList()
    passphrase := buildPassphrase(n, words)
    return passphrase
}

func getWordList() []string {
    listOfBytes, err := ioutil.ReadFile("wordslist.txt")
    var words []string
    var provWord string

    if err != nil {
        log.Print(err)
    }
    for i := 0; i < len(listOfBytes); i++ {
        bit := listOfBytes[i]
        if bit == 10 {
            words = append(words, provWord)
            provWord = ""
        } else {
            provWord += string(bit)
        }
    }

    return words
}

func buildPassphrase(number int64, list[]string) string {
    rand.Seed(time.Now().Unix())
    var i int64
    var passphrase string

    for i = 0; i < number; i++ {
        random := rand.Intn(len(list))
        
        passphrase += list[random]

        if i != number - 1 {
            passphrase += "-"
        }
    }

    return passphrase
}
