package main

import (
    printer "key-generator/printer"
    password "key-generator/password"
    passphrase "key-generator/passphrase"
)

func main() {
    printer.Intro()

    choice := printer.AskChoice()

    if choice == "1" {
        userConfig := printer.SetPasswordConfiguration()
        result := password.GetPassword(userConfig)
        printer.ShowResult(result)
    } else if choice == "2" {
        userConfig := printer.SetPassphraseConfiguration()
        result := passphrase.Generate(userConfig)
        printer.ShowResult(result)
    } else {
        printer.Quit()
    }
}
