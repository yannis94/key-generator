package main

import (
    "log"
    printer "key-generator/printer"
    password "key-generator/generator/password"
    passphrase "key-generator/generator/passphrase"
)

func intro() string {
    var intro string = "------- KEY GENERATOR ------\n\n[0] Password / token\n[1] Passphrase"

    return intro
}

func main() {
    // 1. call introduction func 
    // 2. get user choice
    // 3. redirect to corresponding func
    // 4. output result
    var intro string = printer.Intro()
    log.Println(intro)
    log.Printf(passphrase.Test())
    log.Printf(password.Test())
}
