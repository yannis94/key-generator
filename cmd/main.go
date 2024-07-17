package main

import (
	"fmt"

	"github.com/yannis94/key-generator/internal/password"
)

func main() {
	fmt.Println("KEY GENERATOR")
	pwd := password.Password{}
	cfg := password.NewPasswordConfig(17, 18, 14)
	if err := pwd.InitConfig(*cfg); err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("password: %s\n", pwd.Generate())
	}
}
