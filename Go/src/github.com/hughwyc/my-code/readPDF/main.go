package main

import (
	"fmt"

	"rsc.io/pdf"
)

func main(){
	file, err := pdf.Open(`C:\Users\Hugh\Desktop\test-001\001-140-0.4-1.pdf`)
	if err != nil {
		panic(err)
	}
	fmt.Println(file.NumPage())
	fmt.Println(file.Page(1).Content())
}
