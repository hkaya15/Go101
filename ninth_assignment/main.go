package main

import (
	"fmt"
	"log"
	"github.com/hkaya15/Go101/math"

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
	
	
	num:=math.Double(5)
	fmt.Println(num)

}
