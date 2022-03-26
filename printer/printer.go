package printer

import (
    "log"
    "bufio"
    "os"
    "strings"
)


func getInput (question string, r *bufio.Reader) (string, error) {
    log.SetFlags(0)
    log.Print(question)
    input, err := r.ReadString('\n')

    return strings.TrimRight(input, "\n"), err
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

    answerValid := false
    var userChoice string
    var menu string = "\n\n[1] Password / key\n[2] Passphrase\n[q] Quit"
    mypointer := &userChoice
    
    log.Println(menu)
    
    for !answerValid {
        reader := bufio.NewReader(os.Stdin)
        choice, _ := getInput("\nMake your choice > ", reader)

        switch {
            case choice == "1": 
                answerValid = true
                break
            case choice == "2": 
                answerValid = true
                break
            case choice == "q":
                answerValid = true
            default: 
                log.Println("This choice doen't exist.")
                log.Println(menu)
        }

        *mypointer = choice
    }

    return userChoice
}
