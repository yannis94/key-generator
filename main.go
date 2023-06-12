package main

import (
	"flag"
	"fmt"

	"github.com/yannis94/key-generator/passphrase"
	"github.com/yannis94/key-generator/password"
	"github.com/yannis94/key-generator/printer"
)

func manualUsage(config map[string]int64) string {
    var result string
    printer.Intro()
    choice := printer.AskChoice()

    if choice == "1" {
        config = printer.SetPasswordConfiguration()
        result = password.Generate(config)
    } else if choice == "2" {
        config = printer.SetPassphraseConfiguration()
        result = passphrase.Generate(config)
    } 

    return result
}

func main() {
    passwordGen := flag.Bool("password", false, "Password generation")
    passphraseGen := flag.Bool("phrase", false, "Key generation")
    numLetter := flag.Int64("l", 0, "Specify how many letter you want")
    numNumber := flag.Int64("n", 0, "Specify how many number you want")
    numChar := flag.Int64("c", 0, "Specify how many special character you want")
    numWord := flag.Int64("w", 0, "Specify how many word you want")

    flag.Parse()

    var userConfig map[string]int64
    var result string

    userConfig = make(map[string]int64)
    userConfig["letter"] = *numLetter
    userConfig["number"] = *numNumber
    userConfig["char"] = *numChar
    userConfig["words"] = *numWord

    if *passwordGen {
        result = password.Generate(userConfig)
    } else if *passphraseGen {
        result = passphrase.Generate(userConfig)
    } else {
        result = manualUsage(userConfig)
        printer.ShowResult(result)
        return 
    }

    fmt.Println(result)

}
