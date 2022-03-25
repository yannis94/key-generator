package printer

import "log"

func intro() string {
    log.Println("------- KEY GENERATOR ------\n\n")
    log.Println("[0] Password / token")
    log.Println("[1] Passphrase")
}

func test() string {
    log.Println("Test successful")
}
