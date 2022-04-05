package main

import (
	"fmt"
	"github.com/c3b2a7/picup/apis"
	"os"
)

func main() {
	smms := apis.SMMS{Token: os.Getenv("SMMS_TOKEN")}
	up, err := smms.Up(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(up)
}
