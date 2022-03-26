package main

import (
    "log"
    printer "key-generator/printer"
    password "key-generator/password"
    passphrase "key-generator/passphrase"
)

func main() {
    // 1. call introduction func 
    // 2. get user choice
    // 3. redirect to corresponding func
    // 4. output result
    printer.Intro()

    choice := printer.AskChoice()

    if choice == "1" {
        userConfig := printer.SetPasswordConfiguration()
        result := password.Generate(userConfig)
        log.Print(result)
    } else if choice == "2" {
        log.Printf(passphrase.Generate())
    } else {
        log.Printf("Quitting, bye")
    }
}
