package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/google/uuid"
)

func main() {
	// go mod init *URL
	// go get *URL
	// go env -w GO111MODULE=on
	// go mod tidy
	u, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	 color.Cyan("Prints text in cyan.")
	fmt.Println(u.String())
}
