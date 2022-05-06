package printer

import (
    "log"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func getIntInput(question string, r *bufio.Reader) int64 {
    var answerValid bool = false
    var response int64
    mypointer := &response

    log.SetFlags(0)
    for !answerValid {
        log.Print(question)
        input, err := r.ReadString('\n')

        if err != nil {
            log.Panic("Error, something went wrong")
        }
        
        userValue := strings.TrimRight(input, "\n")

        answer, err := strconv.ParseInt(userValue, 10, 64)

        if err != nil {
            log.Print("Invalid answer.")
        } else {
            answerValid = true
            *mypointer = answer
        }

    }

    return response

}

func getStringInput(question string, acceptedValue []string, r *bufio.Reader) string {
    var answerValid bool = false
    var response string
    mypointer := &response

    log.SetFlags(0)

    for !answerValid {
        log.Print(question)
        log.Print("\nYour choice >")
        input, err := r.ReadString('\n')

        if err != nil {
            log.Panic("Error, something went wrong")
        }
        
        userValue := strings.TrimRight(input, "\n")

        for i, value := range acceptedValue {
            if userValue == value {
                answerValid = true
                *mypointer = userValue
            } else if i - 1 == len(acceptedValue) {
                log.Print("Invalid answer.")
            } else {
                continue
            }
        }
    }

    return response
}

func Intro() {
    var subheader string = "----------------------------------------------------"
    var title string = "\n------------------ KEY GENERATOR -------------------\n"
    var header = subheader + title + subheader + "\n\n"
    var question string = "Welcome to @yannis94 generator,\nplease make your choice."

    log.SetFlags(0)
    log.Println(header + question)

    return 
}

func AskChoice() string {
    log.SetFlags(0)

    var menu string = "\n\n[1] Password / key\n[2] Passphrase\n[q] Quit"
    
    reader := bufio.NewReader(os.Stdin)
    
    userChoice := getStringInput(menu, []string{"1", "2", "q"}, reader)

    return userChoice
}

func SetPasswordConfiguration() map[string]int64 {
    config := make(map[string]int64)
    reader := bufio.NewReader(os.Stdin)

    config["letter"] = getIntInput("How many letter >", reader)
    config["number"] = getIntInput("How many digit >", reader)
    config["char"] = getIntInput("How many special character >", reader)

    return config
}

func SetPassphraseConfiguration() map[string]int64 {
    config := make(map[string]int64)
    reader := bufio.NewReader(os.Stdin)

    config["words"] = getIntInput("How many words do you need >", reader)

    return config
}

func ShowResult(res string) {
    log.Printf("\n------- RESULT -------\n\n")
    log.Print(res)
    log.Printf("\n-------------")
}

func Quit() {
    log.SetFlags(0)
    log.Printf("\nQuitting, bye")
}
